package interfaces

import (
	"github.com/spf13/cobra"
)

type Command interface {
	// @see https://github.com/spf13/cobra
	GetCobraCmd() *cobra.Command

	InitCobraCmd()
	Execute() (status uint8, err error)

	Parent()   Command
	SetParent(Command)
	Children() []Command
	SetChildren([]Command)
	AddChild(Command)
	ApplyDefaultFlags()
}
