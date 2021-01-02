package commands

import (
	"fmt"
	"oci"

	"github.com/apex/log"
)

func Pull(imagesDir string, image string) {
	_, err := oci.PullImage(imagesDir, image)
	if err != nil {
		log.Fatal(fmt.Sprintf("Failed to pull image. Error received: %s", err.Error()))
	}
}
