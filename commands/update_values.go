package commands

import (
	"fmt"
	"os"
	"tony/helm"

	"github.com/urfave/cli/v2"
)

func UpdateValues() *cli.Command {
	return &cli.Command{
		Name:  "setHelmValues",
		Usage: "set the hostname and latest VDD versions in values.yaml",
		Action: func(cCtx *cli.Context) error {

			// create an initial marshalled file
			hn, _ := os.Hostname()
			appVersions := helm.Versions{Foo: "hi", Bar: "there"}
			type stuff struct {
				Name string
				Age  int
			}
			myMoreOther := stuff{"andrew", 12}
			helmValues := helm.HelmValues{
				Other:       "bob",
				MoreOther:   myMoreOther,
				AppVersions: appVersions,
				Hostname:    hn,
			}
			marshalledFile := helmValues.MarshalYAML()
			fmt.Println(string(marshalledFile))

			// unmarshal the file and apply the new versions
			helmValues.UnmarshYAML(marshalledFile)
			helmValues.AppVersions.Foo = "by"
			helmValues.AppVersions.Bar = "now"

			// remarshal with updated versions
			updatedMarshalValues := helmValues.MarshalYAML()
			fmt.Println()
			fmt.Println(string(updatedMarshalValues))

			return nil
		},
	}
}
