package fileOperation

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/rwcarlsen/goexif/exif"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"os"
	"reflect"
	"regexp"
	"time"
)

type FileOperation struct {
	ctx context.Context
}

type LocLatLong struct {
	Latitude  float64
	Longitude float64
}
type ImageResolution struct {
	X int
	Y int
}

type CamMod struct {
	Name     string
	Software string
}

type ExifBase struct {
	Location    LocLatLong
	DateTaken   time.Time
	Size        ImageResolution
	CameraModel CamMod
}
type ExifInfo struct {
	ExifBase
	ExifMeta
}

type ExifResponse struct {
	ExifInfo ExifInfo
	FilePath string
}

type ExifMeta struct {
	ExifIFDPointer    string
	DateTimeDigitized string
	Orientation       string
	YResolution       string
	PixelYDimension   string
	Software          string
	DateTime          string
	ColorSpace        string
	XResolution       string
	PixelXDimension   string
}

func parseKeyValue(keyVaule []string) (string, string) {
	var value string
	if len(keyVaule[3]) == 0 {
		value = keyVaule[2]
	} else {
		value = keyVaule[3]
	}
	return keyVaule[1], value
}

func (e *ExifResponse) ParseExifString(exifStr string) {
	var re = regexp.MustCompile(`(?m)(\w+): ([0-9]+|"([a-zA-Z0-9-: .|\/]*)")`)
	keys := re.FindAllString(exifStr, -1)

	var exifRes ExifMeta

	for _, key := range keys {
		k, v := parseKeyValue(re.FindStringSubmatch(key))
		field := reflect.ValueOf(&exifRes).Elem().FieldByName(k)
		if field.IsValid() {
			field.Set(reflect.ValueOf(v))
		}
	}
	e.ExifInfo.ExifMeta = exifRes

}

func NewLocLatLong(lat, long float64, err error) LocLatLong {
	if err != nil {

		return LocLatLong{}
	}
	return LocLatLong{
		Latitude:  lat,
		Longitude: long,
	}
}

func NewFileOperation() *FileOperation {
	return &FileOperation{}
}

func (f *FileOperation) SetCtx(ctx context.Context) {
	f.ctx = ctx
}

func (f *FileOperation) GetFileInfo(filePath string) string {
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		return ""
	}
	b, err := json.Marshal(fileInfo)

	if err != nil {
		return ""
	}
	return string(b)
}

func (f *FileOperation) GetFileExif(fileName string) ExifResponse {
	fmt.Printf("got file name: '%s' \n", fileName)
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("could not open file. Error: %s", err)
		return ExifResponse{}
	}

	x, err := exif.Decode(file)
	if err != nil {
		fmt.Printf("Could not decode: %s", err.Error())
		return ExifResponse{}
	}

	var response ExifResponse
	response.FilePath = fileName

	response.ExifInfo.DateTaken, err = x.DateTime()
	response.ExifInfo.Location = NewLocLatLong(x.LatLong())
	response.ParseExifString(x.String())
	camModel, err := x.Get(exif.Model)
	if err != nil {
		fmt.Printf("could not get CamModel: %s \n", err.Error())
	} else {
		response.ExifInfo.CameraModel = CamMod{
			Name:     camModel.String(),
			Software: response.ExifInfo.Software,
		}
	}

	return response
}

func (f *FileOperation) OpenFileDialog() string {
	//file, err := runtime.OpenFileDialog(f.ctx, runtime.OpenDialogOptions{})
	files, err := runtime.OpenMultipleFilesDialog(f.ctx, runtime.OpenDialogOptions{})
	if err != nil {
		return ""
	}
	var result []ExifResponse
	for _, file := range files {
		result = append(result, f.GetFileExif(file))
	}
	response, err := json.Marshal(result)
	if err != nil {
		return ""
	}
	return string(response)
}
