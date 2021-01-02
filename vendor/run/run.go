package run

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"syscall"
)

func RunContainer(containersDir string, name string) error {
	containerDir := fmt.Sprintf("%s/%s", containersDir, name)
	imageConfig := GetImageConfig(containerDir)
	os.Chdir(containerDir)
	cmd := buildCommand(imageConfig)
	applyNamespaces(cmd)
	applyChroot(imageConfig)
	err := cmd.Run()
	return err
}

func parseImageName(name string) (string, string) {
	split := strings.Split(name, ":")
	if len(split) != 2 {
		log.Fatal("Invalid image name")
	}
	return split[0], split[1]
}

func buildCommand(imageConfig ImageConfig) *exec.Cmd {
	cmd := exec.Command(imageConfig.ProcessConfig.Args[0], imageConfig.ProcessConfig.Args[1:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Env = imageConfig.ProcessConfig.Env
	return cmd
}

func applyNamespaces(cmd *exec.Cmd) {
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS,
	}
}

func applyChroot(imageConfig ImageConfig) {
	syscall.Chroot("rootfs")
	for _, mount := range imageConfig.MountsConfig {
		var mountFlag uintptr
		if mount.Type == "bind" {
			mountFlag = syscall.MS_BIND
		} else {
			mountFlag = 0
		}
		err := syscall.Mount(mount.Source, mount.Destination, mount.Type, mountFlag, strings.Join(mount.Options, ","))
		if err != nil {
			log.Print(("Failed to mount " + mount.Source + " to " + mount.Destination + " due to " + err.Error()))
		}
		// TODO: Implement more mount options and clean the code to not use if-else for every mount type.
	}
	os.Chdir(imageConfig.ProcessConfig.Cwd)
}
