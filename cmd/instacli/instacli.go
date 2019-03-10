package instacli

import (
	"fmt"

	"github.com/tomotetra/instatemplate/cmd/utils/logger"
	"github.com/urfave/cli"
)

// GenerateTemplate generates a intagram post template
func GenerateTemplate(c *cli.Context) error {
	fileName := c.Args().First()
	if fileName == "" {
		logger.Fatal("Usage: please specify a target file")
	}
	exifs := readExif(fileName)
	text := Template(exifs, "")
	// output result
	fmt.Println(text)
	return nil
}
