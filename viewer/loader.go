package viewer

import (
	"github.com/dev2choiz/f7k/interfaces"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path/filepath"
)

func Load(f string) interfaces.ConfViewer {
	c := &ConfView{}
	load(f, c)

	return c
}

func load(f string,c interfaces.ConfViewer) {
	p, err := os.Getwd()
	s, err := ioutil.ReadFile(filepath.Join(p, f))
	if nil != err {
		panic(err)
	}

	err = yaml.Unmarshal(s, c)
	if nil != err {
		panic(err)
	}
}
