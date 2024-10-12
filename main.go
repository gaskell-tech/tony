package main

import (
	"log"
	"os"
	"tony/hvault"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Commands: []*cli.Command{
			hvault.EncryptAESCommand(),
			hvault.DecryptAESCommand(),
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
