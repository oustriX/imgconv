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
	if isTooShort(len(os.Args)) {
		return
	}

	// open image file
	f, isError := openFile(os.Args[1])
	if isError {
		return
	}

	img, isError := imageDecode(f)
	if isError {
		return
	}

	newFile, isError := createNewFile(os.Args[3])
	if isError {
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

func isTooShort(length int) (isError bool) {
	if length < 4 {
		fmt.Println("You missed some arguments.")
		fmt.Printf("Must be: 3. You sent: %d\n", length-1)
		return true
	}
	return false
}

func openFile(path string) (f *os.File, isError bool) {
	f, err := os.Open(path)
	if err != nil {
		fmt.Printf("cannot open file %s\n", path)
		fmt.Printf("error message: %s\n", err.Error())
		return f, true
	}
	return f, false
}

func imageDecode(f *os.File) (img image.Image, isError bool) {
	img, _, err := image.Decode(f)
	if err != nil {
		fmt.Printf("cannot decode file %s\n", os.Args[1])
		fmt.Printf("error message: %s\n", err.Error())
		return img, true
	}
	return img, false
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

// create file for new image
func createNewFile(destination string) (newImageFile *os.File, isError bool) {
	newFilePath := destination
	newImageFile, err := os.Create(newFilePath)
	if err != nil {
		fmt.Printf("cannot create file %s\n", newFilePath)
		fmt.Printf("error message: %s\n", err.Error())
		return newImageFile, true
	}
	return newImageFile, false
}
