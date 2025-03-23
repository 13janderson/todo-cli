package todo

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"
	"time"
	"github.com/fatih/color"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

type ToDoListSqlite struct {
	dbFileName    string
	db            *sqlx.DB
	toDoTableName string
}

func DefaultToDoListSqlite() (td *ToDoListSqlite) {
	td = &ToDoListSqlite{
		dbFileName:    ".todo.db",
		toDoTableName: "todo",
	}
	td.openDbConnection()
	
	return td
}



func (td ToDoListItem) RemainingTime() time.Duration{
	return time.Until(td.DoBy)
}


func (td *ToDoListSqlite) openDBFile() error{
	if _, err := os.Open(td.dbFileName); err != nil{
		return err
	}
	return nil
}

func (td *ToDoListSqlite) List() error{
	if dbFileOpen := td.openDBFile(); dbFileOpen != nil{
		return errors.New("to list was not initalised.\n run td init first")
	}

	// Select all entries in DB
	var allItems []ToDoListItem
	sqlSelectAll := fmt.Sprintf(`
		SELECT * FROM %s
		ORDER BY doBy DESC
	`, td.toDoTableName)

	td.db.Select(&allItems, sqlSelectAll)

	indent := fmt.Sprintf("%*s", 4, "")
	for _, item := range allItems{
		remainingTime := item.RemainingTime()
		color.Set(color.Bold)
		if remainingTime <= time.Duration(0){
			color.Red("%s [%d] %s EXPIRED", indent, item.Id, item.Do)
		}else{
			color.Green("%s [%d] %s %s", indent, item.Id, item.Do, DurationHumanReadable(remainingTime))
		}
	}
	return nil
}

func DurationHumanReadable(d time.Duration) string{
	var parts []string
	
	day := time.Hour * 24
	days := int(d / day)
	afterDays := d - time.Duration(days)*day
	hours := int(afterDays / time.Hour)
	afterHours := afterDays - time.Duration(time.Hour)*time.Duration(hours)
	mins := int(afterHours / time.Minute)

	if days > 0 {
		parts = append(parts, fmt.Sprintf("%dd", days))
	}
	if hours > 0 {
		parts = append(parts, fmt.Sprintf("%dh", hours))
	}
	if mins > 0 {
		parts = append(parts, fmt.Sprintf("%dm", mins))
	}

	return fmt.Sprintf("%s", parts)

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
			(id INTEGER PRIMARY KEY, do TEXT, doBy TIMESTAMP)
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


func (td *ToDoListSqlite) Add(item ToDoListItem) error{

	sqlInsert := (fmt.Sprintf(`
			INSERT INTO %s
			(do , doBy)
			VALUES
			(:do, :doBy)
		`, td.toDoTableName, ))
	res, err := td.db.NamedExec(sqlInsert, &item)
	if err != nil{
		return err
	}else{
		if rows, err := res.RowsAffected(); rows != 1{
			if err != nil{
				return err
			}
			return errors.New("Multiple rows created for a single insertion.")
		}
	}
	return nil
}

func (td *ToDoListSqlite) Remove(item ToDoListItem) {
	// This function is best effort. We first try to remove entries with item.idId.
	deleteWithId := (fmt.Sprintf(`
			DELETE FROM %s
			WHERE Id='%d'
		`, td.toDoTableName, item.Id))

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
