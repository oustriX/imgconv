package main

import (
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
)

func main() {
	imgconv()
}

func imgconv() {

	// checking for all arguments
	if len(os.Args) < 4 {
		fmt.Println("You missed some arguments.")
		fmt.Printf("Must be: 3. You sent: %d\n", len(os.Args)-1)
		return
	}

	// open image file
	f, err := os.Open(os.Args[1])
	defer f.Close()
	if err != nil {
		fmt.Printf("cannot open file %s\n", os.Args[2])
		fmt.Printf("error message: %s\n", err.Error())
		return
	}

	// decode original image
	img, _, err := image.Decode(f)
	if err != nil {
		fmt.Printf("cannot decode file %s\n", os.Args[1])
		fmt.Printf("error message: %s\n", err.Error())
		return
	}

	// create file for new image
	newFile, err := os.Create(os.Args[3])
	defer newFile.Close()
	if err != nil {
		fmt.Printf("cannot create file %s\n", os.Args[3])
		fmt.Printf("error message: %s\n", err.Error())
		return
	}

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

func convertError(err error) {
	fmt.Println("fail to convert image")
	fmt.Printf("error message: %d\n", err.Error())
}
