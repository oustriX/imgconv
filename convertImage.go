package main

import (
	"errors"
	"github.com/urfave/cli/v2"
	"image"
	"os"
	"strings"
)

type ConvertingFile struct {
	img image.Image
	f *os.File
}


func ConvertImage(cmd *cli.Context) error {
	args := cmd.Args()
	srcImagePath := args.Get(1)
	newFormat := args.Get(2)
	newFilePath := args.Get(3)

	if newFilePath == "" {
		newFilePath = "./newImage" + newFormat
	}

	srcImageFormat := getSrcImageFormat(srcImagePath)

	f, err := os.Open(srcImagePath)
	if err != nil {
		return err
	}

	img, err := decodeImage(f, srcImageFormat)
	if err != nil {
		return err
	}

	// create file for image with new format
	newFile, err := os.Create(newFilePath)
	if err != nil {
		return err
	}

	// encode old image to new file with new format
	err = encodeImage(newFile, img, newFormat)
	if err != nil {
		return err
	}
	return nil
}

func getSrcImageFormat(path string) string {
	parts := strings.Split(path, ".")
	format := parts[len(parts) - 1]
	return format
}

// TODO: change default case
func decodeImage(f *os.File, format string) (image.Image, error){
	switch format {
	case "png":
		return DecodePNG(f)

	case "jpeg", "jpg":
		return DecodeJPEG(f)

	case "gif":
		return DecodeGIF(f)

	case "tiff":
		return DecodeTIFF(f)

	case "bmp":
		return DecodeBMP(f)

	case "webp":
		return DecodeWEBP(f)

	default:
		return nil, errors.New("unsupported format")
	}
}

func encodeImage(f *os.File, img image.Image, newFormat string) error {
	cf := ConvertingFile{f: f, img: img}
	switch newFormat {
	case "png":
		return encodePNG(cf)

	case "jpeg", "jpg":
		return encodeJPEG(cf)

	case "gif":
		return encodeGIF(cf)

	case "tiff":
		return encodeTIFF(cf)

	case "bmp":
		return encodeBMP(cf)

	case "webp":
		return encodeWEBP(cf)

	default:
		return errors.New("unsupported conversion format")
	}
}
