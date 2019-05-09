package instagen

import (
	"fmt"

	"github.com/tomotetra/instagen/src/exif"
	"github.com/tomotetra/instagen/src/template"
	"github.com/tomotetra/instagen/src/utils/logger"
	"github.com/urfave/cli"
)

// GeneratePost generates a intagram post from template
func GeneratePost(c *cli.Context) error {
	fileName := c.Args().First()
	if len(fileName) == 0 {
		logger.Fatal("Usage: please specify a target file")
	}
	exifs := exif.ReadExif(fileName)
	flagParams := mapArgumentToParams(c)
	postText := template.BuildPost(exifs, &flagParams)
	// output result
	fmt.Println(postText)
	return nil
}

func mapArgumentToParams(c *cli.Context) template.Params {
	return template.Params{
		Title:    c.String("title"),
		HashTags: c.String("hashtags"),
	}
}
