package commands

import (
	"fmt"
	"tony/helm"

	"github.com/urfave/cli/v2"
)

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
			AES := "AES256:"
			passwd := cCtx.String("password")
			ct := getCleartext(cCtx.String("filename"))
			fmt.Println(string(ct))
			out, err := helm.EncryptAES(passwd, string(ct))
			if err != nil {
				panic(err)
			}
			fmt.Printf("%s%s", AES, out)
			return nil
		},
	}
}
