package cmd

import (
	// "fmt"
	"os"
	"path/filepath"
	"todo/format"

	"github.com/spf13/cobra"
)

const MAX_DEPTH = 3

type ToDoCommand struct {
	cmd *cobra.Command
	pre func(args ...string) error
	// Normal verison of the command
	run func(args ...string)
	// Recursive version of the command
	recursive bool
	// Help displayed for recursive flag
	recursiveFlagString string
}

const RECURSIVE_FLAG = "recursive"

func NewToDoCommand(toDoCommand ToDoCommand) *cobra.Command {
	var cmd = toDoCommand.cmd
	var recursiveFlagString = "recursively run this command"
	if toDoCommand.recursiveFlagString != "" {
		recursiveFlagString = toDoCommand.recursiveFlagString
	}

	if toDoCommand.recursive {
		cmd.PersistentFlags().BoolP(RECURSIVE_FLAG, "r", false, recursiveFlagString)
	}

	cmd.Run = func(cmd *cobra.Command, args []string) {
		var err = toDoCommand.pre()
		if err != nil {
			format.ShowErrorMessage(err.Error())
		}
		recursive, _ := cmd.Flags().GetBool(RECURSIVE_FLAG)
		if recursive {
			// Recursively run the command
			RunRecursive(FnArgs{
				fn:   toDoCommand.run,
				args: args,
			}, 0)
		} else {
			toDoCommand.run(args...)
		}

	}
	return toDoCommand.cmd
}

type FnArgs struct {
	fn   func(args ...string)
	args []string
}

func (fnArgs FnArgs) Call() {
	fnArgs.fn(fnArgs.args...)
}

func RunRecursive(fnArgs FnArgs, depth int) {
	// fmt.Printf("depth: %d", depth)
	// Call the function
	fnArgs.Call()

	if depth == MAX_DEPTH {
		// fmt.Printf("RETURNING")
		return
	}

	// Call the function recursively across directories
	cwd, _ := os.Getwd()
	dirs, _ := os.ReadDir(cwd)
	// fmt.Printf("dirs, %s\n", dirs)
	for _, dir := range dirs {
		if dir.IsDir() {
			dirPath := filepath.Join(cwd, dir.Name())
			// Need to change the directory
			// fmt.Printf("chdir: %s", cwd)
			os.Chdir(dirPath)
			RunRecursive(fnArgs, depth+1)
			// fmt.Println("chdir: ../")
			os.Chdir("../")
		}
	}
}
