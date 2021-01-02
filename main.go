package main

import (
	"fmt"
	"log"
	"os"

	_ "github.com/containers/image/docker"
	_ "github.com/containers/image/oci/layout"
	"github.com/urfave/cli"
)

func main() {
	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:  "pull",
				Usage: "Pull an image from Dockerhub or a different container registry.",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "image",
						Usage:    "The full image name",
						Required: true,
					},
				},
				Action: func(c *cli.Context) error {
					fmt.Println("test pull " + c.String("image"))
					return nil
				},
			},
			{
				Name:  "run",
				Usage: "Run a container.",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "image",
						Usage:    "The full image name",
						Required: true,
					},
				},
				Action: func(c *cli.Context) error {
					fmt.Println("test run " + c.Args().First())
					return nil
				},
			},
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
