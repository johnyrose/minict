package main

import (
	"log"
	"oci"

	_ "github.com/containers/image/docker"
	_ "github.com/containers/image/oci/layout"
)

func main() {
	_, err := oci.PullImage("/opt/fast_disk/Projects/mini-ct/images", "docker://alpine", "latest")
	if err != nil {
		log.Fatal(err)
	}
	// err := oci.UnpackImage("/opt/fast_disk/Projects/mini-ct/images", "/opt/fast_disk/Projects/mini-ct/containers", "test", "ubuntu", "latest")
	// if err != nil {
	// 	log.Fatal(err)
	// }
}

// (imagesDir string, imageName string, imageTag string)
