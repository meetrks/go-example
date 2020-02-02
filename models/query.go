package models

import "database/sql"

func WriteQuery(query string) *sql.Stmt {
	res, err := Db.Prepare(query)
	if err != nil {
		panic(err.Error())
	}
	defer res.Close()
	_, err = res.Exec()
	if err != nil {
		panic(err.Error())
	}
	return res
}

func ReadQuery(query string) *sql.Rows {
	res, err := Db.Prepare(query)
	if err != nil {
		panic(err.Error())
	}
	defer res.Close()
	rows, err := res.Query()
	if err != nil {
		panic(err.Error())
	}
	return rows
}
