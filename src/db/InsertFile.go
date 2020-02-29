package db

import (
	"database/sql"
	"fmt"
)

var insertFileToCha = `insert into cha(cha_id, path,age,sex) values (?,?,?,?)`
var insertFileToChaStmt *sql.Stmt = nil

//InsertFile insert file information into database
func InsertFile(chaID int64, src string, age int, sex int) error {
	if insertFileToChaStmt == nil {
		database, err := GetDBIns()
		if err != nil {
			println("error when InsertFile obtain instance:", err.Error())
			return err
		}
		insertFileToChaStmt, err = database.Prepare(insertFileToCha)
		if err != nil {
			println("error when preparing statement")
			return err
		}
	}
	m, err := QueryChaID()
	if err != nil {
		println("error when query cha file id")
		return err
	}
	for k, v := range m {
		if v == chaID {
			fmt.Printf("collision with %s with same id %d\n", k, v)
		}
	}
	_, err = insertFileToChaStmt.Exec(chaID, src, age, sex)
	return err
}
