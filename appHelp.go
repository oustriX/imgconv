package main

const appHelpTemplate =
`NAME:
	{{.Name}} -  {{.Usage}}
USAGE:
	imgconv {{if .Commands}} command [command options]{{end}} {{if .ArgsUsage}}{{.ArgsUsage}}{{else}}[arguments...]{{end}}
SUPPORTED FORMATS:
	PNG (png)
	JPEG (jpeg)
	JPG (jpg)
`
