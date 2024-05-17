package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	//"github.com/TinajXD/butterfly"
	"github.com/TinajXD/butterfly/system"
	"github.com/gofiber/fiber/v2"
)

func main() {
	//var timeout int
	port := os.Getenv("PORT")
	timeout, err := strconv.Atoi(os.Getenv("TIMEOUT"))
	if err != nil {
		timeout = 60
		fmt.Println("Can`t parse timeout! Used standard: 60 sec.")
	}
	bodyLimit, err := strconv.Atoi(os.Getenv("BODYLIMIT"))
	if err != nil {
		bodyLimit = 1024*1024*1024*1024
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

	return f.Listen(s.listenAddr)
}

func (s *Server) Put(c *fiber.Ctx) error {
	query := Query{}

	if err := c.BodyParser(&query); err != nil {
		return c.Status(http.StatusBadRequest).Send([]byte("Can`t parse JSON!"))
	}

	if query.Lib == "" || query.Key == "" {
		return c.Status(http.StatusBadRequest).Send([]byte("Lib or Key is empty!"))
	}

	if err := s.Dust.Put(query.Lib+":"+query.Key, query.Value); err != nil {
		return c.Status(http.StatusInternalServerError).Send([]byte(err.Error()))
	} else {
		return c.JSON(map[string]string{"status": "ok"})
	}
}

func (s *Server) Get(c *fiber.Ctx) (err error) {
	query := Query{}
	if err = c.BodyParser(&query); err != nil {
		return c.Status(http.StatusBadRequest).Send([]byte("Can`t parse JSON!"))
	}

	if query.Lib == "" || query.Key == "" {
		return c.Status(http.StatusBadRequest).Send([]byte("Lib or Key is empty!"))
	}

	if query.Value, err = s.Dust.Get(query.Lib + ":" + query.Key); err != nil {
		return c.Status(http.StatusInternalServerError).Send([]byte(err.Error()))
	} else {
		return c.JSON(map[string]string{"value": query.Value})
	}
}

func (s *Server) Update(c *fiber.Ctx) (err error) {
	query := Query{}

	if err = c.BodyParser(&query); err != nil {
		return c.Status(http.StatusBadRequest).Send([]byte("Can`t parse JSON!"))
	}

	if query.Lib == "" || query.Key == "" {
		return c.Status(http.StatusBadRequest).Send([]byte("Lib or Key is empty!"))
	}

	if err := s.Dust.Update(query.Lib+":"+query.Key, query.Value); err != nil {
		return c.Status(http.StatusInternalServerError).Send([]byte(err.Error()))
	} else {
		return c.JSON(map[string]string{"status": "ok"})
	}
}

func (s *Server) Delete(c *fiber.Ctx) (err error) {
	query := Query{}

	if err = c.BodyParser(&query); err != nil {
		return c.Status(http.StatusBadRequest).Send([]byte("Can`t parse JSON!"))
	}

	if query.Lib == "" || query.Key == "" {
		return c.Status(http.StatusBadRequest).Send([]byte("Lib or Key is empty!"))
	}

	if value, err := s.Dust.Delete(query.Lib + ":" + query.Key); err != nil {
		return err
	} else {
		return c.JSON(map[string]string{"status": "ok", "value": value})
	}
}


func (s *Server) Health(c *fiber.Ctx) (err error) {
	memory := system.ReadMemoryStats()

	resp := Health{
		Status:       "ok",
		UTC:          time.Now().UTC().String(),
		StorerType:   "Dust",
		TotalStorage:  memory.MemTotal,
		AvailableStorage: memory.MemAvailable,
	}
	return c.JSON(resp)
}
