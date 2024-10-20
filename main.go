package main

import (
	"log"
	"os"
	internal "tony/commands"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		EnableBashCompletion: true,

		Commands: []*cli.Command{
			internal.EncryptAES(),
			internal.DecryptAES(),
			internal.UpdateValues(),
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
