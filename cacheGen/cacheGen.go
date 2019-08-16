package cacheGen

import (
	"bytes"
	"fmt"
	"github.com/dev2choiz/f7k"
	"github.com/dev2choiz/f7k/interfaces"
	"github.com/dev2choiz/f7k/model/events"
	"github.com/dev2choiz/f7k/pkg/prompt"
	"html/template"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"reflect"
	"strings"
)

type cacheGen struct {}

var cacheGenInstance *cacheGen

var WaitingForListen []func(interfaces.Event)

func Instance() *cacheGen {
	if nil == cacheGenInstance {
		cacheGenInstance = &cacheGen{}
	}
	return cacheGenInstance
}

func (i *cacheGen) Run() *cacheGen {
	i.Print("info", "Cache generation...")

	wd, _ := os.Getwd()
	cachePath := path.Join(wd, "cache")
	_ = exec.Command("rm", "-rf", cachePath).Run()
	err := os.Mkdir(cachePath, os.FileMode(0755))
	if nil != err && ! strings.Contains(err.Error(), "file exists") {
		prompt.New("danger").Println( err.Error())
		panic(err)
	}

	e := &events.CacheGenEvent{}
	f7k.Dispatcher.Dispatch(events.OnCacheGenEvent, e)
	e.ImportCachePackages = f7k.Utils.Slice.DeduplicateStrings(e.ImportCachePackages)
	f7k.Utils.Slice.DeduplicateStrings(e.ImportCachePackages)


	for _, f := range e.GeneratedFiles {
		i.Print("info", f)
	}
	f := i.write(e)
	i.Print("info", f)

	return i
}

func (i *cacheGen) write(e *events.CacheGenEvent) string {
	tmpl, err := template.ParseFiles(filepath.Join(f7k.AppPath.F7kPath(), "cacheGen", "templates.go.tmpl"))
	if nil != err {
		panic(err)
	}


	wd, _ := os.Getwd()
	filename := filepath.Join(path.Join(wd, "cache"), "main.go")
	cachePath := path.Join(f7k.AppConfig.GetImportPath(), "cache") + "/"

	w := bytes.NewBuffer(nil)
	err = tmpl.Execute(w, struct {
		CachePath string
		PreAppLoadFunctions  []string
		PostAppLoadFunctions []string
		ImportCachePackages []string
	}{
		cachePath,
		e.PreAppLoadFunctions,
		e.PostAppLoadFunctions,

		e.ImportCachePackages,
	})
	if nil != err {
		panic(err)
	}

	f, err := os.Create(filename)
	if err != nil {
		panic(err)
	}

	defer f.Close()
	handlePanic(f.Chmod(os.FileMode(0755)))
	handlePanic(f.Write(w.Bytes()))

	c := strings.Split(filename, string(filepath.Separator))
	return strings.Join(c[len(c)-2:], string(filepath.Separator))
}

func handlePanic(errs ...interface{}) {
	for _, err := range errs {
		_, ok := err.(error)
		if ok && nil != err {
			panic(err)
		}
	}
}

func (i *cacheGen) STypedPrint(v interface{}) string {
	var q = ""
	t := reflect.ValueOf(v).Kind()
	switch t {
	case reflect.Slice :

		str := ""
		for _, value := range v.([]interface{}) {
			str += q + fmt.Sprintf("%s%v%s, ", q, i.STypedPrint(value), q)
		}

		return fmt.Sprintf("[]string{%s%s%s}", q, str, q)

	case reflect.String :
		q = "\""
		fallthrough

	default :
		return fmt.Sprintf(`%s%v%s`, q, v, q)
	}
}

func (i *cacheGen) EscapeDoubleQuotes(v string) string {
	return strings.ReplaceAll(v, "\"", "\\\"")
}

func (i *cacheGen) Print(p, s string) {
	if f7k.Verbose {
		prompt.New(p).Printfln(s)
	}
}
