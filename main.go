package main

import (
	"todo/todo"
)

func main(){
	list := todo.NewToDoListSqlite()
	list.Add(
		todo.ToDoListItem{
			Do: "test",
			ByDays: 1,
		},
	)
	list.Remove(todo.ToDoListItem{
		Id: 1,
	})
}
		