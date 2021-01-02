package commands

import (
	"io/ioutil"
	"log"
)

func ListContainers(containersDir string) []string {
	files, err := ioutil.ReadDir(containersDir)
	if err != nil {
		log.Fatal(err)
	}
	var containers []string
	for _, file := range files {
		if file.IsDir() {
			containers = append(containers, file.Name())
		}
	}
	return containers
}
