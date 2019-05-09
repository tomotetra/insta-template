package exif

import (
	"os"

	"github.com/rwcarlsen/goexif/exif"
	"github.com/rwcarlsen/goexif/tiff"
	"github.com/tomotetra/instagen/src/utils/format"
	"github.com/tomotetra/instagen/src/utils/logger"
)

// InstaExifs struct
type InstaExifs struct {
	Date        string
	CameraModel string
	LensModel   string
	Settings    InstaSettings
}

// InstaSettings struct
type InstaSettings struct {
	ISO          string
	FocalLength  string
	FNumber      string
	ShutterSpeed string
}

// ReadExif reads exif data of an image file
func ReadExif(fileName string) InstaExifs {
	targetImage, err := os.Open(fileName)
	if err != nil {
		logger.Fatal("failed to read file")
	}
	x, err := exif.Decode(targetImage)
	if err != nil {
		logger.Fatal("Error: failed to read exif file")
	}
	settings := InstaSettings{
		ISO:          getField(x, "ISOSpeedRatings"),
		FocalLength:  getField(x, "FocalLength"),
		FNumber:      getField(x, "FNumber"),
		ShutterSpeed: getField(x, "ExposureTime"),
	}
	return InstaExifs{
		Date:        format.DateTimeString(getField(x, "DateTimeOriginal")),
		CameraModel: getField(x, "Model"),
		LensModel:   getField(x, "LensModel"),
		Settings:    settings,
	}
}

func getField(x *exif.Exif, fieldName exif.FieldName) string {
	tag, err := x.Get(fieldName)
	if err != nil {
		logger.Fatal(err)
	}
	stringVal := format.RemoveQuotes(tag.String())
	if tag.Format() == tiff.RatVal {
		rat, _ := tag.Rat(0)
		stringVal = format.RatToString(rat)
	}
	return stringVal
}
