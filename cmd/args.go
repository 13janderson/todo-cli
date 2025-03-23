package cmd

import (
	"log"
	"strconv"
)

func GetArgString(args []string, idx int, defaultValue string) string {
	return GetArg(args, idx, defaultValue, Identity[string])
}

func GetArg[T any](args []string, idx int, defaultValue T, convert func(string) (T, error)) T {
	var err error
	var arg T
	if idx < len(args) {
		arg, err = convert(args[idx])
		if err != nil {
			log.Fatal(err.Error())
		}
		return arg
	}
	return defaultValue
}

func Identity[T any](t T) (T, error) {
	return t, nil
}

func StringToInt(s string) (int, error) {
	return strconv.Atoi(s)
}
