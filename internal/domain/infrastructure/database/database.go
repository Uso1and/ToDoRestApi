package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func Init() error {

	connsql := `user=postgres password=root dbname=todo sslmode=disable `

	var err error

	DB, err = sql.Open("postgres", connsql)

	if err != nil {
		return err
	}

	err = DB.Ping()

	if err != nil {
		return err
	}
	fmt.Print("Пинг к бд есть")
	return nil
}
