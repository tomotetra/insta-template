package instacli

import (
	"fmt"
	"log"

	"github.com/urfave/cli"
)

// GenerateTemplate generates a intagram post template
func GenerateTemplate(c *cli.Context) error {
	fileName := c.Args().First()
	if fileName == "" {
		log.Fatal("Usage: please specify a target file")
	}
	myExifs := readExif(fileName)
	fmt.Printf("%+v\n", myExifs)
	return nil
}
