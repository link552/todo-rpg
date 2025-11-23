package storage

import (
	"os"
	"todorpg/internal/core"
	_ "modernc.org/sqlite"
)

func initUsersStorage() {
	sql := `
	CREATE TABLE IF NOT EXISTS Users (
		Id INTEGER PRIMARY KEY AUTOINCREMENT,
		Level INT NOT NULL DEFAULT 1,
		TotalExp INT NOT NULL DEFAULT 0
	);
	`

	if _, err := db.Exec(sql); err != nil {
		// TODO: Handle error.
		println(err.Error())
		os.Exit(1)
	}

	// Insert default user ONLY if the Users table is empty.
	sql = `
	INSERT INTO Users (Level, TotalExp)
	SELECT 1, 0
	WHERE NOT EXISTS (SELECT 1 FROM users);
	`

	_, err := db.Exec(sql)
	if err != nil {
		// TODO: Handle error.
		println(err.Error())
		os.Exit(1)
	}
}

func LoadUser(id int) core.User {
	sql := `
	SELECT Id, Level, TotalExp
	FROM Users
	WHERE Id = ?;
	`

	rows, err := db.Query(sql, id)
	if err != nil {
		// TODO: Handle error.
		println(err.Error())
		os.Exit(1)
	}
	defer rows.Close()

	var u core.User

	if rows.Next() {
		rows.Scan(&u.Id, &u.Level, &u.TotalExp);
	}

	return u
}

func SaveUser(u core.User) {
	sql := `
	UPDATE Users
	SET Level = ?, TotalExp = ?
	WHERE Id = ?
	`

	_, err := db.Exec(sql, u.Level, u.TotalExp, u.Id)
	if err != nil {
		// TODO: Handle error.
		println(err.Error())
		os.Exit(1)
	}
}
