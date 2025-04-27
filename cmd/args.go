package cmd

import (
	"errors"
	"strconv"
	"fmt"
	"regexp"
)


func GetArgString(args []string, idx int) (string, error){
	return GetArg(args, idx, Identity[string])
}

func GetArgInt(args []string, idx int) (int, error){
	return GetArg(args, idx, StringToInt)
}

func GetArg[T any](args []string, idx int, convert func(string) (T, error)) (T, error) {
	var arg T
	var err error
	if idx < len(args) {
		arg, err = convert(args[idx])
	}
	return arg, err
}

func GetArgDefaultValue[T any](args []string, idx int, defaultValue T, convert func(string) (T, error)) (T, error) {
	var arg T
	var err error
	if idx < len(args) {
		arg, err = convert(args[idx])
		return arg, err
	}
	return defaultValue, err
}

func GetArgDefaultString(args []string, idx int, defaultValue string) (string, error){
	return GetArgDefaultValue(args, idx, defaultValue, Identity[string])
}

func GetArgDefaultInt(args []string, idx int, defaultValue int) (int, error){
	return GetArgDefaultValue(args, idx, defaultValue, StringToInt)
}

// TO DO: make this conversion compatible with a default.
// Might need a type T which is a struct which itself contains two strings?
func GetArgTimeUnitString(args []string, idx int) (time string , unit string , e error){
	// Perform regex matching on days and hours arguments
	timeArg, _ := GetArgString(args, idx)
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

func Identity[T any](t T) (T, error) {
	return t, nil
}

func StringToInt(s string) (int, error) {
	return strconv.Atoi(s)
}
