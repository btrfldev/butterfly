package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/btrfldev/butterfly"
	"github.com/btrfldev/butterfly/system"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "9060"
		fmt.Println("Can`t parse PORT! Used standard: 9060.")
	}

	storagePath := os.Getenv("STORAGE_PATH")
	if storagePath == "" {
		storagePath = "./storage/"
		fmt.Println("Can`t parse STORAGE_PATH! Used standard: ./storage/ .")
	}

	publicAddress := os.Getenv("PUBLIC_ADDRESS")
	if publicAddress == "" {
		publicAddress = "http://0.0.0.0"
		fmt.Println("Can`t parse PUBLIC_ADDRESS! Used standard: 0.0.0.0 .")
	}

	dustAddress := os.Getenv("DUST_ADDR")
	if dustAddress == "" {
		dustAddress = "http://0.0.0.0:1106"
		fmt.Println("Can`t parse DUST_ADDR! Used standard: http://0.0.0.0:1106.")
	}

	s := NewServer(port, time.Duration(int(time.Duration.Seconds(60))), storagePath, publicAddress, dustAddress)
	s.InitNode()
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

	return f.Listen(":" + s.listenPort)
}

func (s *Server) Health(c *fiber.Ctx) (err error) {
	memory := system.ReadMemoryStats()
	disk := system.ReadDiskInfo(s.StoragePath)

	resp := butterfly.Health{
		Status:          "ok",
		UTC:             time.Now().UTC().String(),
		NodeType:        "stone",
		Version:         "0.1.0",
		FreeMemory:      memory.FreeMem,
		AvailableMemory: memory.AvailableMem,
		FreeDisk:        disk.AvailableDisk,
		AvailableDisk:   disk.AvailableDisk,
	}
	return c.JSON(resp)
}
