package run

import (
	"archive/tar"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"syscall"

	"github.com/iamsoloma/butterfly/logger"
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

	cmd := exec.Command(c.CallCommand)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	println(c.ContainerRootDir)
	err := syscall.Chdir(c.ContainerRootDir)
	if err != nil {
		c.Logger.Error(err, errLocation, "Can`t change dir")
	}
	err = syscall.Chroot(c.ContainerRootDir)
	if err != nil {
		c.Logger.Error(err, errLocation, "Can`t change root")
	}

	err = syscall.Chdir(c.WorkDir)
	if err != nil {
		c.Logger.Error(err, errLocation, "Can`t change dir")
	}

	dirs, err := os.ReadDir("./")
	if err != nil {
		c.Logger.Error(err, errLocation, "Can`t look workdir")
	}
	for _, dir := range dirs {
		println(dir.Name())
	}
	/*body, err := os.ReadFile("/usr/local/bin/docker-entrypoint.sh")
	if err != nil {
		c.Logger.Error(err, errLocation, "Can`t read call command in container")
	}
	println(string(body))*/

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
	//cmd.Dir = "./"
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	//print(cmd.Output())
	err := cmd.Run()
	if err != nil {
		c.Logger.Error(err, errLocation, "Can`t pull docker image")
	}

}

func UnTar(tarball, target string) error {
	reader, err := os.Open(tarball)
	if err != nil {
		return err
	}
	defer reader.Close()
	tarReader := tar.NewReader(reader)

	for {
		header, err := tarReader.Next()
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}

		path := filepath.Join(target, header.Name)
		info := header.FileInfo()
		if info.IsDir() {
			if err = os.MkdirAll(path, info.Mode()); err != nil {
				return err
			}
			continue
		}

		file, err := os.OpenFile(path, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, info.Mode())
		if err != nil {
			return err
		}
		defer file.Close()
		_, err = io.Copy(file, tarReader)
		if err != nil {
			return err
		}
	}
	return nil
}
