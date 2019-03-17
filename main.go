package main

import (
	"log"
	"os"

	"github.com/tomotetra/instagen/cmd/instagen"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "instagen"
	app.Usage = "create an instagram template"
	app.Action = instagen.GeneratePost
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "title, T",
			Usage: "Work title (default: blank)",
		},
		cli.StringFlag{
			Name:  "tags, t",
			Usage: "Additional hashtags",
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
