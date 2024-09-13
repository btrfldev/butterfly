package main

import (
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/iamsoloma/butterfly"
)

func (s *Server) UploadFromForm(c *fiber.Ctx) error {
	inpName := c.Params("inpname")

	file, err := c.FormFile(inpName)
	if err != nil {
		return c.Status(http.StatusBadRequest).Send([]byte("Can`t find file!"))
	}

	lib := c.Params("lib")
	key := c.Params("*")
	key = strings.ReplaceAll(key, "/", "^")
	key += "^" + file.Filename

	dir, err := os.Stat(s.StoragePath + "/" + lib)
	if os.IsNotExist(err) {
		if err := os.MkdirAll(s.StoragePath+lib, 0666); err != nil {
			return c.Status(http.StatusInternalServerError).Send([]byte("Can`t create a lib`s dir!\n" + err.Error()))
		}
	}
	if !dir.IsDir() {
		return c.Status(http.StatusInternalServerError).Send([]byte("Can`t create a dir! Dir isn`t a dir."))
	}

	//println(s.StoragePath+lib+"/"+key)
	err = c.SaveFile(file, s.StoragePath+lib+"/"+key)
	if err != nil {
		return c.Status(http.StatusInternalServerError).Send([]byte("Can`t save file!\n" + err.Error()))
	}

	cl := fiber.Client{}
	agent := cl.Get(s.DustAddress)
	agent.JSON(butterfly.Object{
		Lib:   lib,
		Key:   key,
		Value: "local",
	})
	statusCode, _, errs := agent.Bytes()
	if len(errs) > 0 {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"errs": errs,
		})
	}
	if statusCode==200{
		return c.JSON(butterfly.Object{
			Lib:   lib,
			Key:   strings.ReplaceAll(key, "^", "/"),
			Value: strconv.FormatInt(file.Size, 10),
		})
	} else {
		return c.Status(http.StatusInternalServerError).Send([]byte("Can`t save meta data!"))
	}
	

}

func (s *Server) Get(c *fiber.Ctx) error {
	lib := c.Params("lib")
	key := c.Params("*")
	key = strings.ReplaceAll(key, "/", "^")

	if lib == "" || key == "" {
		return c.Status(http.StatusBadRequest).Send([]byte("Can`t parse key or lib!"))
	}

	return c.SendFile(s.StoragePath + lib + "/" + key)
}
