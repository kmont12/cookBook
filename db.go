package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func InsertDB(insert string) {
	//log.Println("insert called")
	db, err := sql.Open("mysql", "root:password@/cookbook")
	if err != nil {
		log.Println(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
		return
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
		log.Println(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
		return nil
	}

	rows, err := db.Query(sel)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	return rows
}
