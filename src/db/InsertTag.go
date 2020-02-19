package db

import (
	"database/sql"

	"main.main/src/utils"
)

var insertTagStmt *sql.Stmt = nil
var insertTagQueryString = `
INSERT INTO context(context_id, name)
VALUES (?,?);
`

//InsertTag insert tags into database
func InsertTag(tags []string) error {
	if insertTagStmt == nil {
		database, err := GetDBIns()
		if err != nil {
			return err
		}
		insertTagStmt, err = database.Prepare(insertTagQueryString)
		if err != nil {
			return err
		}
	}
	for i, str := range tags {
		_, err := insertTagStmt.Exec(utils.CreateID(str), str)
		if err != nil {
			println("err in index:", i)
			return err
		}
	}
	return nil
}
