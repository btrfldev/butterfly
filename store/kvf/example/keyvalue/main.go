package main

import (
	"fmt"
	"os"

	"github.com/btrfldev/butterfly/store/kvf"
)

func main() {
	path := "./examples/keyvalue.kv"

	example := make(map[string]string)
	example["Ping"] = "Pong"
	example["Hi!"] = "Bye..."
	example["Black"] = "White"
	example["Cat"] = "Dog"
	example["Chelyabinsk"] = "Chelyabinsk is the administrative center and largest city Chelyabinsk Oblast, Russia. It is the seventh-largest city in Russia, with a population of over 1.1 million people, and the second-largest city in the Ural Federal District, after Yekaterinburg. Chelyabinsk is located to the East behind the South part of the Ural Mountains and runs along the Miass River."

	WriteKV(path, example)
	result := ReadV(path, []string{"Ping", "Hi!", "Chelyabinsk"})
	for k, v := range result {
		fmt.Printf("%s:%s\n", k, v)
	}
}

func WriteKV(filename string, kv map[string]string) {
	Keys, Values := []string{}, []string{}

	for k, v := range kv {
		Keys = append(Keys, k)
		Values = append(Values, v)
	}

	file, err := os.Create(filename)
	check(err)
	defer file.Close()

	Wfile, err := os.OpenFile(filename, os.O_WRONLY, 0666)
	check(err)
	defer Wfile.Close()

	kvf.WriteKeySpace(Wfile, Keys)
	check(err)

	AWfile, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0666)
	check(err)
	defer AWfile.Close()
	Rfile, err := os.OpenFile(filename, os.O_RDONLY, 0666)
	check(err)
	defer Rfile.Close()

	_, err = kvf.AppendValues(AWfile, Rfile, Values)
	check(err)
	//fmt.Println("Last Appended value: " + strconv.Itoa(last))
}

func ReadV(filename string, keys []string) (result map[string]string) {
	Rfile, err := os.OpenFile(filename, os.O_RDONLY, 0666)
	check(err)
	defer Rfile.Close()

	R2file, err := os.OpenFile(filename, os.O_RDONLY, 0666)
	check(err)
	defer R2file.Close()

	result, err = kvf.ReadValues(Rfile, R2file, keys)
	check(err)
	return result
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
