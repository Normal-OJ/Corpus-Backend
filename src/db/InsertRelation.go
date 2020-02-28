package db

import "database/sql"

var insertRelationToMap = `insert into map(cha_id , context_id) values(?,?)`
var insertRelationToMapStmt *sql.Stmt = nil

//InsertRelation insert the relation between cha file and tags into database
func InsertRelation(chaID int64, tagID []int64) error {
	if insertRelationToMapStmt == nil {
		database, err := GetDBIns()
		if err != nil {
			println("error in InsertRelation obtain db instance:", err.Error())
			return err
		}
		insertRelationToMapStmt, err = database.Prepare(insertRelationToMap)
		if err != nil {
			println("error in insertRelation prepare statement:", err.Error())
			return err
		}
	}
	for _, tag := range tagID {
		_, err := insertRelationToMapStmt.Exec(chaID, tag)
		if err != nil {
			return err
		}
	}
	return nil
}
