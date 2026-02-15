package core

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func InitDB() {
	db, err := sql.Open("sqlite3", "./phane_nodes.db")
	if err != nil { log.Fatal(err) }
	defer db.Close()

	// Create table for 24/7 tasks
	sqlStmt := `
	CREATE TABLE IF NOT EXISTS tasks (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT,
		trigger_type TEXT,
		action_skill TEXT,
		last_run DATETIME,
		status TEXT
	);`
	_, err = db.Exec(sqlStmt)
	if err != nil { log.Printf("DB Error: %v", err) }
	log.Println("🗄️ Database Initialized: Tasks Table Ready.")
}
