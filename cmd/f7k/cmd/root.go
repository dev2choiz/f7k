package cmd

import (
	"bufio"
	"github.com/dev2choiz/f7k"
	"github.com/dev2choiz/f7k/appLoader"
	"github.com/dev2choiz/f7k/command"
	"github.com/dev2choiz/f7k/configurator"
	"github.com/dev2choiz/f7k/interfaces"
	"github.com/dev2choiz/f7k/pkg/prompt"
	"github.com/spf13/cobra"
	"log"
	"os"
	"strings"
)

var rootCmd interfaces.Command

func getRootCmd() interfaces.Command {
	if nil == rootCmd {
		c := &command.Command{}
		c.CobraCmd = &cobra.Command {
			Use:   f7k.Name,
			Short: f7k.PrettyName,
			Long:  f7k.PrettyName,
			Run:   func(cmd *cobra.Command, args []string) {
				appLoader.GetCliServerLoader(&configurator.Config{}, "", "",true,  false).Load()
				if err := cmd.Help(); nil != err {
					panic(err)
				}
				os.Exit(0)
			},
		}

		c.AddChild(RunCmd())
		c.AddChild(CreateProjectCmd())
		c.AddChild(GenerateCacheCmd())
		c.AddChild(GrpcBuildCmd())
		c.AddChild(VersionCmd())

		rootCmd = c
	}

	return rootCmd
}

func Execute() (status uint8, err error) {
	return getRootCmd().Execute()
}

func ask(msg string) (string, error) {
	prompt.New("info").Println(msg)
	n, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	n = strings.TrimSpace(n)

	return n, nil
}

func checkFatal(e error) {
	if nil != e {
		log.Fatalf("An error occured : " + e.Error())
	}
}
