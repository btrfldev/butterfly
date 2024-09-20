package main

import (
	"fmt"
	"strings"

	btrflstore "github.com/iamsoloma/butterfly/store"
)

func main() {
	path := "./examples/diskStore.badger"
	example := make(map[string]string)
	example["Ping"] = "Pong"
	example["Hi!"] = "Bye..."
	example["Black"] = "White"
	example["Cat"] = "Dog"
	example["Chelyabinsk"] = "Chelyabinsk is the administrative center and largest city Chelyabinsk Oblast, Russia. It is the seventh-largest city in Russia, with a population of over 1.1 million people, and the second-largest city in the Ural Federal District, after Yekaterinburg. Chelyabinsk is located to the East behind the South part of the Ural Mountains and runs along the Miass River."
	
	kv, err := btrflstore.NewDiskStore(path)
	if err!=nil{
		panic(err)
	}
	defer kv.CloseDiskStore()

	err = kv.Put(example)
	if err!=nil{
		panic(err)
	}
	fmt.Println("Example are putted.")

	res, err := kv.Get([]string{"Ping", "Chelyabinsk"})
	if err!=nil{
		panic(err)
	}
	fmt.Println("Get:")
	printRes(res)

	res, err = kv.Delete([]string{"Cat"})
	if err!=nil{
		panic(err)
	}
	fmt.Println("Deleted:")
	printRes(res)

	list, err := kv.List(func(k, c string) bool {
		if strings.HasPrefix(k, c) {
			return true
		} else {
			return false
		}
	}, "")
	if err!=nil{
		panic(err)
	}

	fmt.Println("List of current keys: ", list)
}

func printRes(res map[string]string) {
	for k,v :=range res{
		fmt.Println("	"+k + " : " + v)
	}
}