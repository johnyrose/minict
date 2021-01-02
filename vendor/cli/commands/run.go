package commands

import (
	"fmt"
	"oci"
	"run"
	"strings"

	"github.com/apex/log"
)

func Run(imagesDir string, containerDir string, image string, name string) {
	imageName, imageTag := parseImageName(image)
	if !doesContainerExist(name, containerDir) {
		err := oci.UnpackImage(imagesDir, containerDir, name, imageName, imageTag)
		if err != nil {
			log.Error(fmt.Sprintf("Failed to unpack image. Error received: %s", err.Error()))
		}
	}
	err := run.RunContainer(imagesDir, containerDir, name)
	if err != nil {
		log.Error(fmt.Sprintf("Failed to run image. Error received: %s", err.Error()))
	}
}

func parseImageName(name string) (string, string) {
	split := strings.Split(name, ":")
	if len(split) != 2 {
		log.Fatal("Invalid image name")
	}
	return split[0], split[1]
}

func doesContainerExist(name string, containerDir string) bool {
	for _, container := range ListContainers(containerDir) {
		if name == container {
			return true
		}
	}
	return false
}
