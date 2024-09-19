package main

import (
	"fmt"

	btrflstore "github.com/iamsoloma/butterfly/store"
)

func main() {
	path := "./examples/diskStore.badger"
	kv, err := btrflstore.NewDiskStore(path)
	if err!=nil{
		panic(err)
	}
	defer kv.CloseDiskStore()

	err = kv.Put("Ping", "Pong")
	if err!=nil{
		panic(err)
	}

	value, err := kv.Get("Ping")
	if err!=nil{
		panic(err)
	}
	fmt.Println("Ping" + ":" + value)
}