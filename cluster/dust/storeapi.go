package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/iamsoloma/butterfly"
)

func (s *Server) Put(c *fiber.Ctx) error {
	query := butterfly.Query{}

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
		if err := s.Carbine.Put(obj.Lib+":"+obj.Key, obj.Value); err != nil {
			return c.Status(http.StatusInternalServerError).Send([]byte(err.Error() + " Object " + fmt.Sprint(i+1)))
		}
	}
	return c.JSON(map[string]string{"status": "ok"})
}

func (s *Server) Get(c *fiber.Ctx) (err error) {
	query := butterfly.Query{}

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
		if query.Objects[i].Value, err = s.Carbine.Get(obj.Lib + ":" + obj.Key); err != nil {
			return c.Status(http.StatusNotFound).Send([]byte(err.Error() + " Object " + fmt.Sprint(i+1)))
		}
	}
	return c.JSON(query)
}

func (s *Server) Update(c *fiber.Ctx) (err error) {
	query := butterfly.Query{}

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
		if err := s.Carbine.Update(obj.Lib+":"+obj.Key, obj.Value); err != nil {
			return c.Status(http.StatusNotFound).Send([]byte(err.Error() + " Object " + fmt.Sprint(i+1)))
		}
	}
	return c.JSON(map[string]string{"status": "ok"})
}

func (s *Server) Delete(c *fiber.Ctx) (err error) {
	query := butterfly.Query{}

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
		if query.Objects[i].Value, err = s.Carbine.Delete(obj.Lib + ":" + obj.Key); err != nil {
			return c.Status(http.StatusNotFound).Send([]byte(err.Error() + " Object " + fmt.Sprint(i+1)))
		}
	}
	return c.JSON(query)
}

func (s *Server) Host(c *fiber.Ctx) (err error) {
	lib, key := c.Params("lib"), c.Params("key")
	value := ""

	//check object
	if lib == "" || key == "" {
		return c.Status(http.StatusBadRequest).Send([]byte("Lib or Key is empty!"))
	}

	//get object
	if value, err = s.Carbine.Get(lib + ":" + key); err != nil {
		return c.Status(http.StatusNotFound).Send([]byte(err.Error()))
	}

	return c.Send([]byte(value))
}

func (s *Server) List(c *fiber.Ctx) (err error) {
	query := butterfly.Query{}
	resp := butterfly.ListResp{
		Lists: []butterfly.List{},
	}

	//parce query
	if err = c.BodyParser(&query); err != nil {
		return c.Status(http.StatusBadRequest).Send([]byte("Can`t parse JSON!"))
	}

	//get lists
	if keys, err := s.Carbine.List(); err != nil {
		return c.Status(http.StatusInternalServerError).Send([]byte(err.Error()))
	} else {
		//sort by prefix
		for i, obj := range query.Objects {
			resp.Lists = append(resp.Lists, butterfly.List{Prefix: obj.Lib + ":" + obj.Key})
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