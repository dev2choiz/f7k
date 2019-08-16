package command

import (
	"github.com/dev2choiz/f7k"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"path/filepath"
)

type DataYaml struct {
	Commands 	map[string]struct {
		ImportPath    string 	`yaml:"importPath"`
		Package       string 	`yaml:"packageName"`
		Function      string 	`yaml:"function"`
	}						 	`yaml:"commands"`
}

var commandYaml = &DataYaml{}

func ParseYaml() error {
	f := filepath.Dir(f7k.AppLoader.ConfFile())
	f = filepath.Join(f, "commands.yaml")

	s, err := ioutil.ReadFile(f)
	if nil != err {
		return err
	}

	err = yaml.Unmarshal(s, commandYaml)
	if nil != err {
		return err
	}

	return nil
}
