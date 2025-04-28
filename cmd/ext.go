/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt" 
	"time"
	"todo/format"
	"todo/todo"
	"strconv"
	"github.com/spf13/cobra"
)

const DEFAULT_EXT_TIME = "1"
const DEFAULT_EXT_UNIT = "d"
const DEFAULT_EXT_TIMEUNIT = DEFAULT_EXT_TIME + DEFAULT_EXT_UNIT

// rmCmd represents the rm command
var extCmd= &cobra.Command{
	Use:   "ext",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 2{
			format.ShowErrorMessage("this command requires exactly 2 arguments. \n proper usage: td ext x ?(d/h)")
		}else{
			id, err := GetArgInt(args, 0)
			if err != nil{
				format.ShowErrorMessage("could not parse id for toDo task")
				return
			}


			
			// Search list for item with id
			itemsWithId, err := todo.DefaultToDoListSqlite().SelectWithId(id)
			if err != nil{
				format.ShowErrorMessage(err.Error())
				return
			}

			noItemsWithId := len(itemsWithId)
			if noItemsWithId == 0{
				format.ShowErrorMessage(fmt.Sprintf("could not find toDo list item with id: %d", id))
				return
			}else if noItemsWithId > 1{
				format.ShowErrorMessage(fmt.Sprintf("found multiple items with same id, this needs to be resolved manually - e.g. by deleting the list all together: %d", id))
				return
			} 

			toDoListItem := itemsWithId[0]
			doBy := toDoListItem.DoBy
			// If the item is currently expired, take the extension of date from now
			now := time.Now()
			if doBy.Before(now){
				doBy = now
			}
			matchedTime, matchedUnit, err := GetArgTimeUnitString(args, 1)
			if err != nil{
				format.ShowWarningMessage(fmt.Sprintf("failed to parse a time from command, using default %s", DEFAULT_EXT_TIMEUNIT))
				matchedTime, matchedUnit = DEFAULT_EXT_TIME, DEFAULT_EXT_UNIT
			}

			intTime, _ := strconv.Atoi(matchedTime)
			if matchedUnit == "h"{
				doBy = doBy.Add(time.Hour * time.Duration(intTime))
			}else{
				doBy = doBy.AddDate(0,0, intTime)
			}

			toDoListItem.DoBy = doBy
			err = todo.DefaultToDoListSqlite().Extend(toDoListItem)
			if err == nil{
				format.ShowSuccessMessage(toDoListItem.String())
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(extCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// rmCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// rmCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
