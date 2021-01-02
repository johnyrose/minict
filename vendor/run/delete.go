package run

import (
	"fmt"
	"os"
	"syscall"

	"github.com/apex/log"
)

func DeleteContainer(containersDir string, name string) {
	containerDir := fmt.Sprintf("%s/%s", containersDir, name)
	imageConfig := GetImageConfig(containerDir)
	unmountAll(containerDir, imageConfig)
	err := os.RemoveAll(containerDir)
	if err != nil {
		log.Error(fmt.Sprintf("Failed to delete container files. Error received: %s", err.Error()))
	}
}

func unmountAll(containerDir string, config ImageConfig) {
	for _, mount := range config.MountsConfig {
		err := syscall.Unmount(fmt.Sprintf("%s/rootfs/%s", containerDir, mount.Destination), 0)
		if err != nil {
			log.Error(fmt.Sprintf("Failed to umount %s, error received: %s", mount.Destination, err.Error()))
		}
	}
}
