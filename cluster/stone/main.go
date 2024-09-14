package main

import (
	"fmt"
	"log"
	"os"

	//"strconv"
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
		fmt.Println("Can`t parse PORT! Used standard: 9060.")
	}

	StoragePath := os.Getenv("STORAGE_PATH")
	if StoragePath == "" {
		StoragePath = "./storage/"
		fmt.Println("Can`t parse STORAGE_PATH! Used standard: ./storage/ .")
	}

	DustAddress := os.Getenv("DUST_ADDR")
	if DustAddress == "" {
		DustAddress = "http://0.0.0.0:1106"
		fmt.Println("Can`t parse DUST_ADDR! Used standard: http://0.0.0.0:1106.")
	}

	s := NewServer(":"+port, time.Duration(int(time.Duration.Seconds(60))), StoragePath, DustAddress)
	log.Fatal(s.Start())
}

func (s *Server) Start() error {
	viewEngine := html.New("cluster/stone/templates", ".html")
	f := fiber.New(
		fiber.Config{
			IdleTimeout:       s.idleTimeout,
			Prefork:           false,
			StreamRequestBody: true,
			Views:             viewEngine,
		},
	)

	//main
	f.Get("/health", s.Health)

	storeapi := f.Group("/store")
	storeapi.Post("/fromForm/:inpname/:lib/*", s.UploadFromForm)

	storeapi.Get("/:lib/*", s.Get)

	ui := f.Group("/ui")
	ui.Get("/upload", s.UploadUI)

	return f.Listen(s.listenAddr)
}

func (s *Server) Health(c *fiber.Ctx) (err error) {
	memory := system.ReadMemoryStats()
	disk := system.ReadDiskInfo(s.StoragePath)

	resp := butterfly.Health{
		Status:          "ok",
		UTC:             time.Now().UTC().String(),
		NodeType:        "stone",
		Version:         "0.1.0",
		FreeMemory:      memory.MemFree,
		AvailableMemory: memory.MemAvailable,
		FreeDisk:        disk.DiskAvailable,
		AvailableDisk:   disk.DiskAvailable,
	}
	return c.JSON(resp)
}
