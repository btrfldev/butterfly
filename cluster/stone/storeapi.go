package main

import (
	"encoding/json"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/btrfldev/butterfly"
	"github.com/gofiber/fiber/v2"
)

func (s *Server) UploadFromForm(c *fiber.Ctx) error {
	inpName := c.Params("inpname")

	file, err := c.FormFile(inpName)
	if err != nil {
		return c.Status(http.StatusBadRequest).Send([]byte("Can`t find file!"))
	}

	lib := c.Params("lib")
	key := c.Params("*")
	key += "/" + file.Filename
	keyLocal := strings.ReplaceAll(key, "/", "^")

	_, err = os.Stat(s.StoragePath + "/" + lib)
	if os.IsNotExist(err) {
		println(s.StoragePath + lib + "/")
		if err := os.MkdirAll(s.StoragePath+lib+"/", 0777); err != nil {
			return c.Status(http.StatusInternalServerError).Send([]byte("Can`t create a lib`s dir!\n" + err.Error()))
		}
	}

	println(s.StoragePath + lib + "/" + key)
	err = c.SaveFile(file, s.StoragePath+lib+"/"+keyLocal)
	if err != nil {
		return c.Status(http.StatusInternalServerError).Send([]byte("Can`t save file!\n" + err.Error()))
	}

	cl := fiber.Client{}
	agent := cl.Get(s.DustAddress + "/put")
	agent.JSON(butterfly.Query{Objects: []butterfly.Object{
		{
			Lib:   "stone_obj_" + lib,
			Key:   key,
			Value: strconv.Itoa(s.NodeInfo.ID),
		},
	}})
	statusCode, _, errs := agent.Bytes()
	if len(errs) > 0 {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"errs": errs,
		})
	}
	if statusCode == 200 {
		//println(string(body[:]))
		return c.JSON(butterfly.Object{
			Lib:   lib,
			Key:   key,
			Value: strconv.FormatInt(file.Size, 10) + " " + strconv.Itoa(s.NodeInfo.ID),
		})
	} else {
		return c.Status(http.StatusInternalServerError).Send([]byte("Can`t save meta data!\n"))
	}

}

func (s *Server) Get(c *fiber.Ctx) error {
	lib := c.Params("lib")
	key := c.Params("*")

	if lib == "" || key == "" {
		return c.Status(http.StatusBadRequest).Send([]byte("Can`t parse key or lib!"))
	}

	cl := fiber.Client{}
	agent := cl.Get(s.DustAddress + "/get")
	println("stone_obj_" + lib + " | " + strings.ReplaceAll(key, "%20", " "))
	agent.JSON(butterfly.Query{Objects: []butterfly.Object{
		{
			Lib:   "stone_obj_" + lib,
			Key:   strings.ReplaceAll(key, "%20", " "),
			Value: "",
		},
	}})
	statusCode, body, errs := agent.Bytes()
	if len(errs) > 0 {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"errs": errs,
		})
	}
	resp := butterfly.Query{}
	if statusCode == 200 {
		if err := json.Unmarshal(body, &resp); err != nil {
			return c.Status(http.StatusInternalServerError).Send([]byte("Can`t unmarshal database`s responce."))
		}

		if resp.Objects[0].Value == strconv.Itoa(s.NodeInfo.ID) {
			println(resp.Objects[0].Value + " | " + strconv.Itoa(s.NodeInfo.ID))
			return c.SendFile(s.StoragePath + lib + "/" + strings.ReplaceAll(key, "/", "^"))
		} else {
			//TODO: Redirect by node id
			return c.Redirect(resp.Objects[0].Value+"/store/"+lib+"/"+key, http.StatusContinue)
		}

	} else {
		return c.Status(http.StatusNotFound).Send([]byte("Can`t found object`s meta!\n" + lib + ":" + key + "\n" + string(body[:])))
	}

}
