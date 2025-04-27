/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"todo/todo"
	"todo/format"
	"github.com/spf13/cobra"
)

// rmCmd represents the rm command
var rmCmd = &cobra.Command{
	Use:   "rm",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 1{
			format.ShowErrorMessage("this command cannot take more than one argument. \n proper usage: rm x where x is either an id or a string to match with")
		}else if len(args) == 0{
			format.ShowErrorMessage("this command requires at least one argument. \n proper usage: td rm x")
		}else{
			id, err := GetArgInt(args, 0)

			var toDoListItem todo.ToDoListItem;

			// Try to parse an id from the string. Failing that we try to match with the string argument
			if err == nil{
				toDoListItem.Id = id
			}else{
				toDoListItem.Do, _ = GetArgString(args, 0)
			}
			
			deleted := todo.DefaultToDoListSqlite().Remove(toDoListItem)
			format.RemovedMessage(fmt.Sprintf("Removed %d task(s)", deleted))
		}
	},
}

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
