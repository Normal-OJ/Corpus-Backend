package db

import (
	"database/sql"

	"main.main/src/utils"
)

var stm *sql.Stmt = nil
var queryString = `
INSERT INTO context(context_id, name)
VALUES (?,?);
`

//InsertTag insert tags into database
func InsertTag(tags []string) error {
	if stm == nil {
		database, err := GetDBIns()
		if err != nil {
			return err
		}
		stm, err = database.Prepare(queryString)
		if err != nil {
			return err
		}
	}
	for i, str := range tags {
		_, err := stm.Exec(utils.CreateID(str), str)
		if err != nil {
			println("err in index:", i)
			return err
		}
	}
	return nil
}
