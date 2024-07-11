package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/iamsoloma/butterfly"
	"github.com/iamsoloma/butterfly/system"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "1106"
		fmt.Println("Can`t parse port! Used standard: 1106.")
	}

	timeout, err := strconv.Atoi(os.Getenv("TIMEOUT"))
	if err != nil {
		timeout = int(time.Duration.Seconds(60))
		fmt.Println("Can`t parse timeout! Used standard: 60 sec.")
	}

	bodyLimit, err := strconv.Atoi(os.Getenv("BODYLIMIT"))
	if err != nil {
		bodyLimit = 1024 * 1024 * 1024 * 1024
		fmt.Println("Can`t parse BodyLimit! Used standard: 1024*1024*1024*1024.")
	}

	var cacheLifeTime time.Duration
	cacheLifeTimeInp, err := strconv.Atoi(os.Getenv("cacheLifeTime"))
	if err != nil {
		cacheLifeTime = 1 * time.Hour
		fmt.Println("Can`t parse cacheLifeTime! Used standard: 1(hour).")
	} else {
		cacheLifeTime = time.Duration(cacheLifeTimeInp) * time.Hour
	}

	s := NewServer(":"+port, bodyLimit, time.Duration(timeout), cacheLifeTime)
	log.Fatal(s.Start())
}

func (s *Server) Start() error {
	f := fiber.New(
		fiber.Config{
			BodyLimit:   s.bodyLimit,
			IdleTimeout: s.idleTimeout,
			Prefork:     true,
		},
	)

	//main
	f.Get("/health", s.Health)

	//storeapi
	f.Get("/put", s.Put)
	f.Get("/get", s.Get)
	f.Get("/update", s.Update)
	f.Get("/delete", s.Delete)
	f.Get("/list", s.List)
	f.Get("/host/:lib/:key", s.Host)

	//cacheapi
	f.Get("/cacheurl/+", s.CacheURL)
	f.Get("/+", s.CacheURL)

	return f.Listen(s.listenAddr)
}

func (s *Server) Health(c *fiber.Ctx) (err error) {
	memory := system.ReadMemoryStats()

	resp := butterfly.Health{
		Status:           "ok",
		UTC:              time.Now().UTC().String(),
		NodeType:         "dust",
		Version:          "0.2.3",
		TotalStorage:     memory.MemTotal,
		AvailableStorage: memory.MemAvailable,
	}
	return c.JSON(resp)
}
