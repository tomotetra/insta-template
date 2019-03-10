package instacli

import (
	"fmt"

	"github.com/tomotetra/instatemplate/cmd/utils/logger"
	"github.com/urfave/cli"
)

type flagParams struct {
	title string
	tags  string
}

// GeneratePost generates a intagram post from template
func GeneratePost(c *cli.Context) error {
	fileName := c.Args().First()
	params := setParams(c)
	if fileName == "" {
		logger.Fatal("Usage: please specify a target file")
	}
	exifs := readExif(fileName)
	text := Template(exifs, params)
	// output result
	fmt.Println(text)
	return nil
}

func setParams(c *cli.Context) *flagParams {
	return &flagParams{
		title: c.GlobalString("title"),
		tags:  c.GlobalString("tags"),
	}
}
