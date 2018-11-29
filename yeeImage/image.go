/**
 * Created by angelina on 2017/5/15.
 */

package yeeImage

import (
	"code.google.com/p/graphics-go/graphics"
	"errors"
	"github.com/nfnt/resize"
	"github.com/oliamb/cutter"
	"github.com/yeeyuntech/yeego/yeeFile"
	"image"
	"image/draw"
	"image/gif"
	"image/jpeg"
	"image/png"
	"math"
	"os"
	"path"
	"strings"
)

// LoadImage
// 从文件中解码出image
func LoadImage(inputImagePath string) (img image.Image, err error) {
	file, err := os.Open(inputImagePath)
	if err != nil {
		return
	}
	defer file.Close()
	img, _, err = image.Decode(file)
	return
}

// SaveImage
// 将image保存到文件中，根据后缀判断保存形式，如果后缀不匹配，则png
func SaveImage(outputImagePath string, img image.Image) (err error) {
	imgFile, err := os.Create(outputImagePath)
	if err != nil {
		return err
	}
	defer imgFile.Close()
	switch strings.ToLower(path.Ext(outputImagePath)) {
	case ".png":
		err = png.Encode(imgFile, img)
	case ".jpg", ".jpeg":
		err = jpeg.Encode(imgFile, img, nil)
	case ".gif":
		err = gif.Encode(imgFile, img, nil)
	default:
		err = png.Encode(imgFile, img)
	}
	return
}

// GetImageInfo
// 获取图片的宽高信息
func GetImageInfo(inputFilePath string) (int, int, error) {
	imageFile, err := os.Open(inputFilePath)
	if err != nil {
		return 0, 0, err
	}
	defer imageFile.Close()
	img, _, err := image.Decode(imageFile)
	if err != nil {
		return 0, 0, err
	}
	imgWidth := img.Bounds().Max.X
	imgHeight := img.Bounds().Max.Y
	return imgWidth, imgHeight, nil
}

// ResizeImage
// 图片缩放
func ResizeImage(inputImagePath, outputImagePath string, width, height uint) error {
	inputImage, err := LoadImage(inputImagePath)
	if err != nil {
		return err
	}
	outputImage := resize.Resize(width, height, inputImage, resize.Lanczos3)
	err = SaveImage(outputImagePath, outputImage)
	if err != nil {
		return err
	}
	return nil
}

// ThumbImage
// 生成缩略图,根据传入的最大宽高生成缩略图,不会改变图片的比例
func ThumbImage(inputImagePath, outputImagePath string, maxWidth, maxHeight uint) error {
	inputFile, err := os.Open(inputImagePath)
	if err != nil {
		return err
	}
	defer inputFile.Close()
	switch strings.ToLower(path.Ext(inputImagePath)) {
	case ".jpg", ".jpeg":
		inputImage, err := jpeg.Decode(inputFile)
		if err != nil {
			return err
		}
		newImage := resize.Thumbnail(maxWidth, maxHeight, inputImage, resize.Lanczos3)
		outputFile, err := os.Create(outputImagePath)
		if err != nil {
			return err
		}
		defer outputFile.Close()
		err = jpeg.Encode(outputFile, newImage, &jpeg.Options{Quality: 80})
		if err != nil {
			return err
		}
	case ".png":
		inputImage, err := png.Decode(inputFile)
		if err != nil {
			return err
		}
		newImage := resize.Thumbnail(maxWidth, maxHeight, inputImage, resize.Lanczos3)
		outputFile, err := os.Create(outputImagePath)
		if err != nil {
			return err
		}
		defer outputFile.Close()
		err = png.Encode(outputFile, newImage)
		if err != nil {
			return err
		}
	case ".gif":
		inputImage, err := gif.Decode(inputFile)
		if err != nil {
			return err
		}
		newImage := resize.Thumbnail(maxWidth, maxHeight, inputImage, resize.Lanczos3)
		outputFile, err := os.Create(outputImagePath)
		if err != nil {
			return err
		}
		defer outputFile.Close()
		err = gif.Encode(outputFile, newImage, nil)
		if err != nil {
			return err
		}
	default:
		return errors.New("图片后缀错误，请传入jpg,jpeg,png,gif后缀的图片")
	}
	return nil
}

// RotateImage
// 图片旋转，顺时针
func RotateImage(inputImagePath, outputImagePath string, angle float64) error {
	imageExt := path.Ext(inputImagePath)
	if !(imageExt == ".png" || imageExt == ".jpg" || imageExt == ".jpeg" || imageExt == ".gif") {
		return errors.New("图片后缀错误，请传入jpg,jpeg,png,gif后缀的图片")
	}
	srcImage, err := LoadImage(inputImagePath)
	if err != nil {
		return err
	}
	w, h, err := GetImageInfo(inputImagePath)
	if err != nil {
		return err
	}
	newWidth := float64(w)*math.Cos(angle) + float64(h)*math.Sin(angle)
	newHeight := float64(w)*math.Sin(angle) + float64(h)*math.Cos(angle)
	dstImage := image.NewRGBA(image.Rect(0, 0, int(newWidth), int(newHeight)))
	err = graphics.Rotate(dstImage, srcImage, &graphics.RotateOptions{Angle: angle})
	if err != nil {
		return err
	}
	err = SaveImage(outputImagePath, dstImage)
	if err != nil {
		return err
	}
	return nil
}

// ImageProcessor
// 图片剪裁参数
type ImageProcessor struct {
	//距离左边距离
	LeftPoint int
	//距离上边距离
	TopPoint int
	//裁剪后的宽
	Width int
	//剪裁后的高
	Height int
}

// ProcessImage
// 图片裁剪
func ProcessImage(inputImagePath, outputImagePath string, processConfig ImageProcessor) error {
	if !yeeFile.FileExists(inputImagePath) {
		return errors.New("图片文件不存在,请检查地址")
	}
	if processConfig.Width == 0 && processConfig.Height == 0 {
		return errors.New("剪裁参数错误")
	}
	imageExt := path.Ext(inputImagePath)
	if !(imageExt == ".png" || imageExt == ".jpg" || imageExt == ".jpeg" || imageExt == ".gif") {
		return errors.New("图片后缀错误，请传入jpg,jpeg,png,gif后缀的图片")
	}
	var cutImageWidth = processConfig.Width
	var cutImageHeight = processConfig.Height
	var cutLeftPoint = processConfig.LeftPoint
	var cutTopPoint = processConfig.TopPoint
	img, err := LoadImage(inputImagePath)
	if err != nil {
		return err
	}
	w, h, err := GetImageInfo(inputImagePath)
	if err != nil {
		return err
	}
	if cutImageWidth > w && cutImageHeight > h {
		return nil
	}
	var outImagePointer image.Image
	if cutImageHeight != 0 && cutImageWidth != 0 {
		outImagePointer, _ = cutter.Crop(img, cutter.Config{
			Width:  cutImageWidth,
			Height: cutImageHeight,
			Anchor: image.Point{X: cutLeftPoint, Y: cutTopPoint},
			Mode:   cutter.TopLeft, // optional, default value
		})
	}
	// 判断如果裁剪点坐标存在小于0的情况，需要手动处理
	var jpg draw.Image
	jpg = image.NewRGBA(image.Rect(0, 0, cutImageWidth, cutImageHeight))
	if cutLeftPoint < 0 && cutTopPoint < 0 {
		draw.Draw(jpg, jpg.Bounds(), outImagePointer,
			outImagePointer.Bounds().Min.Sub(image.Pt(-cutLeftPoint, -cutTopPoint)), draw.Over)
	} else if cutLeftPoint < 0 {
		draw.Draw(jpg, jpg.Bounds(), outImagePointer,
			outImagePointer.Bounds().Min.Sub(image.Pt(-cutLeftPoint, 0)), draw.Over)
	} else if cutTopPoint < 0 {
		draw.Draw(jpg, jpg.Bounds(), outImagePointer,
			outImagePointer.Bounds().Min.Sub(image.Pt(0, -cutTopPoint)), draw.Over)
	}
	// 输出图片
	outputImage, err := os.Create(outputImagePath)
	if err != nil {
		return err
	}
	defer outputImage.Close()
	if cutLeftPoint >= 0 && cutTopPoint >= 0 {
		switch imageExt {
		case ".jpg", ".jpeg":
			err = jpeg.Encode(outputImage, outImagePointer, nil)
			if err != nil {
				return err
			}
		case ".png":
			err = png.Encode(outputImage, outImagePointer)
			if err != nil {
				return err
			}
		case ".gif":
			err = gif.Encode(outputImage, outImagePointer, nil)
			if err != nil {
				return err
			}
		}
	} else {
		switch imageExt {
		case ".jpg", ".jpeg":
			err = jpeg.Encode(outputImage, jpg, nil)
			if err != nil {
				return err
			}
		case ".png":
			err = png.Encode(outputImage, jpg)
			if err != nil {
				return err
			}
		case ".gif":
			err = gif.Encode(outputImage, jpg, nil)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
