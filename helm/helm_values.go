package helm

import "gopkg.in/yaml.v3"

type HelmValues struct {
	Other       any      `yaml:"othe"`
	MoreOther   any      `yaml:"moreOther"`
	Release     string   `yaml:"release"`
	Date        string   `yaml:"date"`
	AppVersions Versions `yaml:"appVersions"`
	Hostname    string   `yaml:"hostname"`
}

type Versions struct {
	Foo string `yaml:"foo"`
	Bar string `yaml:"bar"`
}

func (hv *HelmValues) UnmarshYAML(in []byte) {
	err := yaml.Unmarshal(in, &hv)
	if err != nil {
		panic(err)
	}
}

func (hv *HelmValues) MarshalYAML() []byte {
	out, err := yaml.Marshal(hv)
	if err != nil {
		panic(err)
	}
	return out
}
