package main

import (
	"log"
	"net/http"
	"os"

	"github.com/TinajXD/butterfly"
	"github.com/gofiber/fiber/v2"
)

type Server struct {
	listenAddr string
	Dust       butterfly.Storer[string, string]
}

func NewServer(listenAddr string) *Server {
	return &Server{
		listenAddr: listenAddr,
		Dust:       butterfly.NewDustStore[string, string](),
	}
}

func (s *Server) Put(c *fiber.Ctx) error {
	key := c.Params("key")
	value := c.Params("value")

	if err := s.Dust.Put(key, value); err != nil {
		return err
	} else {
		return c.JSON(map[string]string{"status": "ok"})
	}
}

func (s *Server) Get(c *fiber.Ctx) (err error) {
	query := new(GetQuery)
	if err = c.BodyParser(query); err != nil {
		c.Status(http.StatusInternalServerError).Send([]byte("Can`t parse JSON!"))
	}
	//key := c.Params("key")

	if query.Value, err = s.Dust.Get(query.Key); err != nil {
		return err
	} else {
		return c.JSON(map[string]string{"value": query.Value})
	}
}

func (s *Server) Update(c *fiber.Ctx) error {
	key := c.Params("key")
	value := c.Params("value")

	if err := s.Dust.Update(key, value); err != nil {
		return err
	} else {
		return c.JSON(map[string]string{"status": "ok"})
	}
}

func (s *Server) Delete(c *fiber.Ctx) error {
	key := c.Params("key")

	if value, err := s.Dust.Delete(key); err != nil {
		return err
	} else {
		return c.JSON(map[string]string{"status": "ok", "value": value})
	}
}

func (s *Server) Start() error {
	f := fiber.New()

	f.Get("/put/:key/:value", s.Put)
	f.Get("/get/:key", s.Get)
	f.Get("/update/:key/:value", s.Update)
	f.Get("/delete/:key", s.Delete)

	return f.Listen(s.listenAddr)
}

func main() {
	port := os.Getenv("PORT")
	s := NewServer(":" + port)
	log.Fatal(s.Start())
}
