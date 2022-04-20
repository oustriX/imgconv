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

func decodePNG(f *os.File) (image.Image, error) {
	return png.Decode(f)
}

func decodeJPEG(f *os.File) (image.Image, error) {
	return jpeg.Decode(f)
}

func decodeGIF(f *os.File) (image.Image, error) {
	return gif.Decode(f)
}

func decodeTIFF(f *os.File) (image.Image, error) {
	return tiff.Decode(f)
}

func decodeBMP(f *os.File) (image.Image, error) {
	return bmp.Decode(f)
}

func decodeWEBP(f *os.File) (image.Image, error) {
	return webp.Decode(f)
}
