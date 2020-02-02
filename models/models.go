package models

import (
	"database/sql"
	"fmt"
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

func (b *Blog) GetAll() []Blog {
	query := "Select * From Blog"
	rows := ReadQuery(query)

	result := []Blog{}

	data := Blog{}
	for rows.Next() {
		err := rows.Scan(&b.Title, &b.Description)
		if err != nil {
			log.Fatal(err)
		}
		data.Title = b.Title
		data.Description = b.Description

		result = append(result, data)
	}
	return result
}
