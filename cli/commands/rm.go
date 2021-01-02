package commands

import (
	"fmt"
	"log"
	"run"
)

func RemoveContainer(containersDir string, name string) {
	if !doesContainerExist(name, containersDir) {
		log.Fatal(fmt.Sprintf("Container with name %s does not exist.", name))
	}
	run.DeleteContainer(containersDir, name)
}
