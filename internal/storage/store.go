package storage

import (
	"database/sql"
	"todorpg/internal/task"
	"time"
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

	sql := `
	CREATE TABLE IF NOT EXISTS Tasks (
		Id INTEGER PRIMARY KEY AUTOINCREMENT,
		Title TEXT NOT NULL,
		Short INT NOT NULL,
		Long INT NOT NULL,
		Energy INT NOT NULL,
		CompletedOn TEXT DEFAULT NULL
	);
	`
	if _, err := db.Exec(sql); err != nil {
		// TODO: Handle error.
		println(err.Error())
		os.Exit(1)
	}
}

func Deinit() {
	db.Close()
}

func SelectCurrentTasks() []task.Task {
	sql := `
	SELECT Id, Title, Short, Long, Energy
	FROM Tasks
	WHERE CompletedOn IS NULL;
	`

	rows, err := db.Query(sql)
	if err != nil {
		// TODO: Handle error.
		println(err.Error())
		os.Exit(1)
	}
	defer rows.Close()

	var currentTasks []task.Task

	for rows.Next() {
		var t task.Task
		rows.Scan(&t.Id, &t.Title, &t.Short, &t.Long, &t.Energy);
		currentTasks = append(currentTasks, t)
	}

	return currentTasks
}

func SelectCompletedTasks() []task.Task {
	sql := `
	SELECT Id, Title, Short, Long, Energy, CompletedOn
	FROM Tasks
	WHERE CompletedOn IS NOT NULL
	`

	rows, err := db.Query(sql)
	if err != nil {
		// TODO: Handle error.
		println(err.Error())
		os.Exit(1)
	}
	defer rows.Close()

	var currentTasks []task.Task

	for rows.Next() {
		var completedOnStr string
		var t task.Task
		rows.Scan(&t.Id, &t.Title, &t.Short, &t.Long, &t.Energy, &completedOnStr);
		t.CompletedOn, err = time.Parse(time.DateTime, completedOnStr)
		if err != nil {
			// TODO: Handle error.
		}
		currentTasks = append(currentTasks, t)
	}

	return currentTasks
}

func InsertTask(t task.Task) {
	sql := `
	INSERT INTO Tasks (Title, Short, Long, Energy)
	VALUES (?, ?, ?, ?)
	`

	_, err := db.Exec(sql, t.Title, t.Short, t.Long, t.Energy)
	if err != nil {
		// TODO: Handle error.
		println(err.Error())
		os.Exit(1)
	}
}

func UpdateTask(t task.Task) {
	sql := `
	UPDATE Tasks
	SET Title = ?, Short = ?, Long = ?, Energy = ?
	WHERE Id = ?
	`

	_, err := db.Exec(sql, t.Title, t.Short, t.Long, t.Energy, t.Id)
	if err != nil {
		// TODO: Handle error.
		println(err.Error())
		os.Exit(1)
	}
}

func CompleteTask(id int) {
	sql := `
	UPDATE Tasks
	SET CompletedOn = ?
	WHERE Id = ?
	`
	completedOnStr := time.Now().Format(time.DateTime)
	_, err := db.Exec(sql, completedOnStr, id)
	if err != nil {
		// TODO: Handle error.
		println(err.Error())
		os.Exit(1)
	}
}

func DeleteTask(id int) {
	sql := `
	DELETE FROM Tasks
	WHERE Id = ?
	`

	_, err := db.Exec(sql, id)
	if err != nil {
		// TODO: Handle error.
		println(err.Error())
		os.Exit(1)
	}
}
