package main

import (
	"log"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "makerpool-cli"
	app.Usage = "manage maker's resources"
	app.Version = "0.0.1"

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
