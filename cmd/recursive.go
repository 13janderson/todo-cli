package cmd

import (
	"github.com/spf13/cobra"
	"todo/format"
)

// If something satisfies this interface then we can construct a
// recursive command for it.
type Recursive struct {
	cmd       *cobra.Command
	pre       func(args ...string) error
	normal    func()
	recursive func()
}

func NewRecursiveCommand(rec Recursive, args ...any) *cobra.Command {
	var cmd = rec.cmd
	cmd.Run = func(cmd *cobra.Command, args []string) {
		var err = rec.pre()
		if err != nil {
			format.ShowErrorMessage(err.Error())
		}
		recursive, _ := cmd.Flags().GetBool("recursive")
		if recursive {
			rec.recursive()
		} else {
			rec.normal()
		}

	}
	return rec.cmd

}
