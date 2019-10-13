package grpcBuilder

import (
	"bytes"
	"fmt"
	"github.com/dev2choiz/f7k"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

func (g *gen) createServer() *gen {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	srv := filepath.Join(wd, "grpc", "server")
	err = os.MkdirAll(srv, os.FileMode(0755))
	if err != nil {
		panic(err)
	}

	f := filepath.Join(srv, "server.go")

	data := struct{
		Data map[string]map[string]string
		Imports []string
		Port int
		Protocol string
	}{}

	data.Port 	  = int(g.YamlData.GrpcPort)
	data.Protocol = g.YamlData.GrpcProtocol
	data.Imports = append(data.Imports, "\"fmt\"", "\"net\"", "\"google.golang.org/grpc\"")
	data.Data = make(map[string]map[string]string)
	for _, proto := range g.YamlData.Protos {
		name := proto.Name
		data.Data[name] = make(map[string]string)
		data.Data[name]["Name"]          = name
		data.Data[name]["PbPackageName"] = proto.ProtobufName
		data.Data[name]["PbPackage"]     = proto.ProtoPackage
		data.Data[name]["UCFirstName"]   = strings.ToUpper(name[:1]) + name[1:]
		data.Data[name]["ServerStruct"]   = proto.Server.Struct
		data.Data[name]["AliasServerPackage"]   = getAliasPachage(proto)
		data.Imports = append(data.Imports, "\"" + proto.ProtobufPackage + "\"")
		data.Imports = append(data.Imports, getFormattedImport(proto))
	}

	tmpl := template.New("server.gohtml")
	tmpl, err = tmpl.ParseFiles(filepath.Join(f7k.AppPath.F7kPath(), "internal",  "grpcBuilder", "server.gohtml"))
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

func getAliasPachage(p *proto) string {
	s := strings.Split(p.Server.ServerPath, "/")
	for k, path := range s {
		if 0 != k {
			s[k] = strings.ToUpper(path[:1]) + path[1:]
		}
	}

	return strings.Join(s, "")
}

func getFormattedImport(p *proto) string {
	return fmt.Sprintf("%s \"%s\"", getAliasPachage(p), p.Server.ServerPackage)
}