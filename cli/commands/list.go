package commands

import (
	"fmt"
	"io/ioutil"
	"log"
	"oci"
)

func ListContainers(containersDir string) []string {
	files, err := ioutil.ReadDir(containersDir)
	if err != nil {
		log.Fatal(err)
	}
	var containers []string
	for _, file := range files {
		if file.IsDir() {
			containers = append(containers, file.Name())
		}
	}
	return containers
}

func ListImages(imagesDir string) []string {
	var images []string
	manifests := oci.ListManifests(imagesDir)
	for _, manifest := range manifests {
		images = append(images, fmt.Sprintf("%s:%s", manifest.Name, manifest.Annotations.Tag))
	}
	return images
}
