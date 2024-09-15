package main

import (
	"fmt"

	"github.com/iamsoloma/butterfly"
	"github.com/iamsoloma/butterfly/cluster/dust/client"
)

func main() {
	agent := client.Agent{
		DustAddress: "http://0.0.0.0:1106",
	}

	Objects := []butterfly.Object{
		{
			Lib:   "example",
			Key:   "123",
			Value: "321",
		},
		{
			Lib: "example",
			Key: "Hello, Dust!",
		    Value: "Hi!",
		},
	}

	if err:=agent.Put(Objects);err!=nil{
		panic(err)
	} else {
		fmt.Println("Objects` are putted!")
	}
}