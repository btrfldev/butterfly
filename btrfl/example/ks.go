package main

import (
	"fmt"
	"os"

	"github.com/iamsoloma/butterfly/btrfl"
)

func main() {
	KeySpace := []string{"Key", "Space", "Test", "File", ".", "It`s", "work", "!"}

	file, err := os.Create("bloom.btrfl")
	check(err)
	file.Close()

	WRfile, err := os.OpenFile("bloom.btrfl", os.O_RDWR, 0666)
	check(err)
	defer WRfile.Close()
	err = btrfl.WriteKeySpace(WRfile, KeySpace)
	check(err)


	Rfile, err := os.OpenFile("bloom.btrfl", os.O_RDONLY, 0666)
	check(err)
	defer Rfile.Close()
	KeySpace, err = btrfl.GetKeySpace(Rfile)
	check(err)
	fmt.Println(KeySpace)

}

func check(err error) {
	if err!=nil{
		panic(err)
	}
}