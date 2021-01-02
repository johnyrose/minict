package main

import (
	"cli"
	"log"
	"os"
	"run"

	_ "github.com/containers/image/docker"
	_ "github.com/containers/image/oci/layout"
)

func main() {
	app := cli.GetCli()
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
	// _, err := oci.PullImage("/opt/fast_disk/Projects/mini-ct/images", "ubuntu:20.04")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// err = oci.UnpackImage("/opt/fast_disk/Projects/mini-ct/images", "/opt/fast_disk/Projects/mini-ct/containers", "ubuntu-test", "ubuntu", "20.04")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	run.RunContainer("/opt/fast_disk/Projects/mini-ct/images", "/opt/fast_disk/Projects/mini-ct/containers", "ubuntu-test", "ubuntu:20.04")
}
