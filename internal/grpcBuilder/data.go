package grpcBuilder

import (
	"github.com/dev2choiz/f7k"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"
)

type proto struct{
		Name            string
		ProtoFile       string `yaml:"proto-file"`
		ProtoPath       string `yaml:"proto-path"`
		ProtoPackage    string `yaml:"proto-package"`
		ProtobufPath    string `yaml:"protobuf-path"`
		ProtobufName    string
		ProtobufPackage string

		Server struct {
			Struct     string    `yaml:"struct"`
			ServerPath string    `yaml:"path"`
			ServerPackage string `yaml:"package"`
		} 					     `yaml:"server"`
}

type yamlData struct {
	GrpcProtocol 	string  			`yaml:"grpc-protocol"`
	GrpcPort 		uint16  			`yaml:"grpc-port"`
	GrpcEndpoint 	string  			`yaml:"grpc-endpoint"`
	RestEndpoint 	string  			`yaml:"rest-endpoint"`
	RestPort 		uint16  			`yaml:"rest-port"`
	Protos 			map[string]*proto	`yaml:"protos"`
}

func (g *gen) readYaml() *gen {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	str, err := ioutil.ReadFile(filepath.Join(wd, "conf", "grpc.yaml"))
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(str, g.YamlData)
	if err != nil {
		panic(err)
	}

	for k, proto := range g.YamlData.Protos {
		name := strings.TrimSuffix(proto.ProtoFile, ".proto")
		g.YamlData.Protos[k].Name = name
		g.YamlData.Protos[k].ProtoPackage = f7k.AppConfig.GetImportPath() + "/" + proto.ProtoPath
		if "" == proto.ProtobufPath {
			g.YamlData.Protos[k].ProtobufPath = filepath.Join("grpc", "res", name, proto.Name + "pb")
		}

		if "" == proto.ProtobufName {
			g.YamlData.Protos[k].ProtobufName = path.Base(proto.ProtobufPath)
		}
		if "" == proto.ProtobufPackage {
			g.YamlData.Protos[k].ProtobufPackage = filepath.Join(f7k.AppConfig.GetImportPath(), g.YamlData.Protos[k].ProtobufPath)
		}

		g.YamlData.Protos[k].Server.ServerPackage = f7k.AppConfig.GetImportPath() + "/" + proto.Server.ServerPath
	}

	return g
}

type tmplData struct {
	Protos       map[string]map[string]string
	Imports      []string
	GrpcPort     int
	RestPort     int
	GrpcEndpoint string
	RestEndpoint string
}

func (t *tmplData) populate(p *proto) {
	name :=  p.Name
	t.Protos[name] = make(map[string]string)
	t.Protos[name]["Name"]            = name
	t.Protos[name]["UCFirstName"]     = strings.ToUpper(name[:1]) + name[1:]
	t.Protos[name]["PbPackageName"]   = p.ProtobufName
	t.Protos[name]["PbPackage"]       = p.ProtobufPackage
	t.Protos[name]["UCFirstProtobuf"] = strings.ToUpper(p.ProtobufName[:1]) + p.ProtobufName[1:]
	t.Protos[name]["ServerStruct"]    = p.Server.Struct
	t.Protos[name]["ServerPackage"]   = p.Server.ServerPackage
}
