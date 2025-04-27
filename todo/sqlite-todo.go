package todo

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

type ToDoListSqlite struct {
	dbFileName    string
	db            *sqlx.DB
	toDoTableName string
}

func DefaultToDoListSqliteInDirectory(directory string) (td *ToDoListSqlite) {
	td = &ToDoListSqlite{
		dbFileName:    filepath.Join(directory, ".todo.db"),
		toDoTableName: "todo",
	}
	td.openDbConnection()
	return td
}


func DefaultToDoListSqlite() (td *ToDoListSqlite) {
	td = &ToDoListSqlite{
		dbFileName:    ".todo.db",
		toDoTableName: "todo",
	}
	td.openDbConnection()
	
	return td
}


func (td *ToDoListSqlite) openDBFile() error{
	if _, err := os.Open(td.dbFileName); err != nil{
		return err
	}
	return nil
}

func (td *ToDoListSqlite) List() ([]ToDoListItem, error){
	var allItems []ToDoListItem
	if dbFileOpen := td.openDBFile(); dbFileOpen != nil{
		return allItems, errors.New("to list was not initalised.\nrun td init first")
	}

	// Select all entries in DB
	sqlSelectAll := fmt.Sprintf(`
		SELECT * FROM %s
		ORDER BY doBy DESC
	`, td.toDoTableName)

	td.db.Select(&allItems, sqlSelectAll)

	return allItems, nil

}


func (td *ToDoListSqlite) openDbConnection() {
	db, err := sqlx.Open("sqlite3", td.dbFileName)
	td.db = db
	if err != nil{
		log.Fatal(err.Error())
	}
}

func (td *ToDoListSqlite) Init() error{
	if dbFileOpen := td.openDBFile(); dbFileOpen == nil{
		return errors.New("to do list already exists in this directory.\n to remove it type: td rm")
	}

	td.ExecLogError((fmt.Sprintf(`
			CREATE TABLE IF NOT EXISTS %s
			(id INTEGER PRIMARY KEY, do TEXT, doBy TIMESTAMP, createdAt TIMESTAMP)
		`, td.toDoTableName)))
		
	return nil
}

 func (td *ToDoListSqlite) ExecLogError(sql string) sql.Result{
	db := td.db
	res, err := db.Exec(sql)
	if err != nil {
		log.Fatal(err)
	}
	return res
}


func (td *ToDoListSqlite) Add(item *ToDoListItem) error{

	sqlInsert := (fmt.Sprintf(`
			INSERT INTO %s
			(do , doBy, createdAt)
			VALUES
			(:do, :doBy, :createdAt)
		`, td.toDoTableName, ))
	res, err := td.db.NamedExec(sqlInsert, &item)
	if err != nil{
		return err
	}else{
		if rows, err := res.RowsAffected(); rows != 1{
			if err != nil{
				return err
			}
			return errors.New("multiple rows created for a single insertion")
		}else{
			insertedId, _ := res.LastInsertId()
			item.Id = (int) (insertedId)
		}
	}
	return nil
}

func (td *ToDoListSqlite) SelectWithId(id int) ([]ToDoListItem, error){
	var itemsWithId []ToDoListItem

	selectWithId := (fmt.Sprintf(`
		SELECT * FROM %s 
		WHERE id=%d
	`, td.toDoTableName, id))

	err := td.db.Select(&itemsWithId, selectWithId)
	return itemsWithId, err
}

func (td *ToDoListSqlite) Remove(item ToDoListItem) int{
	// This function is best effort. We first try to remove entries with item.idId.
	deleteWithId := (fmt.Sprintf(`
			DELETE FROM %s
			WHERE Id='%d'
	`, td.toDoTableName, item.Id))
	
	var deleted int64 
	res := td.ExecLogError(deleteWithId)
	deleted, _ = res.RowsAffected()
	if(deleted > 0){
		return (int) (deleted)
	}

	// Failing that, we try to remove any items with a matching Do
	deleteLikeDo := (fmt.Sprintf(`
			DELETE FROM %s
			WHERE Do LIKE '%%%s'
		`, td.toDoTableName, item.Do))

	res = td.ExecLogError(deleteLikeDo)
	deleted, _ = res.RowsAffected()

	return (int) (deleted)
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
			fmt.Printf("Removed %d records",  100)
			return
		}

	}else if numRecords > 1{
		log.Fatalf("failed to determine most recent to do list item. Got %d records", numRecords)
	}
}

func (td *ToDoListSqlite) Extend(item ToDoListItem) error {
	updateWithId := (fmt.Sprintf(`
		UPDATE %s SET doBy=:doBy
		WHERE id=:id
	`, td.toDoTableName))

	fmt.Println(item.String())
	_ , err := td.db.NamedExec(updateWithId, &item)
	return err
}

func (td *ToDoListSqlite) Complete(item ToDoListItem) {
	td.Remove(item)
}

// Function must be called to close db conection
func (td *ToDoListSqlite) Close() {
	td.db.Close()
}
