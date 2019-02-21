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
	text := Template(myExifs, "")
	logger.Success("done!")
	fmt.Println("=========== begin ===========")
	fmt.Println(text)
	fmt.Println("=========== end ===========")
	return nil
}
