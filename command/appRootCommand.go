package command

import (
	"github.com/dev2choiz/f7k/interfaces"
	"github.com/spf13/cobra"
)

type AppRootCommand struct {
	*Command
}

func NewRootCommand(cmds []interfaces.Command, f func()) *AppRootCommand {
	c := &AppRootCommand{&Command{}}
	c.CobraCmd = &cobra.Command {
		Use:   "main",
		Short: "Run application",
		Long:  "Run application",
		Run: func(cmd *cobra.Command, args []string) {
			f()
		},
	}

	for k := range cmds {
		c.AddChild(cmds[k])
	}

	return c
}

func (c *AppRootCommand) Run() (status uint8, err error) {
	status, err = c.Execute()
	if nil != err {
		panic(err)
	}

	return status, err
}
