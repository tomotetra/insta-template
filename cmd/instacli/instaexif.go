package instacli

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/rwcarlsen/goexif/tiff"

	"github.com/rwcarlsen/goexif/exif"
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

func readExif(fileName string) InstaExifs {
	targetImage, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("Error: failed to read file '%s'", fileName)
		os.Exit(1)
	}
	x, err := exif.Decode(targetImage)
	if err != nil {
		log.Fatal("Error: failed to read exif file")
	}
	settings := InstaSettings{
		ISO:          retrieveFormattedField(x, "ISOSpeedRatings"),
		FocalLength:  retrieveFormattedField(x, "FocalLength"),
		FNumber:      retrieveFormattedField(x, "FNumber"),
		ShutterSpeed: retrieveFormattedField(x, "ExposureTime"),
	}
	return InstaExifs{
		Date:        retrieveFormattedField(x, "DateTime"),
		CameraModel: retrieveFormattedField(x, "Model"),
		LensModel:   retrieveFormattedField(x, "LensModel"),
		Settings:    settings,
	}
}

func retrieveFormattedField(x *exif.Exif, fieldName exif.FieldName) string {
	val, err := x.Get(fieldName)
	if err != nil {
		log.Fatal(err)
	}
	stringVal := val.String()
	if val.Format() == tiff.RatVal {
		// trim denominator if 1
		stringVal = strings.Replace(stringVal, "/1", "", -1)
	}
	return stringVal
}
