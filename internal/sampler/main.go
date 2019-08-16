package main

import (
	"github.com/dev2choiz/f7k/internal/sampler/sampler"
	"os"
	"os/exec"
	"path"
)

//go:generate go run main.go
//go:generate go-bindata -pkg installer -o ./../createProject/installer/sampleAsset.go ./../createProject/sample/...
func main() {
	wd, _ := os.Getwd()
	f := path.Join(wd, "../createProject/installer/sampleAsset.go")
	_ = exec.Command("rm", "-rf", f).Run()
	f = path.Join(wd, "../createProject/sample")
	_ = exec.Command("rm", "-rf", f).Run()

	sampler.New().Execute()
}
