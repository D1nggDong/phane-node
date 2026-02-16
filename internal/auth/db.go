package auth

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() error {
	var err error
	DB, err = sql.Open("sqlite3", "./phane.db")
	if err != nil {
		return err
	}

	statement, _ := DB.Prepare("CREATE TABLE IF NOT EXISTS vault (id INTEGER PRIMARY KEY, key_name TEXT, encrypted_value TEXT)")
	statement.Exec()
	return nil
}
