package main

import (
	"database/sql"
	"fmt"

	_ "modernc.org/sqlite"
)

func main() {
	db, err := sql.Open("sqlite", "/Users/gobpo2002/Tools/sqlite-tools-osx-x64-3500400/school.sql")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()
}
