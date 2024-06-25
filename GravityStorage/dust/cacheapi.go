package main

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
)


func (s *Server) CacheURL(c *fiber.Ctx) (err error) {
	link := c.Params("+")
	value := ""

	if link == "" {
		return c.Status(http.StatusBadRequest).Send([]byte("Link is empty!"))
	}

	if value, err = s.CacheStorage.Get("URLsData" + ":" + link); err != nil {
		resp, err := http.Get("https://"+link)
		if err != nil {
			resp, err = http.Get("http://"+link)
			if err != nil {
				return c.Status(http.StatusInternalServerError).Send([]byte("Can`t make cache request!\n" + err.Error()))
			}
		}
		defer resp.Body.Close()
		if resp.StatusCode != http.StatusOK {
			return c.Status(http.StatusNoContent).Send([]byte("Bad server responce! " + strconv.Itoa(resp.StatusCode)))
		}

		header := resp.Header.Clone()
		jsonHeader, err := json.Marshal(&header) 
		if err!=nil{
			return c.Status(http.StatusInternalServerError).Send([]byte("Can`t read headers of server responce!"))
		}
		bvalue, err := io.ReadAll(resp.Body)
		if err != nil {
			return c.Status(http.StatusInternalServerError).Send([]byte("Can`t read server responce!"))
		}
		deathTime := time.Now().Add(s.cacheLifeTime)

		if err = s.CacheStorage.Put("URLsData"+":"+link, string(bvalue)); err != nil {
			return c.Status(http.StatusInternalServerError).Send([]byte("Can`t cache responce!"))
		}
		if err = s.CacheStorage.Put("URLsHeader"+":"+link, string(jsonHeader)); err != nil {
			return c.Status(http.StatusInternalServerError).Send([]byte("Can`t cache responce!"))
		}
		if err = s.CacheStorage.Put("URLsDeathTime"+":"+link, deathTime.String()); err != nil {
			return c.Status(http.StatusInternalServerError).Send([]byte("Can`t save death time of cache!"))
		}
		value = string(bvalue)
	}
	header, err := s.CacheStorage.Get("URLsHeader"+":"+link)
	if err!=nil{
		return c.Status(http.StatusInternalServerError).Send([]byte("Can`t read cached header!"))
	}
	var respHeader http.Header
	if json.Unmarshal([]byte(header), &respHeader);err!=nil{
		return c.Status(http.StatusInternalServerError).Send([]byte("Can`t decode cached header!"))
	}

	for k, vv := range respHeader{
		v := strings.Join(vv, "/")
		c.Set(k, v)
	}
	return c.Send([]byte(value))
}
