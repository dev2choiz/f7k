package command

import (
	"bytes"
	"github.com/dev2choiz/f7k"
	"github.com/dev2choiz/f7k/cacheGen"
	"github.com/dev2choiz/f7k/model/events"
	"github.com/dev2choiz/f7k/pkg/prompt"
	"html/template"
	"os"
	"path"
	"path/filepath"
	"sort"
)

type cacheCom struct {
	*cacheGen.CacheListener
}

func OnCacheGenCommand(event *events.CacheGenEvent) {
	err := ParseYaml()
	if nil != err {
		prompt.New("info").Println("no command was found")
		return
	}
	cc := &cacheCom{cacheGen.NewCacheListener()}
	f := cc.writeCache()
	if "" != f {
		event.GeneratedFiles = append(event.GeneratedFiles, f)
		event.PreAppLoadFunctions = append(event.PreAppLoadFunctions, "command.PrepareCommands")
		event.ImportCachePackages = append(event.ImportCachePackages, "command")
	}
}

func (cc *cacheCom) writeCache() string {
	data := struct {
		*DataYaml
		Imports []string
	}{
		commandYaml,
		make([]string, 0),
	}

	for _, cmd := range commandYaml.Commands {
		if !cc.contains(data.Imports, cmd.ImportPath) {
			data.Imports = append(data.Imports, cmd.ImportPath)
		}
	}

	if 0 == len(data.Commands) {
		return ""
	}

	data.Imports = append(data.Imports, path.Join(f7k.AppPath.F7kGitImportPath(), "command"))
	data.Imports = append(data.Imports, path.Join(f7k.AppPath.F7kGitImportPath(), "interfaces"))
	sort.Strings(data.Imports)

	tmpl, err := template.ParseFiles(filepath.Join(f7k.AppPath.F7kPath(), "command", "cacheCom.go.tmpl"))
	if nil != err {
		panic(err)
	}
	w := bytes.NewBuffer(nil)
	err = tmpl.Execute(w, data)
	if nil != err {
		panic(err)
	}
	wd, _ := os.Getwd()
	dir := path.Join(wd, "cache", "command")
	_ = os.MkdirAll(dir, os.FileMode(0755))
	filename := filepath.Join(dir, "commands.go")
	f, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	_ = f.Chmod(os.FileMode(0755))
	_, err = f.Write(w.Bytes())
	cc.handlePanic(err)

	return filename
}

func (cc *cacheCom) contains(a []string, x string) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}

func (cc *cacheCom) handlePanic(errs ...error) {
	for _, err := range errs {
		if nil != err {
			panic(err)
		}
	}
}
