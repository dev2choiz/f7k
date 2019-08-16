package overrider

import (
	"bytes"
	"errors"
	"github.com/dev2choiz/f7k"
	"github.com/dev2choiz/f7k/cacheGen"
	"github.com/dev2choiz/f7k/interfaces"
	"github.com/dev2choiz/f7k/model/events"
	"gopkg.in/yaml.v2"
	"html/template"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
)

type cacheOve struct {
	*cacheGen.CacheListener
}

var dataYaml = make(map[string]Override)

func OnCacheGenOverloader(e interfaces.Event) {
	event := e.(*events.CacheGenEvent)
	co := &cacheOve{cacheGen.NewCacheListener()}
	f := co.readYaml().writeCache()
	if "" != f {
		event.GeneratedFiles = append(event.GeneratedFiles, f)
		event.PreAppLoadFunctions = append(event.PreAppLoadFunctions, "overrider.PrepareOverriders")
		event.ImportCachePackages = append(event.ImportCachePackages, "overrider")
	}
}

func (co *cacheOve) readYaml() *cacheOve {
	p, err := os.Getwd()
	f := "conf/override.yaml"
	if _, err := os.Stat(f); os.IsNotExist(err) {
		f7k.Prompt.New("info").Printfln("%s")
		co.SetAbort(true)
		return co
	}

	s, err := ioutil.ReadFile(filepath.Join(p, f))
	if nil != err {
		panic(err)
	}

	err = yaml.Unmarshal(s, dataYaml)
	if nil != err {
		panic(err)
	}

	for _, o := range dataYaml {
		e := checkFields(&o)
		if nil != e {
			panic(e)
		}
	}
	return co
}

func checkFields(o interfaces.Overrider) error {
	if "" == o.GetTargetImportPath() {
		return errors.New("'importPath' not found in override.yaml")
	}
	if "" == o.GetTargetPackageName() {
		return errors.New("'packageName' not found in override.yaml")
	}
	if "" == o.GetFunctionName() {
		return errors.New("'functionName' not found in override.yaml")
	}

	return nil
}

func (co *cacheOve) writeCache() string {
	if co.Abort() {
		return ""
	}

	data := struct {
		Imports		[]string
		Overrides	map[string]Override
		Targets	    map[string]string
	}{}

	data.Overrides = dataYaml
	data.Targets = make(map[string]string)
	if nil == data.Overrides {
		return ""
	}

	if 0 < len(data.Overrides) {
		data.Imports = append(data.Imports, f7k.AppPath.F7kGitImportPath())
	}
	for n, o := range data.Overrides {
		switch n {
		case "eventDispatcher":
			data.Targets[n] = "f7k.Dispatcher"
			break
		case "kernel":
			data.Targets[n] = "f7k.Kernel"
			break
		// todo to complete
		}
		data.Imports = append(data.Imports, o.GetTargetImportPath())
	}

	wd, _ := os.Getwd()
	tmpl, err := template.ParseFiles(filepath.Join(f7k.AppPath.F7kPath(), "overrider", "cacheOve.go.tmpl"))
	if nil != err {
		panic(err)
	}
	w := bytes.NewBuffer(nil)
	err = tmpl.Execute(w, data)

	dir := path.Join(wd, "cache", "overrider")
	_ = os.MkdirAll(dir, os.FileMode(0755))
	filename := filepath.Join(dir, "overrider.go")
	f, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	_ = f.Chmod(os.FileMode(0755))
	_, _ = f.Write(w.Bytes())

	return filename
}
