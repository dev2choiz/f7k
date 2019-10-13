package grpcBuilder

import (
	"bytes"
	"fmt"
	"github.com/dev2choiz/f7k"
	"io/ioutil"
	"os"
	"path/filepath"
	"text/template"
)

func (g *gen) createRestServer() *gen {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	srv := filepath.Join(wd, "rest")
	err = os.MkdirAll(srv, os.FileMode(0755))
	if err != nil {
		panic(err)
	}

	f := filepath.Join(srv, "main.go")


	data := tmplData{}

	data.RestPort = int(g.YamlData.RestPort)
	data.RestEndpoint = g.YamlData.RestEndpoint
	data.GrpcPort = int(g.YamlData.GrpcPort)
	data.GrpcEndpoint = g.YamlData.GrpcEndpoint
	data.Imports = append(
		data.Imports,
		"flag",
		"net/http",
		"github.com/golang/glog",
		"github.com/grpc-ecosystem/grpc-gateway/runtime",
		"golang.org/x/net/context",
		"google.golang.org/grpc",
	)

	data.Protos = make(map[string]map[string]string)
	for _, proto := range g.YamlData.Protos {
		data.populate(proto)
		data.Imports = append(data.Imports, proto.ProtobufPackage)
	}

	tmpl := template.New("rest.gohtml")
	tmpl, err = tmpl.ParseFiles(filepath.Join(f7k.AppPath.F7kPath(), "internal",  "grpcBuilder", "rest.gohtml"))
	if nil != err {
		panic(err)
	}

	w := bytes.NewBuffer(nil)
	err = tmpl.Execute(w, data)
	if nil != err {
		panic(err)
	}

	err = ioutil.WriteFile(f, w.Bytes(), 0755)
	if err != nil {
		fmt.Printf("Unable to write file: %v", err)
	}

	return g
}

