package run

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"syscall"
)

func RunContainer(imagesDir string, containersDir string, name string, image string) {
	// imageName, imageTag := parseImageName(image)
	// oci.UnpackImage(imagesDir, containersDir, name, imageName, imageTag)
	containerDir :""= containersDir + "/" + name
	imageConfig := GetImageConfig(containerDir)
	os.Chdir(containerDir)
	cmd = buildCommand(imageConfig)
	applyNamespaces(cmd)
	
}

func parseImageName(name string) (string, string) {
	split := strings.Split(name, ":")
	if len(split) != 2 {
		log.Fatal("Invalid image name")
	}
	return split[0], split[1]
}

func buildCommand(imageConfig ImageConfig) *exec.Cmd {
	cmd := exec.Command(imageConfig.ProcessConfig.Args[0], imageConfig.ProcessConfig.Args[1:]...)
	// TODO: Add env variables and changing of working directory
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd
}

func applyNamespaces(cmd *exec.Cmd) {
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS
	}
}
