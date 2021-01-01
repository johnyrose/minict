package run

import (
	"fmt"
	"log"
	"strings"
)

func RunContainer(imagesDir string, containersDir string, name string, image string) {
	// imageName, imageTag := parseImageName(image)
	// oci.UnpackImage(imagesDir, containersDir, name, imageName, imageTag)
	imageConfig := GetImageConfig(containersDir + "/" + name)
	fmt.Print(imageConfig)
}

func parseImageName(name string) (string, string) {
	split := strings.Split(name, ":")
	if len(split) != 2 {
		log.Fatal("Invalid image name")
	}
	return split[0], split[1]
}
