package instacli

import (
	"fmt"
	"os"
	"strings"

	"github.com/tomotetra/instatemplate/cmd/utils/logger"

	"github.com/tomotetra/instatemplate/cmd/utils/io"
)

// Template builds the post template from exifs
func Template(x InstaExifs, params *flagParams) string {
	var lines []string
	if len(params.title) > 0 {
		lines = append(lines, params.title)
	}
	lines = append(lines, x.Date)
	lines = append(lines, "---")
	lines = append(lines, x.CameraModel)
	lines = append(lines, x.LensModel)
	lines = append(lines, fmt.Sprintf("ISO%s %smm f/%s %ss",
		x.Settings.ISO, x.Settings.FocalLength, x.Settings.FNumber, x.Settings.ShutterSpeed))
	lines = append(lines, ".")
	lines = append(lines, buildHashTags(params.tags))
	return strings.Join(lines, "\n")
}

func buildHashTags(t string) string {
	MaxHashtags := 30
	var tags []string
	if len(t) > 0 {
		tags = append(tags, strings.Split(t, " ")...)
	}
	commonTags, _ := io.ReadLines(fmt.Sprintf("%s/.instatemplate_tags", os.Getenv("HOME")))
	// commonTags, _ := io.ReadLines("./common_tags.txt")
	if len(commonTags) > 0 {
		tags = append(tags, commonTags...)
	}
	if len(tags) > MaxHashtags {
		logger.Warn(fmt.Sprintf("Too many hashtags. (provided: %d)", len(tags)))
		logger.Warn(fmt.Sprintf("Following hashtags were trimmed:\n%s\n", strings.Join(tags[MaxHashtags+1:], ", ")))
		tags = tags[:MaxHashtags]
	}

	var hashTags []string
	for _, tag := range tags {
		if len(tags) > 0 {
			hashTags = append(hashTags, fmt.Sprintf("#%s", tag))
		}
	}
	return strings.Join(hashTags, " ")
}
