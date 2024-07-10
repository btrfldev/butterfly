package main

import (
	"fmt"
	"os"

	"github.com/iamsoloma/butterfly/store/kvf"
)

func main() {
	KeySpace := []string{"Key", "Space", "Test", "File", ".", "It`s", "work", "!"}

	file, err := os.Create("bloom.kv")
	check(err)
	file.Close()

	WRfile, err := os.OpenFile("bloom.kv", os.O_RDWR, 0666)
	check(err)
	defer WRfile.Close()
	err = kvf.WriteKeySpace(WRfile, KeySpace)
	check(err)


	Rfile, err := os.OpenFile("bloom.kv", os.O_RDONLY, 0666)
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