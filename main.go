package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func main() {

	db, err := sql.Open("sqlite3", "./mck.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("select count(*) from mck")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	names, err := rows.Columns()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("NUM COLS", len(names))
	rawBuffer := make([]sql.RawBytes, len(names))
	scanCallArgs := make([]interface{}, len(names))
	for i := range rawBuffer {
		scanCallArgs[i] = rawBuffer[i]
	}
	for rows.Next() {
		err = rows.Scan(scanCallArgs...)
		if err != nil {
			log.Fatal(err)
		}
		for i := range scanCallArgs {
			fmt.Println(names[i], rawBuffer[i])
		}
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

}
