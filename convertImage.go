package main

import (
	"errors"
	"github.com/urfave/cli/v2"
	"image"
	"os"
	"strings"
)

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

	return nil
}

func getSrcImageFormat(path string) string {
	parts := strings.Split(path, ".")
	format := parts[len(parts) - 1]
	return format
}

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
