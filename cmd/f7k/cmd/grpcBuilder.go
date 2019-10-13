package cmd

import (
	"github.com/dev2choiz/f7k/appLoader"
	"github.com/dev2choiz/f7k/command"
	"github.com/dev2choiz/f7k/configurator"
	"github.com/dev2choiz/f7k/interfaces"
	"github.com/dev2choiz/f7k/internal/grpcBuilder"
	"github.com/dev2choiz/f7k/pkg/prompt"
	"github.com/spf13/cobra"
)

func GrpcBuildCmd() interfaces.Command {
	c := &command.Command{}
	c.CobraCmd = &cobra.Command{
		Use:   "grpc-build",
		Short: "Generate grpc and rest files",
		Long:  "Generate grpc and rest files according grpc.yaml file.",
		Run: func(cmd *cobra.Command, args []string) {
			grpcBuildExec(cmd, args)
		},
	}

	return c
}

func grpcBuildExec(cmd *cobra.Command, args []string) {
	appLoader.DefaultCliServerLoader(&configurator.Config{}).LoadApp()
	prompt.New("info").Println("grpc build ...")
	grpcBuilder.New().Execute()
}
