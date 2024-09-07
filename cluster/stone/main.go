package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"github.com/iamsoloma/butterfly"
	"github.com/iamsoloma/butterfly/system"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "9060"
		fmt.Println("Can`t parse PORT! Used standard: 1106.")
	}

	timeout, err := strconv.Atoi(os.Getenv("TIMEOUT"))
	if err != nil {
		timeout = int(time.Duration.Seconds(60))
		fmt.Println("Can`t parse TIMEOUT! Used standard: 60 sec.")
	}

	bodyLimit, err := strconv.Atoi(os.Getenv("BODYLIMIT"))
	if err != nil {
		bodyLimit = 1024 * 1024 * 1024 * 1024
		fmt.Println("Can`t parse BODYLIMIT! Used standard: 1024*1024*1024*1024.")
	}

	StoragePath := os.Getenv("STORAGE_PATH")
	if StoragePath == "" {
		StoragePath = "./storage/"
		fmt.Println("Can`t parse STORAGE_PATH! Used standard: ./storage .")
	}

	s := NewServer(":" + port, bodyLimit, time.Duration(timeout), StoragePath)
	log.Fatal(s.Start())
}

func (s *Server) Start() error {
	viewEngine := html.New("cluster/stone/templates", ".html")
	f := fiber.New(
		fiber.Config{
			BodyLimit:         s.bodyLimit,
			IdleTimeout:       s.idleTimeout,
			Prefork:           false,
			StreamRequestBody: true,
			Views: viewEngine,
		},
	)

	//main
	f.Get("/health", s.Health)

	storeapi:=f.Group("/store")
	storeapi.Post("/upload/fromHTML/:inpname/:lib/*", s.Upload)

	ui:=f.Group("/ui")
	ui.Get("/upload", s.UploadUI)

	return f.Listen(s.listenAddr)
}

func (s *Server) Health(c *fiber.Ctx) (err error) {
	memory := system.ReadMemoryStats()

	resp := butterfly.Health{
		Status:           "ok",
		UTC:              time.Now().UTC().String(),
		NodeType:         "stone",
		Version:          "0.1.0",
		TotalMemory:     memory.MemTotal,
		AvailableMemory: memory.MemAvailable,
	}
	return c.JSON(resp)
}
