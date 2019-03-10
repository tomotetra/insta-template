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
	myExifs := readExif(fileName)
	logger.Success(fmt.Sprintf("%+v\n", myExifs))
	return nil
}
