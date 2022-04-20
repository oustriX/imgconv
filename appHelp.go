package main

const appHelpTemplate = `NAME:
	{{.Name}} -  {{.Usage}}
USAGE:
	imgconv [src] [newFormat] [dst]
EXAMPLE:
	imgconv ./123.jpg png ./1/321.png

SUPPORTED FORMATS:
	PNG (png)
	JPEG (jpeg)
	JPG (jpg)
	GIF (gif)
	TIFF (tiff)
	BMP (bmp)
	WEBP (webp)
`
