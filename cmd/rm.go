/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"errors"
	"github.com/fatih/color"
	"strconv"
	"todo/todo"
	"github.com/spf13/cobra"
)

// rmCmd represents the rm command
var rmCmd = &cobra.Command{
	Use:   "rm",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: func(cmd *cobra.Command, args []string) error{
		if len(args) > 1{
			return errors.New("this command requires atr least one argument. \n proper usage: rm x where x is either an id or a string to match with")
		}else if len(args) == 0{

			return errors.New("this command requires at least one argument. \n proper usage: td add x ?d ?h")
		}else{
			arg := args[0]
			id, err := strconv.Atoi(arg)

			var toDoListItem todo.ToDoListItem;

			// Try to parse an id from the string. Failing that we try to match with the Do string 
			if err == nil{
				toDoListItem.Id = id
			}else{
				toDoListItem.Do = arg
			}
			
			deleted := todo.DefaultToDoListSqlite().Remove(toDoListItem)
			color.Set(color.Bold)
			color.Red("%s Removed %d task(s)", indent(), deleted)
			return nil
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
