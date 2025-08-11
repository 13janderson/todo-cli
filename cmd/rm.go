/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"todo/format"
	"todo/todo"
)

var rmCmd = NewToDoCommand(ToDoCommand{
	cmd: &cobra.Command{
		Use: "rm",
	},
	recursive: true,
	pre: func(args ...string) error {
		if len(args) > 1 {
			return errors.New("this command cannot take more than one argument. \n proper usage: rm x where x is either an id or a string pattern to match with")
		}

		return nil
	},
	run: func(_ bool, args ...string) {
		parser := NewParser(args)
		id, err := parser.GetArgInt(0)

		var toDoListItem todo.ToDoListItem

		// Try to parse an id from the string. Failing that we try to match with the string argument
		if err == nil {
			toDoListItem.Id = id
		} else {
			toDoListItem.Do, _ = parser.GetArgString(0)
		}

		deleted := todo.DefaultToDoListSqliteCwd().Remove(toDoListItem)
		format.RemovedMessage(fmt.Sprintf("Removed %d task(s)", deleted))
	},
})

func init() {
	rootCmd.AddCommand(rmCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// rmCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// rmCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
