package run

import (
	"strings"
	"syscall"
)

func performMount(mount MountsConfig) error {
	var mountFlag uintptr
	if mount.Type == "bind" {
		prepareBindMount(mount)
		mountFlag = syscall.MS_BIND
	} else {
		mountFlag = 0
	}
	return syscall.Mount(mount.Source, "rootfs"+mount.Destination, mount.Type, mountFlag, strings.Join(mount.Options, ","))
}
