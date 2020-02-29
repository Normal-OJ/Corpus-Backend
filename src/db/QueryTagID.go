package db

import "database/sql"

var queryTagFromContext = `select context_id,name from context`
var queryTagFromContextStmt *sql.Stmt = nil

//QueryTagID query all the exsisted tag and tagID inside database
func QueryTagID() (map[string]int64, error) {
	if queryTagFromContextStmt == nil {
		database, err := GetDBIns()
		if err != nil {
			println("error  when QueryTagID obtain db instance:", err.Error())
			return nil, err
		}
		queryTagFromContextStmt, err = database.Prepare(queryTagFromContext)
		if err != nil {
			println("error when QueryTagID preparing  statement:", err.Error())
			return nil, err
		}
	}
	rows, err := queryTagFromContextStmt.Query()
	if err == sql.ErrNoRows {
		return nil, err
	} else if err != nil {
		return nil, err
	}

	ret := make(map[string]int64)

	for rows.Next() {
		var tag string = ""
		var id int64 = 0
		rows.Scan(&id, &tag)
		ret[tag] = id

		if rows.Err() != nil {
			return nil, rows.Err()
		}
	}
	return ret, nil
}
