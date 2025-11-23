package storage

import (
	"todorpg/internal/core"
	"os"
	"time"
	"fmt"
	"strings"
	_ "modernc.org/sqlite"
)

func initEventsStorage() {
	sql := `
	CREATE TABLE IF NOT EXISTS Events (
		Id INTEGER PRIMARY KEY AUTOINCREMENT,
		UserId INTEGER NOT NULL,
		Type TEXT NOT NULL,
		ProcessedOn TEXT DEFAULT NULL,
		FOREIGN KEY (UserId) REFERENCES Users(Id)
	);
	`

	if _, err := db.Exec(sql); err != nil {
		// TODO: Handle error.
		println(err.Error())
		os.Exit(1)
	}
}

func LoadPendingUserEvents(userId int) []core.Event {
	sql := `
	SELECT Id, UserId, Type, ProcessedOn
	FROM Events
	WHERE UserId = ? AND ProcessedOn IS NULL
	`

	rows, err := db.Query(sql)
	if err != nil {
		// TODO: Handle error.
		println(err.Error())
		os.Exit(1)
	}
	defer rows.Close()

	var events []core.Event

	for rows.Next() {
		var processedOnStr string
		var e core.Event
		rows.Scan(&e.Id, &e.UserId, &e.Type, &processedOnStr);
		e.ProcessedOn, err = time.Parse(time.DateTime, processedOnStr)
		if err != nil {
			// TODO: Handle error.
		}
		events = append(events, e)
	}

	return events
}

func CreateEvent(userId int, _type string) int {
	sql := `
	INSERT INTO Events (UserId, Type)
	VALUES (?, ?)
	`

	res, err := db.Exec(sql, userId, _type)
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

func SaveEvent(e core.Event) {
	sql := `
	UPDATE Events
	SET UserId = ?, Type = ?, ProcessedOn = ?
	WHERE Id = ?
	`

	processedOnStr := e.ProcessedOn.Format(time.DateTime)
	_, err := db.Exec(sql, e.UserId, e.Type, processedOnStr, e.Id)
	if err != nil {
		// TODO: Handle error.
		println(err.Error())
		os.Exit(1)
	}
}

func CloseEvents(ids []int) {
	valueStrings := make([]string, 0, len(ids))
	valueArgs := make([]any, 0, len(ids) * 3)

	processedOnStr := time.Now().Format(time.DateTime)
	valueArgs = append(valueArgs, processedOnStr)

	for _, lp := range ids {
		valueStrings = append(valueStrings, "?")
		valueArgs = append(valueArgs, lp)
	}

	sql := "UPDATE Events SET ProcessedOn = ? WHERE Id IN (%s)"

	stmt := fmt.Sprintf(sql, strings.Join(valueStrings, ","))
	_, err := db.Exec(stmt, valueArgs...)
	if err != nil {
		// TODO: Handle error.
		println(err.Error())
		os.Exit(1)
	}
}
