package run

import (
	"context"
	"os"
	"os/exec"
	"syscall"

	"github.com/iamsoloma/butterfly/logger"

	"github.com/codeclysm/extract"
)

type Container struct {
	Name             string
	Image            string
	WorkDir          string
	CallCommand      string
	ContainerRootDir string
	Logger           logger.Logger
}

func (c *Container) Run() error {
	errLocation := "btrfl.run.Container.Run"

	println(c.ContainerRootDir)
	err := syscall.Chdir(c.ContainerRootDir)
	if err != nil {
		c.Logger.Error(err, errLocation, "Can`t change dir")
	}

	err = syscall.Chroot(c.ContainerRootDir)
	if err != nil {
		c.Logger.Error(err, errLocation, "Can`t change root")
	}

	cmd := exec.Command(c.CallCommand)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = syscall.Chdir(c.WorkDir)
	if err != nil {
		c.Logger.Error(err, errLocation, "Can`t change dir")
	}

	println(c.CallCommand)
	err = cmd.Run()
	if err != nil {
		c.Logger.Error(err, errLocation, "Can`t execute call command in container")
	}

	return nil
}

func (c *Container) PullDockerImage(path string) {
	errLocation := "btrfl.run.Container.PullDockerImage"

	cmd := exec.Command("./pull", c.Image)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		c.Logger.Error(err, errLocation, "Can`t pull docker image")
	}

}

func UnTar(tarball string, target string) error {
	// fmt.Printf("Extracting %s %s\n", source, dst)
	r, err := os.Open(tarball)
	if err != nil {
		return err
	}
	defer r.Close()

	ctx := context.Background()
	return extract.Archive(ctx, r, target, nil)
}