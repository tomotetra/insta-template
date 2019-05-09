package flag

import (
	"github.com/urfave/cli"
)

// ArgumentFlags is a list of flags permitted in cli argument
var ArgumentFlags = map[string]cli.StringFlag{
	"Title": cli.StringFlag{
		Name:  "title, T",
		Usage: "Work title (default: blank)",
	},
	"Hashtags": cli.StringFlag{
		Name:  "hashtags, t",
		Usage: "Additional hashtags",
	},
}

// ArgumentKeys list of keys'
var ArgumentKeys = extractKeys(ArgumentFlags)

// ArgumentStringFlags list of cli.StringFlags'
var ArgumentStringFlags = extractStringFlags(ArgumentFlags)

func extractKeys(m map[string]cli.StringFlag) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

func extractStringFlags(m map[string]cli.StringFlag) []cli.Flag {
	flags := make([]cli.Flag, 0, len(m))
	for _, v := range m {
		flags = append(flags, v)
	}
	return flags
}
