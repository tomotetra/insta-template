package instacli

import (
	"fmt"
	"strings"

	"github.com/tomotetra/instatemplate/cmd/utils/io"
)

// Template builds the post template from exifs
func Template(x InstaExifs, hashtags string) string {
	var lines []string
	lines = append(lines, x.Date)
	lines = append(lines, "---")
	lines = append(lines, x.CameraModel)
	lines = append(lines, x.LensModel)
	lines = append(lines, fmt.Sprintf("ISO%s %smm f/%s %ss",
		x.Settings.ISO, x.Settings.FocalLength, x.Settings.FNumber, x.Settings.ShutterSpeed))
	lines = append(lines, ".")
	lines = append(lines, buildHashTags(hashtags))
	return strings.Join(lines, "\n")
}

func buildHashTags(t string) string {
	var tags []string
	if len(t) > 0 {
		tags = append(tags, strings.Split(t, " ")...)
	}
	commonTags, _ := io.ReadLines("common_tags.txt")
	if len(commonTags) > 0 {
		tags = append(tags, commonTags...)
	}

	var hashTags []string
	for _, tag := range tags {
		if len(tags) > 0 {
			hashTags = append(hashTags, fmt.Sprintf("#%s", tag))
		}
	}
	return strings.Join(hashTags, " ")
}
