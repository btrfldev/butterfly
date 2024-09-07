package main

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/iamsoloma/butterfly"
)

func (s *Server) Upload(c *fiber.Ctx) error {
	inpName := c.Params("inpname")

	file, err := c.FormFile(inpName)
	if err != nil {
		return c.Status(http.StatusBadRequest).Send([]byte("Can`t find file!"))
	}

	lib := c.Params("lib")
	key := c.Params("*")
	key = strings.ReplaceAll(key, "/", "^")
	key += "^" + file.Filename

	err = c.SaveFile(file, "./tmp/uploads/"+lib+"$"+key)
	if err != nil {
		return c.Status(http.StatusInternalServerError).Send([]byte("Can`t save file!"))
	}
	println("./tmp/uploads/" + lib + "$" + key)

	return c.JSON(butterfly.Object{
		Lib:   lib,
		Key:   strings.ReplaceAll(key, "^", "/"),
		Value: strconv.FormatInt(file.Size, 10),
	})
}
