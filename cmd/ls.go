/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"todo/todo"
	"todo/format"
	"github.com/spf13/cobra"
)

// lsCmd represents the ls command
var lsCmd = &cobra.Command{
	Use:   "ls",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			format.ShowErrorMessage("this command does not support any arguments. \n proper usage: td ls")
		}

		items, err := todo.DefaultToDoListSqlite().List()
		if err != nil{
			format.ShowWarningMessage(err.Error())
		}else{
			format.ShowToDoListItemsNormalised(items)
		}
	},
}

func init() {
	rootCmd.AddCommand(lsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// lsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// lsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
