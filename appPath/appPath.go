package appPath

import (
	"errors"
	"github.com/dev2choiz/f7k/interfaces"
	"os"
	"path"
	"path/filepath"
	"runtime"
)

const F7kImportPath = "github.com/dev2choiz/f7k"

type appPath struct {
	f7kGitImportPath 		string
	f7kPath          		string
	appRoot                	string
	gopathSrc              	string
	gopath                 	string
}

var instance  interfaces.AppPath

func Instance() interfaces.AppPath {
	return instance
}

func init() {
	i := &appPath{}

	i.gopath = os.Getenv("GOPATH")
	i.gopathSrc = path.Join(i.gopath, "src")
	i.f7kGitImportPath = F7kImportPath

	p, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	i.appRoot = p

	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		panic(errors.New("unable to retrieve f7k path"))
	}

	i.f7kPath = filepath.Dir(filepath.Dir(filename))

	instance = i
}


func (a *appPath) F7kGitImportPath() string {
	return a.f7kGitImportPath
}

func (a *appPath) SetF7kGitImportPath(f7kGitImportPath string) {
	a.f7kGitImportPath = f7kGitImportPath
}

func (a *appPath) F7kPath() string {
	return a.f7kPath
}

func (a *appPath) SetF7kPath(f7kPath string) {
	a.f7kPath = f7kPath
}

func (a *appPath) AppRoot() string {
	return a.appRoot
}

func (a *appPath) SetAppRoot(appRoot string) {
	a.appRoot = appRoot
}

func (a *appPath) GopathSrc() string {
	return a.gopathSrc
}

func (a *appPath) SetGopathSrc(gopathSrc string) {
	a.gopathSrc = gopathSrc
}

func (a *appPath) Gopath() string {
	return a.gopath
}

func (a *appPath) SetGopath(gopath string) {
	a.gopath = gopath
}
