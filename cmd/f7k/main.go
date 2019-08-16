package main

import (
	"github.com/dev2choiz/f7k/cmd/f7k/cmd"
	"github.com/dev2choiz/f7k/pkg/prompt"
	"os"
)

func main() {
	s, e := cmd.Execute()
	if nil != e {
		prompt.New("danger").Printfln("error : ", e.Error())
		os.Exit(int(s))
	}
	os.Exit(0)
}
