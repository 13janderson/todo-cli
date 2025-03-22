package todo

import (
	// "crypto/sha256"
	"encoding/json"
	"bufio"
	"fmt"
	"hash"
	"os"
)

// Need a simple way to keep track of to do's in specific repos or directories
// Question of what the best way to store these is, simple JSON file with a default TTL is the first thought

type ToDoListItem struct {
	Id int `json:"do" db:"id"`
	Do      string    `json:"do" db:"do"`
	ByDays int       `json:"byDays" db:"byDays"`
	ByHours int       `json:"byHours" db:"byHours"`
	Hash hash.Hash `json:"hash" db:"hash"`
}

type ToDoList interface {
	// Init new to do list
	Init()
	Add(item ToDoListItem)
	// Remove specific item
	Remove(item ToDoListItem)
	// Removes most recent to do added
	Pop()

	Complete()
}

type ToDoListJson struct {
	jsonFileName string
	jsonFile     *os.File
	jsonFileScanner *bufio.Scanner
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
	writeFile, err := os.OpenFile(td.jsonFileName, os.O_CREATE | os.O_APPEND | os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("Error opening file", err.Error())
		return
	}
	fmt.Println("Initialised")
	td.jsonFile = writeFile
	readFile, _:= os.OpenFile(td.jsonFileName, os.O_RDONLY, 0666)
	td.jsonFileScanner = bufio.NewScanner(readFile)
}

func (td *ToDoListJson) Add(item ToDoListItem) {
	bytes, _ := json.Marshal(item)
	_, err := td.jsonFile.WriteString(fmt.Sprintf("%s\n", string(bytes)))
	if(err != nil){
		fmt.Println("Error writing file", err.Error())
	}
}

func (td *ToDoListJson) Remove(item ToDoListItem) {
	// This function is best effort. First tries to remove by a matching hash, then by a matching name
	// failing that, Pop is called to remove the most recent item
	scanner := td.jsonFileScanner
	tdItem := &ToDoListItem{}
	for scanner.Scan(){
		fmt.Println(scanner.Text())
		json.Unmarshal(scanner.Bytes(), &tdItem)
	}
	fmt.Println(tdItem.Do)
}

func (td *ToDoListJson) Pop() {

}

func (td *ToDoListJson) Complete(){

}