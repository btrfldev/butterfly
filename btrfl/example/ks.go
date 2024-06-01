package main

import (
	"fmt"
	"os"

	"github.com/iamsoloma/butterfly/btrfl"
)

func main() {
	KeySpace := []string{"Petya", "Egor", "Vasya", "Olga", "Sveta"}
	
	/*Wfile, err := os.OpenFile("bloom.btrfl", os.O_WRONLY, 0666)
	check(err)
	defer Wfile.Close()
	Rfile, err := os.OpenFile("bloom.btrfl", os.O_RDONLY, 0666)
	check(err)
	defer Rfile.Close()*/

	WRfile, err := os.OpenFile("bloom.btrfl", os.O_RDWR, 0666)
	check(err)
	defer WRfile.Close()
	err = btrfl.UpdateKeySpace(WRfile, KeySpace)
	check(err)


	Rfile, err := os.OpenFile("bloom.btrfl", os.O_RDONLY, 0666)
	check(err)
	defer Rfile.Close()
	KeySpace, err = btrfl.GetKeySpace(Rfile/*"./bloom.btrfl"*/)
	check(err)
	fmt.Println(KeySpace)

}

func check(err error) {
	if err!=nil{
		panic(err)
	}
}