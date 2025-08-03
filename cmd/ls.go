/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"errors"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
	"todo/format"
	"todo/todo"
)

var lsCmd = NewRecursiveCommand(Recursive{
	cmd: &cobra.Command{
		Use: "ls",
	},
	pre: func(args ...string) error {
		if len(args) > 0 {
			return errors.New("this command does not support any arguments. \n proper usage: td ls")
		}
		return nil
	},
	recursive: func() { showListInDirectoryRecursive(0, ".") },
	normal:    func() { showList() },
})

const MAX_DEPTH = 3

func showListInDirectoryRecursive(currentDepth int, directory string) {

	showListDirectory(directory)
	if currentDepth == MAX_DEPTH {
		return
	}

	dirs, _ := os.ReadDir(directory)
	for _, dir := range dirs {
		if dir.IsDir() {
			dirPath := filepath.Join(directory, dir.Name())
			showListInDirectoryRecursive(currentDepth+1, dirPath)
		}
	}
}

func showListDirectory(directory string) {
	items, err := todo.DefaultToDoListSqliteInDirectory(directory).List()
	if err == nil {
		format.ShowDirectoryMessage(directory)
		format.ShowToDoListItems(items)
	}
}

func showList() {
	items, err := todo.DefaultToDoListSqlite().List()
	if err != nil {
		format.ShowWarningMessage(err.Error())
	} else {
		format.ShowToDoListItems(items)
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
	lsCmd.PersistentFlags().BoolP("recursive", "r", false, "Recursive listing with max depth of 5.")
}
