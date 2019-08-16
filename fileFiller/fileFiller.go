package fileFiller

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

type fileFiller struct {
	DefaultPlaceholders map[string]string
	Placeholders 		map[string]string
}

func New() *fileFiller {
	ff := &fileFiller{make(map[string]string), make(map[string]string)}

	return ff
}

func (ff *fileFiller) setPlaceholders(p map[string]string) *fileFiller {
	ff.Placeholders = p

	return ff
}

func (ff *fileFiller) FillFiles(dir string) *fileFiller {
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}

		ff.FillFile(path)
		return nil
	})
	if err != nil {
		panic(err)
	}

	return ff
}

func (ff *fileFiller) FillFile(f string) *fileFiller {
	c, err := ioutil.ReadFile(f)
	if err != nil {
		panic(err)
	}

	for p, v := range ff.Placeholders {
		p = fmt.Sprintf("%s", p)
		c = bytes.ReplaceAll(c, []byte(p), []byte(v))
	}

	if err = ioutil.WriteFile(f, c, os.FileMode(0755)); err != nil {
		panic(err)
	}

	return ff
}
