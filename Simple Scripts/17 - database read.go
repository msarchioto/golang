package main

import "fmt"
import (
	_ "github.com/a-palchikov/sqlago"
	"database/sql"
	"log"
)

func main() {

	db, err := sql.Open("sqlany", "uid=dba;pwd=sql;Host=10.28.1.178:2638")
	if err != nil {
		log.Fatalf("Unable to connect to db: %s", err)
	}

	fmt.Println("OK")

	var name string
	err = db.QueryRow("select COLUMN_NAME from meta.T_BL_PREDEFINED_COLS").Scan(&name)
	if err != nil {
		log.Fatalf("Select failed: %s", err)
	}

	fmt.Println(name)
}
