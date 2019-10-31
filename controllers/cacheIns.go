package controllers

import (
	"bytes"
	"github.com/dev2choiz/f7k"
	"github.com/dev2choiz/f7k/cacheGen"
	"github.com/dev2choiz/f7k/model/events"
	"github.com/dev2choiz/f7k/pkg/goParser"
	"go/ast"
	"html/template"
	"os"
	"path"
	"path/filepath"
)

type cacheIns struct {
	*cacheGen.CacheListener
}

func OnCacheGenInstantiate(event *events.CacheGenEvent) {
	cr := RegistryInstance()
	cr.Imports = make([]string, 0) // re-init imports
	ci := cacheIns{cacheGen.NewCacheListener()}
	f := ci.populateMetadata().writeCache()
	if "" != f {
		event.GeneratedFiles = append(event.GeneratedFiles, f)
		event.PreAppLoadFunctions = append(event.PreAppLoadFunctions, "controller.PrepareControllerLoader")
		event.ImportCachePackages = append(event.ImportCachePackages, "controller")
	}
}

func (ci *cacheIns) populateMetadata() *cacheIns {
	dir := path.Join(f7k.AppPath.AppRoot(), f7k.AppConfig.GetControllerDir())
	specs, err := goParser.Instance().ExtractStructsSpecs(dir)
	if nil != err {
		panic(err)
	}
	RegistryInstance().ControllersSpecs = ci.filterSpec(specs)

	return ci
}

func (ci *cacheIns) filterSpec(specs []*ast.TypeSpec) map[string]*ast.TypeSpec {
	newSpecs := make(map[string]*ast.TypeSpec, 0)
	for _, spec := range specs {
		typ := spec.Type.(*ast.StructType)
		for _, field := range typ.Fields.List {
			if nil != field.Names {
				continue	// because we are searching a unnamed embedded struct
			}
			fieldType := field.Type.(*ast.SelectorExpr)
			if "controllers" != fieldType.X.(*ast.Ident).Name {
				continue
			}
			if "Controller" != fieldType.Sel.Name {
				continue
			}

			// Here the struct is embedding controllers.Controller
			newSpecs[f7k.AppConfig.GetControllerDir() + "." + spec.Name.Name] = spec
		}
	}

	return newSpecs
}

func (ci *cacheIns) writeCache() string {
	cr := RegistryInstance()
	if 0 < len(cr.ControllersSpecs) {
		cr.Imports = append(cr.Imports, path.Join(f7k.AppPath.F7kGitImportPath(), "controllers"))
		cr.Imports = append(cr.Imports, path.Join(f7k.AppConfig.GetImportPath(), "ctrl"))
	}

	tmpl, err := template.ParseFiles(filepath.Join(f7k.AppPath.F7kPath(), "controllers", "cacheIns.go.tmpl"))
	if nil != err {
		panic(err)
	}
	w := bytes.NewBuffer(nil)
	err = tmpl.Execute(w, cr)
	if nil != err {
		panic(err)
	}
	wd, _ := os.Getwd()
	dir := path.Join(wd, "cache", "controller")
	_ = os.MkdirAll(dir, os.FileMode(0755))
	filename := filepath.Join(dir, "instances.go")
	f, err := os.Create(filename)
	if err != nil {
		panic(err)
	}

	defer f.Close()
	_ = f.Chmod(os.FileMode(0755))
	_, _ = f.Write(w.Bytes())

	return filename
}
