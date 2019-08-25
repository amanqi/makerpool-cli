package main

import (
	"fmt"
	"log"
	"os"

	wip "./wip"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "makerpool-cli"
	app.Usage = "manage maker's resources"
	app.Version = "0.1"

	allFlags := []cli.Flag{
		cli.StringFlag{
			Name:  "platform, p",
			Usage: "set maker platform: wip",
		},
		cli.BoolFlag{
			Name:  "json",
			Usage: "output data in JSON",
		},
	}
	app.Flags = allFlags

	app.Commands = []cli.Command{
		{
			Name:    "profile",
			Aliases: []string{"p"},
			Usage:   "show authenticated user's profile",
			Action: func(c *cli.Context) error {
				outputFormat := ""
				if c.Bool("json") {
					outputFormat = "json"
				}
				if c.String("platform") == "" {
					return fmt.Errorf("platform not specified")
				}

				switch c.String("platform") {
				case "wip":
					err := wip.ShowProfile(outputFormat)
					if err != nil {
						return err
					}

				default:
					return fmt.Errorf(fmt.Sprintf("platform %s not supported",
						c.String("platform")))
				}

				return nil
			},
			Flags: allFlags,
		},
	}

	err := app.Run(os.Args)
	fatal(err)
}

func fatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
