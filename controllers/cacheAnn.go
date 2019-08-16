package controllers

import (
	"bytes"
	"github.com/dev2choiz/f7k"
	"github.com/dev2choiz/f7k/cacheGen"
	"github.com/dev2choiz/f7k/interfaces"
	"github.com/dev2choiz/f7k/model/events"
	"github.com/dev2choiz/f7k/pkg/annotation"
	"os"
	"path"
	"path/filepath"
	"text/template"
)

type cacheAnn struct {
	*cacheGen.CacheListener
}

func OnCacheGenAnnotations(e interfaces.Event) {
	event := e.(*events.CacheGenEvent)
	ca := &cacheAnn{cacheGen.NewCacheListener()}
	f := ca.writeControllersAnnotations(e)
	if "" != f {
		event.GeneratedFiles = append(event.GeneratedFiles, f)
		event.PreAppLoadFunctions = append(event.PreAppLoadFunctions, "controller.PrepareControllerAnnotationsLoader")
		event.ImportCachePackages = append(event.ImportCachePackages, "controller")
	}
}

func (ca *cacheAnn) writeControllersAnnotations(e interfaces.Event) string {
	return ca.
		fetchControllersAnnotations().
		writeControllerCacheAnnotations()
}

func (ca *cacheAnn) fetchControllersAnnotations() *cacheAnn {
	dir := path.Join(f7k.AppPath.AppRoot(), f7k.AppConfig.GetControllerDir())
	RegistryInstance().Annotations = annotation.New().GetPackageAnnotations(dir)

	return ca
}

func (ca *cacheAnn) writeControllerCacheAnnotations() string {
	cr := RegistryInstance()
	if 0 < len(cr.Annotations) {
		cr.Imports = append(cr.Imports, path.Join(f7k.AppPath.F7kGitImportPath(), "pkg/annotation"))
		cr.Imports = append(cr.Imports, path.Join(f7k.AppPath.F7kGitImportPath(), "controllers"))
	}

	tmpl := template.New("cacheAnn.go.tmpl").Funcs(template.FuncMap{
		"sTypedPrint": cacheGen.Instance().STypedPrint,
		"escapeDoubleQuotes": cacheGen.Instance().EscapeDoubleQuotes,
	})
	tmpl, err := tmpl.ParseFiles(filepath.Join(f7k.AppPath.F7kPath(), "controllers", "cacheAnn.go.tmpl"))
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
	filename := filepath.Join(dir, "annotations.go")
	f, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	_ = f.Chmod(os.FileMode(0755))
	_, _ = f.Write(w.Bytes())

	return filename
}
