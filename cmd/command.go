package cmd

import (
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
	"todo/format"
)

const MAX_DEPTH = 3

type FnArgs struct {
	fn             func(additionalArgs AdditionalArgs, args ...string)
	additionalArgs AdditionalArgs
	args           []string
}

type AdditionalArgs struct {
	recursive bool
	depth     int
}

type ToDoCommand struct {
	cmd *cobra.Command
	pre func(args ...string) error
	// What to run for each invocation of this command, this is ran recursively
	// if recursive is passed as true
	run func(additionalArgs AdditionalArgs, args ...string)
	// Whether this command supports recursive use over directories
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
		var err = toDoCommand.pre(args...)
		if err != nil {
			format.ShowErrorMessage(err.Error())
			return
		}
		recursive, err := cmd.Flags().GetBool(RECURSIVE_FLAG)
		if recursive {
			// Recursively run the command
			RunRecursive(FnArgs{
				additionalArgs: AdditionalArgs{
					recursive: true,
				},
				fn:   toDoCommand.run,
				args: args,
			}, 0)
		} else if !recursive || err != nil {
			toDoCommand.run(AdditionalArgs{
				recursive: false,
			}, args...)
		}

	}
	return toDoCommand.cmd
}

func (fnArgs FnArgs) Call() {
	fnArgs.fn(fnArgs.additionalArgs, fnArgs.args...)
}

func RunRecursive(fnArgs FnArgs, depth int) {
	// fmt.Printf("depth: %d", depth)
	// Call the function
	fnArgs.Call()

	if depth == MAX_DEPTH {
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
			fnArgs.additionalArgs.depth = depth
			RunRecursive(fnArgs, depth+1)
			// fmt.Println("chdir: ../")
			os.Chdir("../")
		}
	}
}
