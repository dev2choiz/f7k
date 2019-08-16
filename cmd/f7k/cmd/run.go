package cmd

import (
	"bufio"
	"fmt"
	"github.com/dev2choiz/f7k/command"
	"github.com/dev2choiz/f7k/interfaces"
	"github.com/dev2choiz/f7k/pkg/prompt"
	"github.com/spf13/cobra"
	"os"
	"os/exec"
	"strings"
)

func RunCmd() interfaces.Command {
	c := &command.Command{}
	c.CobraCmd = &cobra.Command{
		Use:   "run",
		Short: "Run the project",
		Long:  "Run the project",
		Run: func(cmd *cobra.Command, args []string) {
			runExec()
		},
	}

	return c
}

func runExec() {
	cacheExec()

	wd, _ := os.Getwd()
	str := fmt.Sprintf("run %s/main.go ", wd)
	str += strings.Split(strings.Join(os.Args, " "), " run ")[1]
	cmd := exec.Command("go", strings.Split(str, " ")...)

	stdOut, e := cmd.StdoutPipe()
	checkFatal(e)
	checkFatal(cmd.Start())
		scOut := bufio.NewScanner(stdOut)
		scOut.Split(bufio.ScanLines)
		for scOut.Scan() {
			m := scOut.Text()
			prompt.New("info").Println(m)
		}
	_ = cmd.Wait()
}
