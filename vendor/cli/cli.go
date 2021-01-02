package cli

import (
	"cli/commands"
	"fmt"

	"github.com/urfave/cli"
)

func GetCli() *cli.App {
	config := GetAppConfig()
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
				},
				Action: func(c *cli.Context) error {
					fmt.Println("test run " + c.Args().First())
					return nil
				},
			},
		},
	}
	return app
}
