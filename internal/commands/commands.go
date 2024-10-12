package internal

import (
	"fmt"
	"log"
	"strings"
	"tony/helm"

	"github.com/urfave/cli/v2"
)

const AES = "AES256:"

func EncryptAES() *cli.Command {

	return &cli.Command{
		Name:    "encryptAES",
		Aliases: []string{"e"},
		Usage:   "encrypt file content",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "password",
				Aliases:  []string{"p"},
				Usage:    "encryption password",
				Required: true,
			},
		},
		Action: func(cCtx *cli.Context) error {
			p := cCtx.String("password")
			s := cCtx.Args().First()
			out, err := helm.EncryptAES(p, s)
			if err != nil {
				panic(err)
			}
			fmt.Printf("%s%s", AES, out)
			return nil
		},
	}
}

func DecryptAES() *cli.Command {
	return &cli.Command{
		Name:    "decryptAES",
		Aliases: []string{"d"},
		Usage:   "decrypt value",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "password",
				Aliases:  []string{"p"},
				Usage:    "decryption password",
				Required: true,
			},
		},
		Action: func(cCtx *cli.Context) error {
			p := cCtx.String("password")
			s := cCtx.Args().First()
			if strings.HasPrefix(cCtx.Args().First(), AES) {
				s = strings.TrimPrefix(s, AES)
				o, err := helm.DecryptAES(p, s)
				if err != nil {
					log.Fatalln(err)
				}
				fmt.Println(o)
			} else {
				fmt.Println("The string does not start with the prefix: ", AES)
			}
			return nil
		},
	}
}
