package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/iamsoloma/butterfly"
	"github.com/iamsoloma/butterfly/system"
	"github.com/gofiber/fiber/v2"
)

func main() {
	port := os.Getenv("PORT")
	timeout, err := strconv.Atoi(os.Getenv("TIMEOUT"))
	if err != nil {
		timeout = 60
		fmt.Println("Can`t parse timeout! Used standard: 60 sec.")
	}
	bodyLimit, err := strconv.Atoi(os.Getenv("BODYLIMIT"))
	if err != nil {
		bodyLimit = 1024 * 1024 * 1024 * 1024
		fmt.Println("Can`t parse BodyLimit! Used standard: 1024*1024*1024*1024.")
	}
	s := NewServer(":"+port, bodyLimit, time.Duration(timeout))
	log.Fatal(s.Start())
}

func (s *Server) Start() error {
	f := fiber.New(
		fiber.Config{
			BodyLimit:   s.bodyLimit,
			IdleTimeout: s.idleTimeout,
		},
	)

	f.Get("/health", s.Health)
	f.Get("/put", s.Put)
	f.Get("/get", s.Get)
	f.Get("/update", s.Update)
	f.Get("/delete", s.Delete)
	f.Get("/list", s.List)

	return f.Listen(s.listenAddr)
}


func (s *Server) Health(c *fiber.Ctx) (err error) {
	memory := system.ReadMemoryStats()

	resp := butterfly.Health{
		Status:           "ok",
		UTC:              time.Now().UTC().String(),
		StorerType:       "Dust",
		Version:          "0.2.3",
		TotalStorage:     memory.MemTotal,
		AvailableStorage: memory.MemAvailable,
	}
	return c.JSON(resp)
}
