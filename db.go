package main

import (
	"database/sql"
)

func InsertDB(insert string) {
	//log.Println("insert called")
	db, err := sql.Open("mysql", "root:password@/cookbook")
	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}

	_, err2 := db.Query(insert)
	if err2 != nil {
		panic(err2.Error())
	}
	defer db.Close()
}

func SearchDB(sel string) *sql.Rows {
	//log.Println("select called")
	db, err := sql.Open("mysql", "root:password@/cookbook")
	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}

	rows, err := db.Query(sel)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	return rows
}
