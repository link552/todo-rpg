package storage

import (
	"todorpg/internal/core"
	"time"
	"os"
	_ "modernc.org/sqlite"
)

func initTasksStorage() {
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

func LoadCurrentTasks() []core.Task {
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

	var currentTasks []core.Task

	for rows.Next() {
		var t core.Task
		rows.Scan(&t.Id, &t.Title, &t.Short, &t.Long, &t.Energy);
		currentTasks = append(currentTasks, t)
	}

	return currentTasks
}

func LoadCompletedTasks() []core.Task {
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

	var currentTasks []core.Task

	for rows.Next() {
		var completedOnStr string
		var t core.Task
		rows.Scan(&t.Id, &t.Title, &t.Short, &t.Long, &t.Energy, &completedOnStr);
		t.CompletedOn, err = time.Parse(time.DateTime, completedOnStr)
		if err != nil {
			// TODO: Handle error.
		}
		currentTasks = append(currentTasks, t)
	}

	return currentTasks
}

func LoadTask(id int) core.Task {
	sql := `
	SELECT Id, Title, Short, Long, Energy, CompletedOn
	FROM Tasks
	WHERE Id = ?
	`

	rows, err := db.Query(sql, id)
	if err != nil {
		// TODO: Handle error.
		println(err.Error())
		os.Exit(1)
	}
	defer rows.Close()

	var t core.Task

	if rows.Next() {
		var completedOnStr string
		rows.Scan(&t.Id, &t.Title, &t.Short, &t.Long, &t.Energy, &completedOnStr);
		t.CompletedOn, err = time.Parse(time.DateTime, completedOnStr)
		if err != nil {
			// TODO: Handle error.
		}
	}

	return t
}

func CreateTask(title string, short int, long int, energy int) int {
	sql := `
	INSERT INTO Tasks (Title, Short, Long, Energy)
	VALUES (?, ?, ?, ?)
	`

	res, err := db.Exec(sql, title, short, long, energy)
	if err != nil {
		// TODO: Handle error.
		println(err.Error())
		os.Exit(1)
	}

	id, err := res.LastInsertId()
	if err != nil {
		// TODO: Handle error.
		println(err.Error())
		os.Exit(1)
	}

	 return int(id)
}

func SaveTask(t core.Task) {
	sql := `
	UPDATE Tasks
	SET Title = ?, Short = ?, Long = ?, Energy = ?, CompletedOn = ?
	WHERE Id = ?
	`

	completedOnStr := time.Now().Format(time.DateTime)
	_, err := db.Exec(sql, t.Title, t.Short, t.Long, t.Energy, completedOnStr, t.Id)
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
