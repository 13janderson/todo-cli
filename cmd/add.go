/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
	"todo/format"
	"todo/todo"

	"github.com/spf13/cobra"
)

const DEFAULT_TO_DO_TIME = "2"
const DEFAULT_TO_DO_UNIT = "h"
const DEFAULT_TO_DO_TIMEUNIT = DEFAULT_TO_DO_TIME + DEFAULT_TO_DO_UNIT

var addCmd = NewToDoCommand(
	ToDoCommand{
		cmd: &cobra.Command{
			Use: "add",
			Example: strings.Join([]string{
				"td add foo adds a new entry called foo with the default duration",
				"td add bar 1h adds a new entry called bar with a duration of 1 hour",
				"td add bar 2d adds a new entry called bar with a duration of 2 days",
			}, "\n"),
		},
		pre: func(args ...string) error {
			if len(args) == 0 {
				return errors.New("this command requires at least one argument. \n proper usage: td add x ?(d/h). For example td add 'have a pint' 1d to give yourself a day to have a pint")
			}
			return nil
		},
		recursive: false,
		run: func(_ bool, args ...string) {
			parser := NewParser(args)

			do, err := parser.GetArgString(0)
			if err != nil {
				// Fatal if we cannot resolve a string from the first argument
				format.ShowErrorMessage("could not parse string for toDo task")
				return
			}

			createdAt := time.Now()
			matchedTime, matchedUnit, err := parser.GetArgTimeUnitString(1)
			if err != nil {
				format.ShowWarningMessage(fmt.Sprintf("failed to parse a time from command, using default %s", DEFAULT_TO_DO_TIMEUNIT))
				matchedTime, matchedUnit = DEFAULT_TO_DO_TIME, DEFAULT_TO_DO_UNIT
			}

			var doBy time.Time
			intTime, _ := strconv.Atoi(matchedTime)
			if matchedUnit == "h" {
				doBy = createdAt.Add(time.Hour * time.Duration(intTime))
			} else {
				doBy = createdAt.AddDate(0, 0, intTime)
			}

			td := todo.ToDoListItem{
				Do:        do,
				DoBy:      doBy,
				CreatedAt: createdAt,
			}

			err = todo.DefaultToDoListSqlite().Add(&td)

			if err != nil {
				format.ShowErrorMessage(err.Error())
			} else {
				format.ShowSuccessMessage(fmt.Sprintf("Added %s", td.String()))
			}
		},
	})

func init() {
	rootCmd.AddCommand(addCmd)
}
