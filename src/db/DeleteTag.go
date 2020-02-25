package db

import "database/sql"

var deleteTagFromContext = `delete from context where name == ?`
var deleteTagFromMap = `delete from map where context_id not in (select context_id from context)`

var deleteTagFromContextStmt *sql.Stmt = nil
var deleteTagFromMapStmt *sql.Stmt = nil

//DeleteTag delete specify tag
func DeleteTag(tag string) {
	database, err := GetDBIns()
	if err != nil {
		println("error in delete Tag:", err.Error())
	}

	if deleteTagFromContextStmt == nil {
		deleteTagFromContextStmt, err = database.Prepare(deleteTagFromContext)
		if err != nil {
			println("error in deleteTag  init context:", err.Error())
		}
	}

	if deleteTagFromMapStmt == nil {
		deleteTagFromMapStmt, err = database.Prepare(deleteTagFromMap)
		if err != nil {
			println("error in deleteTag  init map:", err.Error())
		}
	}

	deleteTagFromContextStmt.Exec(tag)
	deleteTagFromMapStmt.Exec()
}
