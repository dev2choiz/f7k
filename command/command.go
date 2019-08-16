package command

import (
	"github.com/dev2choiz/f7k"
	"github.com/dev2choiz/f7k/interfaces"
	"github.com/spf13/cobra"
	"sync"
)

//this variable receive commands yaml data from the cache
var AppCommands []interfaces.Command

var once sync.Once

type Command struct {
	CobraCmd      *cobra.Command

	parent        interfaces.Command
	children      []interfaces.Command
}

func New() interfaces.Command {
	c := &Command{}
	c.CobraCmd = &cobra.Command {
		Use:   "default",
		Short: "Default command.",
		Long:  "Default command.",
	}

	return c
}

func (c *Command) Parent() interfaces.Command {
	return c.parent
}

func (c *Command) SetParent(parent interfaces.Command) {
	c.parent = parent
}

func (c *Command) Children() []interfaces.Command {
	return c.children
}

func (c *Command) SetChildren(children []interfaces.Command) {
	c.children = children
}

func (c *Command) GetCobraCmd() *cobra.Command {
	return c.CobraCmd
}

func (c *Command) InitCobraCmd() {
}

func (c *Command) SetExecute(f func(cmd *cobra.Command, args []string)) {
	c.CobraCmd.Run = f
}

func (c *Command) Execute() (uint8, error) {
	c.preExecute()
	for _, child := range c.Children() {
		c.CobraCmd.AddCommand(child.GetCobraCmd())
	}

	once.Do(func() {
		c.ApplyDefaultFlags()
	})

	if e := c.CobraCmd.Execute(); nil != e {
		return 1, e
	}

	return 0, nil
}

func (c *Command) preExecute() {
}

func (c *Command) AddChild(command interfaces.Command) {
	command.SetParent(c)
	c.SetChildren(append(c.Children(), command))
}

func (c *Command) ApplyDefaultFlags() {
	c.CobraCmd.PersistentFlags().BoolVarP(&f7k.Verbose, "verbose", "v", false, "verbose output")
}

func (c *Command) IsVerbose() bool {
	b, err := c.CobraCmd.PersistentFlags().GetBool("verbose")
	if nil != err {
		return false
	}

	return b
}
