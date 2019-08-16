package cmd

import (
	"fmt"
	"github.com/dev2choiz/f7k"
	"github.com/dev2choiz/f7k/appLoader"
	"github.com/dev2choiz/f7k/command"
	"github.com/dev2choiz/f7k/configurator"
	"github.com/dev2choiz/f7k/interfaces"
	"github.com/dev2choiz/f7k/internal/createProject/installer"
	"github.com/dev2choiz/f7k/pkg/prompt"
	"github.com/spf13/cobra"
	"os"
	"path"
	"strconv"
	"strings"
)

func CreateProjectCmd() interfaces.Command {
	c := &command.Command{}
	c.CobraCmd = &cobra.Command{
		Use:   "create",
		Short: "Create a new project",
		Long:  fmt.Sprintf(`Create a new golang web project based on "%s" framework`, f7k.PrettyName),
		Run: func(cmd *cobra.Command, args []string) {
			createProjectExec(cmd, args)
		},
	}

	c.CobraCmd.PersistentFlags().StringP( "name", "n", "", "Application name")
	c.CobraCmd.PersistentFlags().IntP( "port", "p", 0, "listen port")
	c.CobraCmd.PersistentFlags().StringP("import-path", "i", "", "Import path (ex : github.com/dev2choiz/f7k)")

	return c
}

func createProjectExec(cmd *cobra.Command, args []string) {
	appLoader.
		GetCliServerLoader(
			&configurator.Config{},
			"",
			"",
			false,
			false).
		Load()

	prompt.New("info").Println("Create application...")
	i := installer.New()
	i.CurrentDir, _ = os.Getwd()
	i.ProjectName = askProjectName(cmd, args)
	i.AppImportPath = askImportPath(cmd, args)
	i.Port = uint16(askPort(cmd, args))

	i.Execute()
	prompt.New("info").Printfln("Application created in \"%s\".", path.Join(i.CurrentDir, i.ProjectName))
}

func askProjectName(cmd *cobra.Command, args []string) string {
	n, err := cmd.PersistentFlags().GetString("name")
	if nil != err || "" == n {
		n, _ = ask("Project name:")
	}
	return strings.ToLower(n)
}

func askImportPath(cmd *cobra.Command, args []string) string {
	n, err := cmd.PersistentFlags().GetString("import-path")
	if nil != err || "" == n {
		n, _ = ask("import path (ex : github.com/dev2choiz/f7k) :")
	}
	return n
}

func askPort(cmd *cobra.Command, args []string) int {
	n, err := cmd.PersistentFlags().GetInt("port")
	if nil == err && 0 != n {
		return n
	}

	var i int
	for {
		s, _ := ask("Port :")
		i, err = strconv.Atoi(s)
		if nil == err {
			break
		}
		prompt.New("danger").Printfln("given port %s is invalid", s)
	}

	return i
}
