package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/TinajXD/butterfly/system"
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

func (s *Server) Put(c *fiber.Ctx) error {
	query := Query{}

	//parse query
	if err := c.BodyParser(&query); err != nil {
		return c.Status(http.StatusBadRequest).Send([]byte("Can`t parse JSON!"))
	}

	//check all objects
	for i, obj := range query.Objects {
		if obj.Lib == "" || obj.Key == "" {
			return c.Status(http.StatusBadRequest).Send([]byte("Lib or Key is empty! Object " + fmt.Sprint(i+1)))
		}

		if strings.Contains(obj.Lib, ":") || strings.Contains(obj.Key, ":") {
			return c.Status(http.StatusBadRequest).Send([]byte("You can`t use ':' in Lib or Key! Object " + fmt.Sprint(i+1)))
		}
	}

	//put all objects
	for i, obj := range query.Objects {
		if err := s.Dust.Put(obj.Lib+":"+obj.Key, obj.Value); err != nil {
			return c.Status(http.StatusInternalServerError).Send([]byte(err.Error() + " Object " + fmt.Sprint(i+1)))
		}
	}
	return c.JSON(map[string]string{"status": "ok"})
}

func (s *Server) Get(c *fiber.Ctx) (err error) {
	query := Query{}

	//parse query
	if err = c.BodyParser(&query); err != nil {
		return c.Status(http.StatusBadRequest).Send([]byte("Can`t parse JSON!"))
	}

	//check all objects
	for i, obj := range query.Objects {
		if obj.Lib == "" || obj.Key == "" {
			return c.Status(http.StatusBadRequest).Send([]byte("Lib or Key is empty! Object " + fmt.Sprint(i+1)))
		}
	}

	//get all objects
	for i, obj := range query.Objects {
		if query.Objects[i].Value, err = s.Dust.Get(obj.Lib + ":" + obj.Key); err != nil {
			return c.Status(http.StatusInternalServerError).Send([]byte(err.Error() + " Object " + fmt.Sprint(i+1)))
		}
	}
	return c.JSON(query)
}

func (s *Server) Update(c *fiber.Ctx) (err error) {
	query := Query{}

	//parse query
	if err = c.BodyParser(&query); err != nil {
		return c.Status(http.StatusBadRequest).Send([]byte("Can`t parse JSON!"))
	}

	//check all objects
	for i, obj := range query.Objects {
		if obj.Lib == "" || obj.Key == "" {
			return c.Status(http.StatusBadRequest).Send([]byte("Lib or Key is empty!" + " Object " + fmt.Sprint(i+1)))
		}
	}

	//update all objects
	for i, obj := range query.Objects {
		if err := s.Dust.Update(obj.Lib+":"+obj.Key, obj.Value); err != nil {
			return c.Status(http.StatusInternalServerError).Send([]byte(err.Error() + " Object " + fmt.Sprint(i+1)))
		}
	}
	return c.JSON(map[string]string{"status": "ok"})
}

func (s *Server) Delete(c *fiber.Ctx) (err error) {
	query := Query{}

	//parse query
	if err = c.BodyParser(&query); err != nil {
		return c.Status(http.StatusBadRequest).Send([]byte("Can`t parse JSON!"))
	}

	//check all objects
	for i, obj := range query.Objects {
		if obj.Lib == "" || obj.Key == "" {
			return c.Status(http.StatusBadRequest).Send([]byte("Lib or Key is empty!" + " Object " + fmt.Sprint(i+1)))
		}
	}

	//delete all objects
	for i, obj := range query.Objects {
		if query.Objects[i].Value, err = s.Dust.Delete(obj.Lib + ":" + obj.Key); err != nil {
			return c.Status(http.StatusInternalServerError).Send([]byte(err.Error() + " Object " + fmt.Sprint(i+1)))
		}
	}
	return c.JSON(query)
}

func (s *Server) List(c *fiber.Ctx) (err error) {
	query := Query{}
	resp := ListResp{
		Lists: []List{},
	}

	//parce query
	if err = c.BodyParser(&query); err != nil {
		return c.Status(http.StatusBadRequest).Send([]byte("Can`t parse JSON!"))
	}

	//get lists
	if keys, err := s.Dust.List(); err != nil {
		return c.Status(http.StatusInternalServerError).Send([]byte(err.Error()))
	} else {
		//sort by prefix
		for i, obj := range query.Objects {
			resp.Lists = append(resp.Lists, List{Prefix: obj.Lib + ":" + obj.Key})
			for _, key := range keys {
				if strings.HasPrefix(key, obj.Lib+":"+obj.Key) {
					resp.Lists[i].Count += 1
					resp.Lists[i].Keys = append(resp.Lists[i].Keys, key)
				}
			}
		}

		return c.JSON(resp)
	}
}

func (s *Server) Health(c *fiber.Ctx) (err error) {
	memory := system.ReadMemoryStats()

	resp := Health{
		Status:           "ok",
		UTC:              time.Now().UTC().String(),
		StorerType:       "Dust",
		Version:          "0.2.3",
		TotalStorage:     memory.MemTotal,
		AvailableStorage: memory.MemAvailable,
	}
	return c.JSON(resp)
}
