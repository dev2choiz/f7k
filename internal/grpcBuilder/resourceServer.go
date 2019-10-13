package grpcBuilder

import (
	"bytes"
	"fmt"
	"github.com/dev2choiz/f7k"
	"github.com/dev2choiz/f7k/pkg/goParser"
	"github.com/dev2choiz/f7k/pkg/prompt"
	"io/ioutil"
	"os"
	"path/filepath"
	"text/template"
)

func (g *gen) createResourceServer() *gen {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	for _, proto := range g.YamlData.Protos {
		srv := filepath.Join(wd, proto.Server.ServerPath)
		if g.fileExists(srv) {
			if goParser.Instance().IsStructExist(srv, proto.Server.Struct) {
				prompt.New("info").Printfln("%s struct already exist in %s", proto.Server.Struct, srv)
				continue
			}
		} else {
			err = os.MkdirAll(srv, os.FileMode(0755))
			if err != nil {
				panic(err)
			}
		}

		data := tmplData{}
		data.RestPort = int(g.YamlData.RestPort)
		data.RestEndpoint = g.YamlData.RestEndpoint
		data.GrpcPort = int(g.YamlData.GrpcPort)
		data.GrpcEndpoint = g.YamlData.GrpcEndpoint
		data.Protos = make(map[string]map[string]string)
		data.populate(proto)

		tmpl := template.New("resourceServer.gohtml")
		tmpl, err = tmpl.ParseFiles(filepath.Join(f7k.AppPath.F7kPath(), "internal",  "grpcBuilder", "resourceServer.gohtml"))
		if nil != err {
			panic(err)
		}

		w := bytes.NewBuffer(nil)
		err = tmpl.Execute(w, data)
		if nil != err {
			panic(err)
		}

		f := filepath.Join(srv, "server.go")
		err = ioutil.WriteFile(f, w.Bytes(), 0755)
		if err != nil {
			fmt.Printf("Unable to write file: %v", err)
		}

		prompt.New("info").Printfln("%s generated", f)
	}

	return g
}

func (g *gen) fileExists(filename string) bool {
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}

	return true
}

const RESOURCE_SERVER_TMPL = `// Right here your protobuf server who will be registred with grpc server
package server

/*import (
	"%s"
)*/

// Should implement ServiceServer
type %s struct {
}

`
