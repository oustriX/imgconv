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
		Name: "imgconv",
		Usage: "simple image format converter",
		Action: ConvertImage,
		CustomAppHelpTemplate: appHelpTemplate,
		EnableBashCompletion: true,
		Version: "1.0.0",
		Flags: []cli.Flag{
			&cli.PathFlag{
				Name: "src",
				Aliases: []string{"s"},
				Usage: "[required] the `ORIGINAL IMAGE` to be converted",
				Required: true,
			},
			&cli.StringFlag{
				Name: "format",
				Aliases: []string{"f"},
				Usage: "[required] `FORMAT` for the new file",
				Required: true,

			},
			&cli.PathFlag{
				Name: "destination",
				Aliases: []string{"dst", "d"},
				Usage: "[required] `PATH` for save image with new format",
				DefaultText: "./newImage.[format]",
			},
		},
	}

	sort.Sort(cli.CommandsByName(app.Commands))

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
