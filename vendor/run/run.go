package run

import (
	"log"
	"os"
	"os/exec"
	"strings"
	"syscall"
)

func RunContainer(imagesDir string, containersDir string, name string) error {
	// imageName, imageTag := parseImageName(image)
	// oci.UnpackImage(imagesDir, containersDir, name, imageName, imageTag)
	containerDir := containersDir + "/" + name
	imageConfig := GetImageConfig(containerDir)
	os.Chdir(containerDir)
	cmd := buildCommand(imageConfig)
	applyNamespaces(cmd)
	applyChroot(imageConfig)
	err := cmd.Run()
	return err
	// TODO: Unmount the /proc folder that was created once the container exits, even if it exits with an error.
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
	// TODO: Add env variables and changing of working directory
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
	// for _, mount := range imageConfig.MountsConfig {
	// 	var mountFlag uintptr
	// 	if mount.Type == "bind" {
	// 		mountFlag = syscall.MS_BIND
	// 	} else {
	// 		mountFlag = 0
	// 	}
	// 	err := syscall.Mount(mount.Source, mount.Destination, mount.Type, mountFlag, strings.Join(mount.Options, ","))
	// 	if err != nil {
	// 		log.Print(("Failed to mount " + mount.Source + " to " + mount.Destination + " due to " + err.Error()))
	// 	}
	// 	// TODO: Implement more mount options and clean the code to not use if-else for every mount type.
	// }

	// TODO: Return the mounts when all were confirmed to work without possibility to damage the system
	os.Chdir(imageConfig.ProcessConfig.Cwd)
}
