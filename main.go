package main

import (
	"log"
	"oci"
)

func main() {
	err := oci.UnpackImage("/opt/fast_disk/Projects/mini-ct/images", "/opt/fast_disk/Projects/mini-ct/containers", "test", "ubuntu", "latest")
	if err != nil {
		log.Fatal(err)
	}
}
