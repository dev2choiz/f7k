package sampler

import (
	"bytes"
	"github.com/dev2choiz/f7k/fileFiller"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

type archiver struct {
	AppImportPath string
	F7kImportPath string
	ExamplePath   string
	TargetPath    string
}

func (i archiver) Execute() {
	i.copyFiles().fillFiles()
	return
}

func New() *archiver {
	i := &archiver{}
	i.F7kImportPath = "github.com/dev2choiz/f7k"
	i.AppImportPath = "github.com/dev2choiz/f7ktest"
	i.ExamplePath   = "/var/www/golang/f7ktest"
	i.TargetPath    = "/var/www/golang/f7k/internal/createProject/sample"

	return i
}

func (i *archiver) copyFiles() *archiver {
	_ = exec.Command("rm", "-rf", i.TargetPath).Run()

	cmd := exec.Command("cp", "-r", i.ExamplePath, i.TargetPath)
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		outStr, errStr := string(stdout.Bytes()), string(stderr.Bytes())
		log.Fatalf("cannot create the project in %s : %s %s %s\n", outStr, errStr, err.Error())
	}

	_ = os.Remove(i.TargetPath + "/go.sum")
	_ = os.Remove(i.TargetPath + "/debug")
	_ = os.Remove(i.TargetPath + "/__debug_bin")
	_ = exec.Command("rm", "-rf", filepath.Join(i.TargetPath, ".idea")).Run()
	_ = exec.Command("rm", "-rf", filepath.Join(i.TargetPath, "vendor")).Run()

	return i
}

func (i *archiver) fillFiles() *archiver {
	ff := fileFiller.New()
	ff.Placeholders = map[string]string{
		 i.AppImportPath : "{{APP_IMPORT_PATH}}",
		i.F7kImportPath : "{{F7K_IMPORT_PATH}}",
	}

	ff.FillFiles(i.TargetPath)

	return i
}
