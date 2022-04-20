package main

import (
	"fmt"
	"github.com/kolesa-team/go-webp/encoder"
	"github.com/kolesa-team/go-webp/webp"
	"golang.org/x/image/bmp"
	"golang.org/x/image/tiff"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
)

// TODO: Переделать все на cli-библиотеку

func imgconv() {

	// checking for all arguments
	if len(os.Args) < 4 {
		fmt.Println("You missed some arguments.")
		fmt.Printf("Must be: 3. You sent: %d\n", len(os.Args)-1)
		return
	}

	// open image file
	f, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Printf("cannot open file %s\n", os.Args[2])
		fmt.Printf("error message: %s\n", err.Error())
		return
	}
	defer f.Close()

	// decode original image
	img, _, err := image.Decode(f)
	if err != nil {
		fmt.Printf("unsupported format: %s\n", os.Args[1])
		fmt.Printf("error message: %s\n", err.Error())
		return
	}

	// create file for new image
	newFile, err := os.Create(os.Args[3])
	if err != nil {
		fmt.Printf("cannot create file %s\n", os.Args[3])
		fmt.Printf("error message: %s\n", err.Error())
		return
	}
	defer newFile.Close()

	switch os.Args[2] {
	case "-png":
		err := convertToPNG(img, newFile)
		if err != nil {
			convertError(err)
			return
		}

	case "-gif":
		err := convertToGIF(img, newFile)
		if err != nil {
			convertError(err)
			return
		}

	case "-jpeg":
		err := convertToJPEG(img, newFile)
		if err != nil {
			convertError(err)
			return
		}

	case "-tiff":
		err := convertToTIFF(img, newFile)
		if err != nil {
			convertError(err)
			return
		}

	case "-bmp":
		err := convertToBMP(img, newFile)
		if err != nil {
			convertError(err)
			return
		}

	case "webp":
		err := convertToWEBP(img, newFile)
		if err != nil {
			convertError(err)
			return
		}

	default:
		fmt.Println("Imgconv can't convert image to this format(")
		return
	}

	fmt.Println("Done!")

}

func convertToPNG(img image.Image, f *os.File) error {
	err := png.Encode(f, img)
	return err
}

func convertToGIF(img image.Image, f *os.File) error {
	err := gif.Encode(f, img, &gif.Options{})
	return err
}

func convertToJPEG(img image.Image, f *os.File) error {
	options := jpeg.Options{}
	err := jpeg.Encode(f, img, &options)
	return err
}

func convertToTIFF(img image.Image, f *os.File) error {
	options := tiff.Options{}
	err := tiff.Encode(f, img, &options)
	return err
}

func convertToBMP(img image.Image, f *os.File) error {
	err := bmp.Encode(f, img)
	return err
}

func convertToWEBP(img image.Image, f *os.File) error {
	options := encoder.Options{}
	err := webp.Encode(f, img, &options)
	return err
}

func convertError(err error) {
	fmt.Println("fail to convert image")
	fmt.Printf("error message: %d\n", err.Error())
}
