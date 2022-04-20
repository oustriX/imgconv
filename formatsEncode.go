package main

import (
	"golang.org/x/image/bmp"
	"golang.org/x/image/tiff"
	"image/gif"
	"image/jpeg"
	"image/png"
)

func encodePNG(cf ConvertingFile) error {
	return png.Encode(cf.f, cf.img)
}

func encodeJPEG(cf ConvertingFile) error {
	return jpeg.Encode(cf.f, cf.img, &jpeg.Options{})
}

func encodeGIF(cf ConvertingFile) error {
	return gif.Encode(cf.f, cf.img, &gif.Options{})
}

func encodeTIFF(cf ConvertingFile) error {
	return tiff.Encode(cf.f, cf.img, &tiff.Options{})
}

func encodeBMP(cf ConvertingFile) error {
	return bmp.Encode(cf.f, cf.img)
}

func encodeWEBP(cf ConvertingFile) error {
	return bmp.Encode(cf.f, cf.img)
}
