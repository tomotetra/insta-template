package main

import (
	"log"
	"os"

	"github.com/tomotetra/instatemplate/cmd/instacli"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "instatemplate"
	app.Usage = "create an instagram template"
	app.Action = instacli.GenerateTemplate
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
