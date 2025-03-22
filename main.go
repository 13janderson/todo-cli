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

	list.Add(
		todo.ToDoListItem{
			Do: "this is not a drill",
			ByDays: 1,
		},
	)
	list.Remove(todo.ToDoListItem{
		Do: "test",
	})
}
		