package main

import (
	"fmt"
	"os"

	"github.com/iamsoloma/butterfly/store/kvf"
)

func main() {
	path := "./examples/keyspace.kv"
	KeySpace := []string{"Key", "Space", "Test", "File", ".", "It`s", "work", "!"}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		file, err := os.Create(path)
		check(err)
		file.Close()
	}

	WRfile, err := os.OpenFile(path, os.O_RDWR, 0666)
	check(err)
	defer WRfile.Close()
	err = kvf.WriteKeySpace(WRfile, KeySpace)
	check(err)


	Rfile, err := os.OpenFile(path, os.O_RDONLY, 0666)
	check(err)
	defer Rfile.Close()
	KeySpace, err = kvf.GetKeySpace(Rfile)
	check(err)
	fmt.Println(KeySpace)

}

func check(err error) {
	if err!=nil{
		panic(err)
	}
}