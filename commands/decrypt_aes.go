package commands

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"
	"tony/helm"

	"github.com/urfave/cli/v2"
)

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
			AES := "AES256:"
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

// getCleartext returns content of file f
func getCleartext(f string) []byte {
	// determine if path absolute
	if filepath.IsAbs(f) {
		ct, err := os.ReadFile(f)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		return ct
	}

	// if path is not absolute, determine filepath relative to working directory.
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	ct, err := os.ReadFile(path.Join(dir, f))
	if err != nil {
		fmt.Println("error: ", err)
		os.Exit(1)
	}
	return ct
}
