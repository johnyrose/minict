package cli

import (
	"cli/commands"
	"encoding/json"
	"fmt"
	"os"

	"github.com/urfave/cli"
)

func initFolders(config AppConfig) {
	os.MkdirAll(config.ImagesDir, os.ModePerm)
	os.MkdirAll(config.ContainersDir, os.ModePerm)
}

func GetCli() *cli.App {
	config := GetAppConfig()
	initFolders(config)
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
					commands.Pull(config.ImagesDir, c.String("image"))
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
					&cli.StringFlag{
						Name:     "name",
						Usage:    "The name of the container",
						Required: true,
					},
				},
				Action: func(c *cli.Context) error {
					commands.Run(config.ImagesDir, config.ContainersDir, c.String("image"), c.String("name"))
					return nil
				},
			},
			{
				Name:  "rm",
				Usage: "Remove an existing container.",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "name",
						Usage:    "The name of the container",
						Required: true,
					},
				},
				Action: func(c *cli.Context) error {
					commands.RemoveContainer(config.ContainersDir, c.String("name"))
					return nil
				},
			},
			{
				Name:  "list-containers",
				Usage: "List containers that are running or can be ran have been ran.",
				Action: func(c *cli.Context) error {
					containers := commands.ListContainers(config.ContainersDir)
					b, err := json.MarshalIndent(containers, "", "	")
					if err == nil {
						fmt.Println(string(b))
					}
					return err
				},
			},
			{
				Name:  "list-images",
				Usage: "List images that were downloaded/added.",
				Action: func(c *cli.Context) error {
					images := commands.ListImages(config.ImagesDir)
					b, err := json.MarshalIndent(images, "", "	")
					if err == nil {
						fmt.Println(string(b))
					}
					return err
				},
			},
		},
	}
	return app
}
