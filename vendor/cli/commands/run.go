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
	err := oci.UnpackImage(imagesDir, containerDir, name, imageName, imageTag)
	if err != nil {
		log.Error(fmt.Sprintf("Failed to unpack image. Error received: %s", err.Error()))
	}
	run.RunContainer(imagesDir, containerDir, name, image)
}

func parseImageName(name string) (string, string) {
	split := strings.Split(name, ":")
	if len(split) != 2 {
		log.Fatal("Invalid image name")
	}
	return split[0], split[1]
}
