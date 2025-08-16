/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"errors"
	"fmt"
	"todo/format"
	"todo/todo"

	"github.com/spf13/cobra"
)

var lsCmd = NewToDoCommand(ToDoCommand{
	cmd: &cobra.Command{
		Use: "ls",
	},
	pre: func(args ...string) error {
		if len(args) > 0 {
			return errors.New("this command does not support any arguments. \n proper usage: td ls")
		}
		return nil
	},
	run:                 showList,
	recursive:           true,
	recursiveFlagString: fmt.Sprintf("recursive listing with max depth of %d.", MAX_DEPTH),
})

func showList(args AdditionalArgs, _ ...string) {
	items, err := todo.DefaultToDoListSqliteCwd().List()
	// We only want to show the warning that the list was not initialised if we are not running this
	// command recursively
	if err != nil {
		if !args.recursive {
			format.ShowWarningMessage(err.Error())
		}
	} else {
		depth := args.depth
		format.ShowCwdMessage(depth)
		format.ShowToDoListItems(items, depth)
	}
}

func init() {
	rootCmd.AddCommand(lsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// lsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
}
