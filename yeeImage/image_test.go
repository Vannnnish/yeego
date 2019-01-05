/**
 * Created by angelina on 2017/5/15.
 */

package yeeImage

import (
	"testing"
	"yeego"
)

const (
	TestImagePath    = "./testdata/test.jpg"
	ResizeImagePath  = "./testdata/resize.jpg"
	ThumbImagePath   = "./testdata/thumb.jpg"
	RotateImagePath  = "./testdata/rotate.jpg"
	ProcessImagePath = "./testdata/process.jpg"
)

func TestGetImageInfo(t *testing.T) {
	w, h, err := GetImageInfo(TestImagePath)
	yeego.Equal(w, 1920)
	yeego.Equal(h, 1200)
	yeego.Equal(err, nil)
}

func TestResizeImage(t *testing.T) {
	err := ResizeImage(TestImagePath, ResizeImagePath, 200, 200)
	yeego.Equal(err, nil)
	w, h, err := GetImageInfo(ResizeImagePath)
	yeego.Equal(w, 200)
	yeego.Equal(h, 200)
	yeego.Equal(err, nil)
}

func TestThumbImage(t *testing.T) {
	err := ThumbImage(TestImagePath, ThumbImagePath, 200, 200)
	yeego.Equal(err, nil)
	w, h, err := GetImageInfo(ThumbImagePath)
	yeego.Equal(w, 200)
	yeego.Equal(h, 200*1200/1920)
	yeego.Equal(err, nil)
}

func TestProcessImage(t *testing.T) {
	processConf := ImageProcessor{
		LeftPoint: 0,
		TopPoint:  0,
		Width:     200,
		Height:    200,
	}
	err := ProcessImage(TestImagePath, ProcessImagePath, processConf)
	yeego.Equal(err, nil)
}
