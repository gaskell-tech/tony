package hvault

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

func EncryptAESCommand() *cli.Command {
	return &cli.Command{
		Name:    "encryptAES",
		Aliases: []string{"e"},
		Usage:   "encrypt file content",
		Action: func(cCtx *cli.Context) error {
			fmt.Println("encrypting: ", cCtx.Args().First())
			return nil
		},
	}
}

func DecryptAESCommand() *cli.Command {
	return &cli.Command{
		Name:    "decryptAES",
		Aliases: []string{"d"},
		Usage:   "decrypt value",
		Action: func(cCtx *cli.Context) error {
			fmt.Println("decrypting: ", cCtx.Args().First())
			return nil
		},
	}
}
