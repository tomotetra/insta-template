package template

import (
	"fmt"
	"os"
	"strings"

	"github.com/tomotetra/instagen/src/exif"
	"github.com/tomotetra/instagen/src/utils/io"
	"github.com/tomotetra/instagen/src/utils/logger"
)

type Params struct {
	Title    string
	HashTags string
}

// BuildPost builds the post template from exifs
func BuildPost(x exif.InstaExifs, params *Params) string {
	var lines []string
	if len(params.Title) > 0 {
		lines = append(lines, params.Title)
	}
	lines = append(lines, x.Date)
	lines = append(lines, "---")
	lines = append(lines, x.CameraModel)
	lines = append(lines, x.LensModel)
	lines = append(lines, fmt.Sprintf("ISO%s %smm f/%s %ss",
		x.Settings.ISO, x.Settings.FocalLength, x.Settings.FNumber, x.Settings.ShutterSpeed))
	lines = append(lines, ".")
	lines = append(lines, buildHashTagString(params.HashTags))
	return strings.Join(lines, "\n")
}

func buildHashTagString(t string) string {
	// maximum count of hashtags permitted
	MaxHashtags := 30

	var tags []string
	if len(t) > 0 {
		tags = append(tags, strings.Split(t, " ")...)
	}
	commonTags, _ := io.ReadLines(fmt.Sprintf("%s/.instagenTags", os.Getenv("HOME")))
	// commonTags, _ := io.ReadLines("./instagenTags.txt")
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
