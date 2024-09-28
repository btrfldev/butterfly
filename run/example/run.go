package main

import (
	"errors"
	"os"
	"strings"

	"github.com/iamsoloma/butterfly/logger"
	btrflrun "github.com/iamsoloma/butterfly/run"
)

func main() {
	container := btrflrun.Container{
		Image:       "redis",
		Name:        "redis",
		CallCommand: "redis-server",
		WorkDir:     "/usr/local/bin",
		Logger: logger.Logger{
			LogToTerminal: true,
		},
	}
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	dir := path + "/containers/" + container.Name
	err = os.MkdirAll(dir, 0777)
	if err != nil {
		panic(err)
	}
	defer os.RemoveAll(dir)
	//println(dir)
	container.ContainerRootDir = dir

	localImage := strings.ReplaceAll(container.Image, "/", ":")
	if _, err := os.Stat(localImage); errors.Is(err, os.ErrNotExist) {
		container.PullDockerImage("./images")
	}

	err = btrflrun.UnTar("./images/"+localImage+".tar.gz", container.ContainerRootDir)
	if err != nil {
		panic(err)
	}

	err = container.Run()
	if err != nil {
		panic(err)
	}

}
