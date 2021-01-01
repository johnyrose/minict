package runtime

import (
	"log"
	"oci"
	"strings"
)

func RunContainer(imagesDir string, containersDir string, name string, image string) {
	imageName, imageTag := parseImageName(image)
	oci.UnpackImage(imagesDir, containersDir, name, imageName, imageTag)
}

func parseImageName(name string) (string, string) {
	split := strings.Split(name, ":")
	if len(split) != 2 {
		log.Fatal("Invalid image name")
	}
	return split[0], split[1]
}
