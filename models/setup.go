package models

import (
	"database/sql"
	"fmt"
)

var Db *sql.DB

func Setup()  {
	Db, _ = sql.Open("mysql", fmt.Sprintf("root:%v@tcp(localhost:3306)/go_example", "#Neo@1221"))
}
