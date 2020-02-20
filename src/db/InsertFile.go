package db

import (
	"database/sql"
)

var insertFileToCha = `insert into cha(cha_id, path,age,sex) values (?,?,?,?)`
var insertFileToChaStmt *sql.Stmt = nil

//InsertFile insert file information into database
func InsertFile(chaID uint64, src string, age int, sex int) error {
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
	_, err := insertFileToChaStmt.Exec(chaID, src, age, sex)
	return err
}
