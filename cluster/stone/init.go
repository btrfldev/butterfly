package main

import (
	"encoding/json"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/iamsoloma/butterfly"
)

func (s *Server) InitNode() {
	s.initStorage()

	s.registerNewNode()
}

func (s *Server) initStorage() {
	if err := os.MkdirAll(s.StoragePath, 0777); err != nil {
		panic(err)
	}
}

func (s *Server) registerNewNode() {
	cl := fiber.Client{}
	resp := butterfly.ListResp{}
	agent := cl.Get(s.DustAddress + "/list")
	agent.JSON(butterfly.Query{
		Objects: []butterfly.Object{
			{
				Lib:   "stone_nodes",
				Key:   "",
				Value: "",
			},
		},
	})
	statusCode, body, errs := agent.Bytes()
	if len(errs) > 0 {
		for _, err := range errs {
			fmt.Println(err.Error())
		}
		panic("Can`t get nodeID!")
	}
	if statusCode == 200 {
		if err := json.Unmarshal(body, &resp); err != nil {
			panic(err)
		}
	} else {
		panic("Can`t get list of nodes.")
	}

	NewID := 0
	if resp.Lists[0].Count > 0 {
		NewID = 0
		IDs := []int{}
		for _, node := range resp.Lists[0].Keys {
			lk := strings.Split(node, ":")
			NextID, err := strconv.Atoi(lk[1])
			if err != nil {
				panic("can`t parse node`s id: " + lk[1])
			}
			IDs = append(IDs, NextID)
		}
		slices.Sort(IDs)

		if len(IDs) == 1 {
			for IDs[0] >= NewID {
				NewID += 1
			}
		} else {
			for in := range IDs {
				if in == 0 {
					if NewID < IDs[in] {
						break
					} else {
						NewID += 1
					}
				} else if IDs[in-1] < NewID && NewID < IDs[in] {
					break
				} else {
					NewID += 1
				}
			}
		}
	} else {
		NewID = 0
	}

	agent = cl.Get(s.DustAddress + "/put")
	agent.JSON(butterfly.Query{
		Objects: []butterfly.Object{
			{
				Lib:   "stone_nodes",
				Key:   strconv.Itoa(NewID),
				Value: s.NodeInfo.PublicAddress + ":" + s.listenPort,
			},
		},
	})
	statusCode, body, errs = agent.Bytes()
	if len(errs) > 0 {
		for _, err := range errs {
			fmt.Println(err.Error())
		}
		panic("Can`t put nodeID!")
	}
	if statusCode == 200 {
		if err := json.Unmarshal(body, &resp); err != nil {
			panic(err)
		}
		s.NodeInfo.ID = NewID
	} else {
		panic("Can`t get list of nodes.")
	}

	
}
