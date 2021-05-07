package run

import (
	"os"
	"strings"
	"syscall"
)

func performMount(mount MountsConfig) error {
	if mount.Type == "bind" {
		prepareBindMount(mount)
	}
	mountFlag := getMountFlag(mount)
	return syscall.Mount(mount.Source, "rootfs"+mount.Destination, mount.Type, mountFlag, strings.Join(mount.Options, ","))
}

func getMountFlag(mount MountsConfig) uintptr {
	if mount.Type == "bind" {
		return syscall.MS_BIND
	}
	if mount.Type == "tmpfs" {
		return syscall.MS_NOSUID
	}
	return 0
}

func prepareBindMount(mount MountsConfig) {
	sourceFile, _ := os.Stat(mount.Source)
	if sourceFile.IsDir() {
		os.MkdirAll("rootfs"+mount.Destination, os.ModePerm)
	} else {
		filenameSplit := strings.Split(mount.Destination, "/")
		filenameSplit = filenameSplit[:len(filenameSplit)-1]
		os.MkdirAll("rootfs/"+strings.Join(filenameSplit, "/"), os.ModePerm)
		os.Create("rootfs" + mount.Destination)
	}
}
