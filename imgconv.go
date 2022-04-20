package main

import (
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"sort"
)

//var supportedFormats = []string{"png", "jpeg", "jpg"}

func main() {
	app := &cli.App{
		Name:                  "imgconv",
		Usage:                 "simple image format converter",
		Action:                ConvertImage,
		CustomAppHelpTemplate: appHelpTemplate,
		EnableBashCompletion:  true,
		Version:               "1.0.0",
		Commands: []*cli.Command{
			{
				Name:    "src",
				Aliases: []string{"s"},
				Usage:   "[required] the `ORIGINAL IMAGE` to be converted",
			},
			{
				Name:    "newFormat",
				Aliases: []string{"f"},
				Usage:   "[required] `FORMAT` for the new file",
			},
			{
				Name:    "destination",
				Aliases: []string{"dst", "d"},
				Usage:   "`PATH` for save image with new format (default: ./newImage.[newFormat]",
			},
		},
	}

	sort.Sort(cli.CommandsByName(app.Commands))

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
