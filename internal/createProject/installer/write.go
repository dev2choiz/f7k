package installer

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func (i installer) Execute() {
	i.
		checkData().
		writeNewProject().
		writeGoModFile().
		populatePlaceholders().
		replacePlaceholders()
}

func New() *installer {
	i := &installer{}
	i.F7kImportPath = "github.com/dev2choiz/f7k"
	return i
}

func (i *installer) writeNewProject() *installer {
	for _, f := range AssetNames() {
		d, err := Asset(f)
		if nil != err {
			panic(err)
		}

		a := strings.Split(f, "/sample/")[1]
		//t := filepath.Join(append([]string{i.getAbsProjectDir()}, a)...)
		t := filepath.Join(i.getAbsProjectDir(), a)
		err = os.MkdirAll(filepath.Dir(t), os.FileMode(0755))
		if nil != err {
			panic(err)
		}
		err = ioutil.WriteFile(t, d, os.FileMode(0755))
		if err != nil {
			panic(err)
		}
	}
	return i
}

func (i *installer) writeGoModFile() *installer {
	tmpl := fmt.Sprintf(`module %s

go 1.12

require (
)
`, i.AppImportPath)
	f := filepath.Join(i.getAbsProjectDir(), "go.mod")
	err := ioutil.WriteFile(f, []byte(tmpl), os.FileMode(0755))
	if err != nil {
		panic(err)
	}
	return i
}
