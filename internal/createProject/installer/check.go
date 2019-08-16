package installer

import (
	"errors"
	"os"
	"path/filepath"
)


type installer struct {
	ProjectName   string
	AppImportPath string
	F7kImportPath string
	CurrentDir    string
	Port          uint16
}


func (i *installer) checkData() *installer {
	p, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	i.CurrentDir = p
	return i.
		checkString("ProjectName", i.ProjectName).
		checkString("AppImportPath (like github.com/dev2choiz/f7k)", i.AppImportPath).
		checkInt("port number", int(i.Port))

}

func (i *installer) checkString(n, v string) *installer {
	if "" == v {
		panic(errors.New(n + " not provided"))
	}

	return i
}

func (i *installer) checkInt(n string, v int) *installer {
	if 0 == v {
		panic(errors.New(n + " not provided"))
	}

	return i
}

func (i *installer) getAbsProjectDir() string {
	return filepath.Join(i.CurrentDir, i.ProjectName)
}
