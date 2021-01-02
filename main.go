package main

import (
	"log"
	"os"

	"github.com/Ripolak/minict/cli"

	_ "github.com/containers/image/docker"
	_ "github.com/containers/image/oci/layout"
)

func main() {
	app := cli.GetCli()
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
