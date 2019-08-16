package router

import (
	"bytes"
	"github.com/dev2choiz/f7k"
	"github.com/dev2choiz/f7k/cacheGen"
	"github.com/dev2choiz/f7k/controllers"
	"github.com/dev2choiz/f7k/interfaces"
	"github.com/dev2choiz/f7k/model/events"
	"gopkg.in/yaml.v2"
	"html/template"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"
)

type cacheYam struct {
	*cacheGen.CacheListener
}

func OnCacheGenYaml(e interfaces.Event) {
	event := e.(*events.CacheGenEvent)
	cr := controllers.RegistryInstance()
	cr.Imports = make([]string, 0) // re-init imports
	ci := cacheYam{cacheGen.NewCacheListener()}
	f := ci.populateRouterMetadataWithYaml().writeCache()
	if "" != f {
		event.GeneratedFiles = append(event.GeneratedFiles, f)
		event.PreAppLoadFunctions = append(event.PreAppLoadFunctions, "router.PrepareRouterYaml")
		event.ImportCachePackages = append(event.ImportCachePackages, "router")
	}
}

func (ci *cacheYam) populateRouterMetadataWithYaml() *cacheYam {
	dir, err := os.Getwd()
	p, err := ioutil.ReadFile(dir + "/conf/routes.yaml")
	if err != nil {
		panic(err)
	}

	//bug, unmarshall override checkPattern
	// may be set in cache ?
	r := Instance()
	err = yaml.Unmarshal([]byte(p), r)
	if err != nil {
		panic(err)
	}

	for name, route := range r.routerMetadata.Routes {
		if "" == strings.TrimSpace(route.Controller()) || "" == strings.TrimSpace(route.Path()) || "" == strings.TrimSpace(route.Action()) {
			continue
		}
		route.SetName(name)
		r.routerMetadata.Routes[name] = route
	}

	return ci
}

func (ci *cacheYam) writeCache() string {
	imps := make([]string, 0)
	r := Instance()
	if 0 < len(r.routerMetadata.Routes) {
		imps = append(imps, path.Join(f7k.AppPath.F7kGitImportPath(), "router"))
		//imps = append(imps, path.Join(f7k.AppConfig.GetImportPath(), "ctrl"))
	}

	tmpl, err := template.ParseFiles(filepath.Join(f7k.AppPath.F7kPath(), "router", "cacheYam.go.tmpl"))
	if nil != err {
		panic(err)
	}
	w := bytes.NewBuffer(nil)
	err = tmpl.Execute(w, map[string]interface{}{
		"Imports" : imps,
		"Data" 	  : r.routerMetadata.Routes,
	})
	if nil != err {
		panic(err)
	}
	wd, _ := os.Getwd()
	dir := path.Join(wd, "cache", "router")
	_ = os.MkdirAll(dir, os.FileMode(0755))
	filename := filepath.Join(dir, "yaml.go")
	f, err := os.Create(filename)
	if err != nil {
		panic(err)
	}

	defer f.Close()
	_ = f.Chmod(os.FileMode(0755))
	_, _ = f.Write(w.Bytes())

	return filename
}
