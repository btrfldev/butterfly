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
	query := Query{}
	/*key := c.Params("key")
	value := c.Params("value")*/
	if err := c.BodyParser(&query); err != nil {
		return c.Status(http.StatusBadRequest).Send([]byte("Can`t parse JSON!"))
	}

	if err := s.Dust.Put(query.Key, query.Value); err != nil {
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

	if query.Value, err = s.Dust.Get(query.Key); err != nil {
		return c.Status(http.StatusInternalServerError).Send([]byte(err.Error()))
	} else {
		return c.JSON(map[string]string{"value": query.Value})
	}
}

func (s *Server) Update(c *fiber.Ctx) (err error) {
	query := Query{}
	/*key := c.Params("key")
	value := c.Params("value")*/
	if err = c.BodyParser(&query); err != nil {
		return c.Status(http.StatusBadRequest).Send([]byte("Can`t parse JSON!"))
	}

	if err := s.Dust.Update(query.Key, query.Value); err != nil {
		return c.Status(http.StatusInternalServerError).Send([]byte(err.Error()))
	} else {
		return c.JSON(map[string]string{"status": "ok"})
	}
}

func (s *Server) Delete(c *fiber.Ctx) (err error) {
	//key := c.Params("key")
	query:=Query{}

	if err = c.BodyParser(&query); err != nil {
		return c.Status(http.StatusBadRequest).Send([]byte("Can`t parse JSON!"))
	}

	if value, err := s.Dust.Delete(query.Key); err != nil {
		return err
	} else {
		return c.JSON(map[string]string{"status": "ok", "value": value})
	}
}

func (s *Server) Start() error {
	f := fiber.New()

	f.Get("/put", s.Put)
	f.Get("/get", s.Get)
	f.Get("/update", s.Update)
	f.Get("/delete", s.Delete)

	return f.Listen(s.listenAddr)
}

func main() {
	port := os.Getenv("PORT")
	s := NewServer(":" + port)
	log.Fatal(s.Start())
}
