package main

import (
	// "crypto/sha256"
	"encoding/json"
	"fmt"
	"hash"
	"os"
)

// Need a simple way to keep track of to do's in specific repos or directories
// Question of what the best way to store these is, simple JSON file with a default TTL is the first thought

type ToDoListItem struct {
	Do      string    `json:"do"`
	TTLDays int       `json:"by"`
	Hash hash.Hash `json:"hash"`
}

type ToDoList interface {
	// Init new to do list
	Init()
	Add(item ToDoListItem)
	// Remove specific item
	Remove(item ToDoListItem)
	// Removes most recent to do added
	Pop()
}

type ToDoListJson struct {
	jsonFileName string
	jsonFile     *os.File
}

func NewToDoListJson() (td *ToDoListJson) {
	td = &ToDoListJson{
		jsonFileName: ".td.json",
	}
	td.Init()
	return td
}

// In the current directory, initialise a new json to do list
// if one already exists, do nothing
func (td *ToDoListJson) Init() {
	file, err := os.OpenFile(td.jsonFileName, os.O_CREATE | os.O_APPEND | os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("Error opening file", err.Error())
		return
	}
	fmt.Println("Initialised")
	td.jsonFile = file
}

func (td *ToDoListJson) Add(item ToDoListItem) {
	bytes, _ := json.Marshal(item)
	n, err := td.jsonFile.Write(bytes)
	if(err != nil){
		fmt.Println("Error writing file", err.Error())
	}
	fmt.Println(n)
	b := make([]byte, 10)
	td.jsonFile.Read(b)
	fmt.Println(string(b))
}

func (td *ToDoListJson) Remove(item ToDoListItem) {

}

func (td *ToDoListJson) Pop() {

}

func main() {
	td := NewToDoListJson()
	td.Add(ToDoListItem{
		Do: "abc",
	})
}
