package main

import (
	"errors"
	"fmt"
	"github.com/urfave/cli/v2"
	"image"
	"os"
	"strings"
)

type ConvertingFile struct {
	img image.Image
	f   *os.File
}

func ConvertImage(cmd *cli.Context) error {
	srcImagePath := cmd.Args().Get(0)
	newFormat := cmd.Args().Get(1)
	newFilePath := cmd.Args().Get(2)
	fmt.Println(srcImagePath, newFormat, newFilePath)

	if newFilePath == "" {
		newFilePath = "./newImage." + newFormat
	}

	srcImageFormat := getSrcImageFormat(srcImagePath)

	f, err := os.Open(srcImagePath)
	if err != nil {
		fmt.Printf("fail to open src file: %s\n", srcImagePath)
		return err
	}
	defer f.Close()

	img, err := decodeImage(f, srcImageFormat)
	if err != nil {
		fmt.Printf("fail to decode image\n")
		return err
	}

	// create file for image with new format
	newFile, err := os.Create(newFilePath)
	if err != nil {
		fmt.Printf("fail to create new file with path: %s\n", newFilePath)
		return err
	}
	defer newFile.Close()

	// encode old image to new file with new format
	err = encodeImage(newFile, img, newFormat)
	if err != nil {
		fmt.Printf("fail to encode new image\n")
		return err
	}
	return nil
}

func getSrcImageFormat(path string) string {
	parts := strings.Split(path, ".")
	format := parts[len(parts)-1]
	return format
}

func decodeImage(f *os.File, format string) (image.Image, error) {
	switch format {
	case "png":
		return decodePNG(f)

	case "jpeg", "jpg":
		return decodeJPEG(f)

	case "gif":
		return decodeGIF(f)

	case "tiff":
		return decodeTIFF(f)

	case "bmp":
		return decodeBMP(f)

	case "webp":
		return decodeWEBP(f)

	default:
		img, _, err := image.Decode(f)
		return img, err
	}
}

func encodeImage(f *os.File, img image.Image, newFormat string) error {
	cf := ConvertingFile{f: f, img: img}
	switch strings.ToLower(newFormat) {
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
