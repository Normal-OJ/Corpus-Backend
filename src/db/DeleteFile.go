package db

import "database/sql"

var deleteFileFromCha = `delete from cha where path == ?`
var deleteFileFromMap = `delete from map where cha_id not in (select cha_id from cha)`

var deleteFileFromChaStmt *sql.Stmt = nil
var deleteFileFromMapStmt *sql.Stmt = nil

//DeleteFile delete specify file  on given route
func DeleteFile(src string) {
	database, err := GetDBIns()
	if err != nil {
		println("error in delete Tag:", err.Error())
	}

	if deleteFileFromChaStmt == nil {
		deleteFileFromChaStmt, err = database.Prepare(deleteFileFromCha)
		if err != nil {
			println("error in deleteFile init cha:", err.Error())
		}
	}

	if deleteFileFromMapStmt == nil {
		deleteFileFromMapStmt, err = database.Prepare(deleteFileFromMap)
		if err != nil {
			println("error in deleteFile init map:", err.Error())
		}
	}

	deleteFileFromChaStmt.Exec(src)
	deleteFileFromMapStmt.Exec()
}
