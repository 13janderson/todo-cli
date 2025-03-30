/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"
	"fmt"
	"todo/format"
	"todo/todo"
	"path/filepath"
	"github.com/spf13/cobra"
)

// lsCmd represents the ls command
// Currently we are just listing under the current directory, what we want is to call this function recursively with a maximum depth
var lsCmd = &cobra.Command{
	Use:   "ls",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			format.ShowErrorMessage("this command does not support any arguments. \n proper usage: td ls")
		}

		recursive, _ := cmd.Flags().GetBool("recursive")
		if recursive{
			showListInDirectoryRecursive(0, ".")
		}else{
			showList()
		}
	},
}

const MAX_DEPTH = 3
func showListInDirectoryRecursive(currentDepth int, directory string){

	showListDirectory(directory)
	if currentDepth == MAX_DEPTH{
		return
	}

	dirs, _ := os.ReadDir(directory)
	for _, dir := range dirs{
		if dir.IsDir(){
			dirPath := filepath.Join(directory, dir.Name())
			showListInDirectoryRecursive(currentDepth+1, dirPath)
		}
	}
}

func showListDirectory(directory string){
	items, err := todo.DefaultToDoListSqliteInDirectory(directory).List()
	if err == nil{
		format.ShowDirectoryMessage(fmt.Sprintf("/%s", directory))
		format.ShowToDoListItemsNormalised(items)
	}
}


func showList() {
	items, err := todo.DefaultToDoListSqlite().List()
	if err != nil{
		format.ShowWarningMessage(err.Error())
	}else{
		format.ShowToDoListItemsNormalised(items)
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
