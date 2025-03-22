package todo

import (
	"fmt"
	"log"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

func test(){
}
type ToDoListSqlite struct {
	dbFileName    string
	db            *sqlx.DB
	toDoTableName string
}

func NewToDoListSqlite() (td *ToDoListSqlite) {
	td = &ToDoListSqlite{
		dbFileName:    ".todo.db",
		toDoTableName: "todo",
	}
	td.Init()
	return td
}

func (td *ToDoListSqlite) Init() {
	db, err := sqlx.Open("sqlite3", td.dbFileName)
	if err != nil {
		log.Fatal(err)
	}
	td.db = db

	td.ExecLogError((fmt.Sprintf(`
			CREATE TABLE IF NOT EXISTS %s
			(id INTEGER PRIMARY KEY, do TEXT, createdAt TIMESTAMP, byDays int , byHours int )
		`, td.toDoTableName)))
}

 func (td *ToDoListSqlite) ExecLogError(sql string){
	db := td.db
	_, err := db.Exec(sql)
	if err != nil {
		log.Fatal(err)
	}
}


func (td *ToDoListSqlite) Add(item ToDoListItem) {
	sqlInsert := (fmt.Sprintf(`
			INSERT INTO %s
			(do , createdAt, byDays, byHours)
			VALUES
			('%s',  CURRENT_TIMESTAMP, %d, %d)
		`, td.toDoTableName, item.Do, item.ByDays, item.ByHours))
	td.ExecLogError(sqlInsert)
}

func (td *ToDoListSqlite) Remove(item ToDoListItem) {
	// This function is best effort. We first try to remove entries with item.Id.
	sqlSelectById:= (fmt.Sprintf(`
			SELECT Id from %s
			WHERE Id='%d'
		`, td.toDoTableName, item.Id))

	fmt.Println(sqlSelectById)
	var selectReturn []ToDoListItem
	err := td.db.Select(&selectReturn,  sqlSelectById)
	if err != nil{
		log.Fatal(err.Error())
	}

	fmt.Println(selectReturn)

	// Failing that, we try to remove any items with a matching Do

	// Failing that we remove the most recent item

}

func (td *ToDoListSqlite) Pop() {

}

func (td *ToDoListSqlite) Complete() {

}

// Function must be called to close db conection
func (td *ToDoListSqlite) Close() {
	td.db.Close()

}
