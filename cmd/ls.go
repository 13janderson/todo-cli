/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"errors"
	"time"
	"todo/todo"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// lsCmd represents the ls command
var lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) > 0 {
			return errors.New("this command does not support any arguments. \n proper usage: td ls")
		}

		items, err := todo.DefaultToDoListSqlite().List()
		if err != nil{
			return err
		}

		for _, item := range items{
			remainingTime := item.RemainingTime()
			color.Set(color.Bold)
			if remainingTime <= time.Duration(0){
				color.Red("%s [%d] %s EXPIRED", indent(), item.Id, item.Do)
			}else{
				color.Green("%s [%d] %s %s", indent(), item.Id, item.Do, DurationHumanReadable(remainingTime))
			}
		}
		return nil

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
