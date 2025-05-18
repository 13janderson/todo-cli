package cmd

import (
	"errors"
	"strconv"
	"fmt"
	"regexp"
)

// Wrapper around a list string arguments
type Parser struct{
	args []string
}

func NewParser(args []string) Parser{
	return Parser{
		args,
	}
}

// Returns the result of converting the argument args[idx] to a String
func (p *Parser) GetArgString(idx int) (string, error){
	return getArg(p.args, idx, identity[string])
}

// Returns the result of converting the argument args[idx] to an integer
func (p *Parser) GetArgInt(idx int) (int, error){
	return getArg(p.args, idx, stringToInt)
}

// Takes a list of stirng arguments, an index, and a function to convert strings to the type T
// it then returns the result of applying that function on the argument at args[idx]
func getArg[T any](args []string, idx int, convert func(string) (T, error)) (T, error) { var arg T
	var err error
	if idx < len(args) {
		arg, err = convert(args[idx])
	}
	return arg, err
}

// Takes a list of stirng arguments, an index, and a function to convert strings to the type T, and also a default value
// it then returns the result of applying that function on the argument at args[idx]. In the case of errors, it returns 
// the passed in default value.
func getArgDefaultValue[T any](args []string, idx int, defaultValue T, convert func(string) (T, error)) (T, error) {
	var arg T
	var err error
	if idx < len(args) {
		arg, err = convert(args[idx])
		return arg, err
	}
	return defaultValue, err
}

func (p *Parser) GetArgDefaultString(idx int, defaultValue string) (string, error){
	return getArgDefaultValue(p.args, idx, defaultValue, identity[string])
}

func (p *Parser) GetArgDefaultInt(idx int, defaultValue int) (int, error){
	return getArgDefaultValue(p.args, idx, defaultValue, stringToInt)
}

// TODO: add a default function for this conversion
// Probably need a struct for this to make use of existing generic functions
func (p *Parser) GetArgTimeUnitString(idx int) (time string ,unit string , e error){
	// Perform regex matching on days and hours arguments
	timeArg, _ := p.GetArgString(idx)
	regex := regexp.MustCompile(`^(\d+)([hd])$`)

	var matchedTime, matchedUnit string
	var err error
	if regex.MatchString(timeArg){
		groups := regex.FindStringSubmatch(timeArg)
		matchedTime = groups[1]
		matchedUnit = groups[2]
	}else{
		err = errors.New(fmt.Sprintf("Failed to match a valid time from string %s", timeArg))
	}
	return matchedTime, matchedUnit, err
}

func identity[T any](t T) (T, error) {
	return t, nil
}

func stringToInt(s string) (int, error) {
	return strconv.Atoi(s)
}
  
