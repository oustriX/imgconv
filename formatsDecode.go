package main

import (
	"golang.org/x/image/bmp"
	"golang.org/x/image/tiff"
	"golang.org/x/image/webp"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
)

func DecodePNG(f *os.File) (image.Image, error) {
	return png.Decode(f)
}

func DecodeJPEG(f *os.File) (image.Image, error) {
	return jpeg.Decode(f)
}

func DecodeGIF(f *os.File) (image.Image, error) {
	return gif.Decode(f)
}

func DecodeTIFF(f *os.File) (image.Image, error) {
	return tiff.Decode(f)
}

func DecodeBMP(f *os.File) (image.Image, error) {
	return bmp.Decode(f)
}

func DecodeWEBP(f *os.File) (image.Image, error) {
	return webp.Decode(f)
}
