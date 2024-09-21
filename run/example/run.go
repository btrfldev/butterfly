package main

import (
	"os"
	"strings"

	"github.com/iamsoloma/butterfly/logger"
	btrflrun "github.com/iamsoloma/butterfly/run"
)

func main() {
	container := btrflrun.Container{
		Image:       "jupyter/datascience-notebook",
		Name:        "way",
		CallCommand: "./start-notebook.py",
		WorkDir:     "/",
		Logger: logger.Logger{
			LogToTerminal: true,
		},
	}
	container.PullDockerImage("./images")
	dir, err := os.MkdirTemp("", container.Name)
	if err != nil {
		panic(err)
	}
	defer os.RemoveAll(dir)
	//println(dir)
	container.ContainerRootDir = dir

	localImage := strings.ReplaceAll(container.Image, "/", ":")
	err = btrflrun.UnTar("./images/"+localImage+".tar.gz", container.ContainerRootDir)
	if err != nil {
		panic(err)
	}

	err = container.Run()
	if err != nil {
		panic(err)
	}

}
