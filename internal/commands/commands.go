package internal

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
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
