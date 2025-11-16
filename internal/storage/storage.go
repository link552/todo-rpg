package storage

import (
	"database/sql"
	"os"
	_ "modernc.org/sqlite"
)

var db *sql.DB

func Init() {
	sqliteDbPath := os.Getenv("SQLITE_DB_PATH")
	if sqliteDbPath == "" {
		sqliteDbPath = "./todorpg.db"
	}

	var err error
	db, err = sql.Open("sqlite", sqliteDbPath)
	if err != nil {
		// TODO: Handle error.
		println(err.Error)
		os.Exit(1)
	}

	initUsersStorage();
	initTasksStorage();
}

func Deinit() {
	db.Close()
}

