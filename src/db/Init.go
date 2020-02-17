package db

import (
	"database/sql"
	"errors"
)

var dbSchema = `CREATE TABLE IF NOT EXISTS cha (
	cha_id integer primary key, 
	path TEXT, 
	age integer,
	sex integer
  );
  
  CREATE TABLE IF NOT EXISTS context (
	context_id integer primary key, 
	name TEXT
  );
  
  CREATE TABLE IF NOT EXISTS map (
	cha_id integer,
	context_id integer,
	FOREIGN KEY(cha_id) REFERENCES cha(id),
	FOREIGN KEY(context_id) REFERENCES context(id)
  );`
var addSample = `
INSERT INTO cha(cha_id, path, age, sex)
VALUES (1, 'a.cha', 12, 1);

INSERT INTO cha(cha_id, path, age, sex)
VALUES (2, 'b/b.cha', 24, 2);

INSERT INTO context(context_id, name)
VALUES (1, 'toy');

INSERT INTO context(context_id, name)
VALUES (2, 'book');

INSERT INTO map(cha_id, context_id)
VALUES (1, 1);

INSERT INTO map(cha_id, context_id)
VALUES (1, 2);

INSERT INTO map(cha_id, context_id)
VALUES (2, 2);
`
var dbIns *sql.DB = nil

//Init is used for initialization of whole db module
func Init(database *sql.DB) error {
	if dbIns != nil {
		return nil
	}
	dbIns = database
	_, err := database.Exec(dbSchema)
	if err != nil {
		println("error when creating schema:", err.Error())
		return err
	}
	return nil
}

//ErrInstanceNotExisted just like what it said , need to run Init before use
var ErrInstanceNotExisted = errors.New("DB instance not created yet")

//GetDBIns accquire an instance of db from module
func GetDBIns() (*sql.DB, error) {
	if dbIns == nil {
		return nil, ErrInstanceNotExisted
	}
	return dbIns, nil
}

//AddTestSample creates same test sample , will be removed in the future
func AddTestSample() {
	database, err := GetDBIns()
	if err != nil {
		println("Error in init AddTestSample:", err.Error())
		return
	}
	_, err = database.Exec(addSample)
	if err != nil {
		println("error in AddTestSample adding test sample:", err.Error())
	}
}
