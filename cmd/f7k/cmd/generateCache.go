package cmd

import (
	"github.com/dev2choiz/f7k/appLoader"
	"github.com/dev2choiz/f7k/cacheGen"
	"github.com/dev2choiz/f7k/command"
	"github.com/dev2choiz/f7k/configurator"
	"github.com/dev2choiz/f7k/interfaces"
	"github.com/dev2choiz/f7k/pkg/prompt"
	"github.com/spf13/cobra"
	"sync"
)

var once sync.Once

var cacheCmd *command.Command

func GenerateCacheCmd() interfaces.Command {
	once.Do(func() {
		cacheCmd = &command.Command{}
		cacheCmd.CobraCmd = &cobra.Command{
			Use:   "cache",
			Short: "Regenerate cache",
			Long:  "Regenerate application cache",
			Run: func(cmd *cobra.Command, args []string) {
				cacheExec()
			},
		}
	})

	return cacheCmd
}

func cacheExec() {
	appLoader.DefaultCliServerLoader(&configurator.Config{}).LoadApp()
	cacheGen.Instance().Run()
	prompt.New("info").Println("Cache generated")
}