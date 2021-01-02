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
	applyUsers(imageConfig)
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
		Cloneflags: syscall.CLONE_NEWUTS |
			syscall.CLONE_NEWPID |
			syscall.CLONE_NEWNS |
			syscall.CLONE_NEWIPC,
	}
	// TODO: Add the syscall.CLONE_NEWUSER when user support is added.
	// TODO: Add the syscall syscall.CLONE_NEWNET when networking namespace support is added.
}

func applyChroot(imageConfig ImageConfig) {
	for _, mount := range imageConfig.MountsConfig {
		var mountFlag uintptr
		if mount.Type == "bind" {
			mountFlag = syscall.MS_BIND
			prepareBindMount(mount)
		} else {
			mountFlag = 0
		}
		err := syscall.Mount(mount.Source, "rootfs"+mount.Destination, mount.Type, mountFlag, strings.Join(mount.Options, ","))
		if err != nil {
			log.Print(fmt.Sprintf("Failed to mount %s to %s due to %s", mount.Source, mount.Destination, err.Error()))
		}
		// TODO: Implement more mount options and clean the code to not use if-else for every mount type.
	}
	syscall.Chroot("rootfs")
	os.Chdir(imageConfig.ProcessConfig.Cwd)
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

func applyUsers(imageConfig ImageConfig) {
	syscall.Setuid(imageConfig.ProcessConfig.User["uid"])
	syscall.Setgid(imageConfig.ProcessConfig.User["gid"])
}
