package cmd

import (
	"github.com/spf13/cobra"
	"todo/format"
)

type ToDoCommand struct {
	cmd *cobra.Command
	pre func(args ...string) error
	// Normal verison of the command
	normal func(args ...string)
	// Recursive version of the command
	recursive func(args ...string)
	// Help displayed for recursive flag
	recursiveFlagString string
}

const RECURSIVE_FLAG = "recursive"

func NewToDoCommand(rec ToDoCommand) *cobra.Command {
	var cmd = rec.cmd
	var recursiveFlagString = "recursively run this command"
	if rec.recursiveFlagString != "" {
		recursiveFlagString = rec.recursiveFlagString
	}

	if rec.recursive != nil {
		cmd.PersistentFlags().BoolP(RECURSIVE_FLAG, "r", false, recursiveFlagString)
	}

	cmd.Run = func(cmd *cobra.Command, args []string) {
		var err = rec.pre()
		if err != nil {
			format.ShowErrorMessage(err.Error())
		}
		recursive, _ := cmd.Flags().GetBool(RECURSIVE_FLAG)
		if recursive {
			rec.recursive(args...)
		} else {
			rec.normal(args...)
		}

	}
	return rec.cmd
}
