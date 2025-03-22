package todo

import (
	"database/sql"
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

 func (td *ToDoListSqlite) ExecLogError(sql string) sql.Result{
	db := td.db
	res, err := db.Exec(sql)
	if err != nil {
		log.Fatal(err)
	}
	return res
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
	// This function is best effort. We first try to remove entries with item.idId.
	deleteWithId := (fmt.Sprintf(`
			DELETE FROM %s
			WHERE Id='%d'
		`, td.toDoTableName, item.Id))

	fmt.Println(deleteWithId)
	res := td.ExecLogError(deleteWithId)
	if deleted, _ := res.RowsAffected(); deleted > 0{
		fmt.Printf("Removed %d records.", deleted)
		return
	}

	// Failing that, we try to remove any items with a matching Do
	deleteLikeDo := (fmt.Sprintf(`
			DELETE FROM %s
			WHERE Do LIKE '%%%s'
		`, td.toDoTableName, item.Do))

	fmt.Println(deleteWithId)
	res = td.ExecLogError(deleteLikeDo)
	if deleted, _ := res.RowsAffected(); deleted > 0{
		fmt.Printf("Removed %d records.", deleted)
		return
	}

}

func (td *ToDoListSqlite) Pop() {
	selectMaxId := (fmt.Sprintf(`
		SELECT MAX(id) as id FROM %s
	`, td.toDoTableName))
	var returnResult []ToDoListItem
	err := td.db.Select(&returnResult, selectMaxId)
	if err != nil{
		log.Fatal(err.Error())
	}

	if numRecords := len(returnResult); numRecords == 1 {
		maxIdRecord := returnResult[0]
		deleteMaxId := (fmt.Sprintf(`
			DELETE FROM %s WHERE id=%d 
			`, td.toDoTableName, maxIdRecord.Id))
		res := td.ExecLogError(deleteMaxId)
		if deleted, _ := res.RowsAffected(); deleted > 0{
			fmt.Printf("Removed %d records.", deleted)
			return
		}

	}else if numRecords > 1{
		log.Fatalf("Failed to determine most recent to do list item. Got %d records.", numRecords)
	}

}

func (td *ToDoListSqlite) Complete(item ToDoListItem) {
	td.Remove(item)
}

// Function must be called to close db conection
func (td *ToDoListSqlite) Close() {
	td.db.Close()
}
