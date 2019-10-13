package grpcBuilder

import (
	"fmt"
	"github.com/dev2choiz/f7k/pkg/prompt"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"
)

type gen struct {
	ProtoFiles    map[string]map[string]string
	YamlData      *yamlData
}

func New() *gen {
	return &gen{}
}

func (g *gen) Execute() {
	g.ProtoFiles = make(map[string]map[string]string)
	g.YamlData = &yamlData{}

	g.
		readYaml().
		createGoProtobufs().
		createResourceServer().
		createServer().
		createRestServer()

	return
}

func (g *gen) createGoProtobufs() *gen {
	for _, p := range g.YamlData.Protos {
		// genereate protobuf
		err := generateProtobuf(p)
		if err != nil {
			panic(err)
		}
		// genereate gateway
		err = generateGateway(p)
		if err != nil {
			panic(err)
		}
		// genereate swagger
		err = generateSwagger(p)
		if err != nil {
			panic(err)
		}
	}

	return g
}

func generateProtobuf(p *proto) error {
	return execProtoc(p, "go_out=plugins=grpc",  p.ProtobufPath,",pb.go")
}

func generateGateway(p *proto) error {
	return execProtoc(p, "grpc-gateway_out=logtostderr=true", p.ProtobufPath,",pb,gw.go")
}

func generateSwagger(p *proto) error {
	return execProtoc(p, "swagger_out=logtostderr=true", "api/openapi-spec", ",swagger.json")
}

func execProtoc(p *proto, argName, output, ext string) error {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	_, err = os.Stat(path.Join(wd, output))
	if os.IsNotExist(err) {
		err = os.MkdirAll(path.Join(wd, output), os.FileMode(0777))
		if err != nil {
			panic(err)
		}
	} else if err != nil {
		panic(err)
	}

	imp := []string{"-I" + p.ProtoPath}
	cmdStr := fmt.Sprintf(
		"-Ithird_party %s --%s:%s %s",
		strings.Join(imp, " "),
		argName,
		output,
		p.ProtoFile)

	cmd := strings.Split(cmdStr, " ")
	o, err := exec.Command("protoc", cmd...).CombinedOutput()
	if nil != err {
		return fmt.Errorf("%s\nCommand :protoc %s\nCommand output:\n%s", err.Error(), cmdStr, string(o))
	}

	prompt.New("info").Println(fmt.Sprintf(
		"%s generated",
		filepath.Join(wd, output, p.Name + ext)))

	return nil
}
