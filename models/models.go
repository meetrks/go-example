package models

import (
	"database/sql"
	"fmt"
	"golang.org/x/net/dict"
	"log"
)

type Blog struct {
	Title       string
	Description string
}

func (b *Blog) Add() *sql.Stmt {
	query := fmt.Sprintf("Insert into Blog VALUES ('%s', '%s')", b.Title, b.Description)
	res := WriteQuery(query)
	return res
}

func (b *Blog) GetAll() map[string]interface{} {
	query := "Select * From Blog"
	rows := ReadQuery(query)

	result := make(map[dict.Dict]interface{})

	data := make(map[string]interface{})
	for rows.Next() {
		err := rows.Scan(&b.Title, &b.Description)
		if err != nil {
			log.Fatal(err)
		}
		data["title"] = b.Title
		data["description"] = b.Description
	}
	return data
}
