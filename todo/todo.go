package todo

import (
	"log"
	"strconv"
	"time"
)

type ToDoListItem struct {
	Id        int       `json:"id" db:"id"`
	Do        string    `json:"do" db:"do"`
	DoBy time.Time `json:"doBy" db:"doBy"`
}

type ToDoList interface {
	// Init new to do list
	Init()
	// Writes all list items in readable format to stdout
	List() error
	Add(item ToDoListItem)
	// Remove specific item
	Remove(item ToDoListItem)
	// Removes most recent to do added
	Pop()

	Complete()
}


func GetArgString (args []string, idx int, defaultValue string) string {
	return GetArg(args, idx, defaultValue, Identity[string])
}

func GetArg[T any] (args []string, idx int, defaultValue T, convert func(string) (T, error)) T{
	var err error
	var arg T
	if idx < len(args) {
		arg, err = convert(args[idx])
		if err != nil{
			log.Fatal(err.Error())
		}
		return arg
	}
	return defaultValue
}

func Identity[T any](t T) (T, error ){
	return t, nil
}

func StringToInt(s string) (int, error) {
	return strconv.Atoi(s)
}
