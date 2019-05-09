package main

import (
	"os"

	"github.com/tomotetra/instagen/src/flag"
	"github.com/tomotetra/instagen/src/instagen"
	"github.com/tomotetra/instagen/src/utils/logger"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "instagen"
	app.Usage = "create an instagram template"
	app.Action = instagen.GeneratePost
	app.Flags = flag.ArgumentStringFlags
	err := app.Run(os.Args)
	if err != nil {
		logger.Fatal(err)
	}
}
