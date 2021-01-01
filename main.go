package main

import (
	"oci"
)

func main() {
	oci.UnpackImage("/opt/fast_disk/Projects/mini-ct/images", "/opt/fast_disk/Projects/mini-ct/containers", "test", "ubuntu", "latest")
}
