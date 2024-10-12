package internal

import (
	"fmt"
	"os"
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
				Usage:    "encrypt using password `PASSWORD`",
				Required: true,
			},
			&cli.StringFlag{
				Name:     "filename",
				Aliases:  []string{"f"},
				Usage:    "encrypt content of the file `FILENAME`",
				Required: true,
			},
		},
		Action: func(cCtx *cli.Context) error {
			dir, err := os.Getwd()
			if err != nil {
				panic(err)
			}
			filepath := fmt.Sprint(dir, "\\", cCtx.String("filename"))
			fmt.Println(filepath)
			pwd := cCtx.String("password")
			ct, err := os.ReadFile(filepath)
			if err != nil {
				panic(err)
			}
			out, err := helm.EncryptAES(pwd, string(ct))
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
				o, _ := helm.DecryptAES(p, s)
				fmt.Println(o)
			} else {
				fmt.Println("The string does not start with the prefix: ", AES)
			}
			return nil
		},
	}
}
