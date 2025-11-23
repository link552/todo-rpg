package storage

import (
	"todorpg/internal/core"
	"os"
	"fmt"
	"strings"
	_ "modernc.org/sqlite"
)

func initLevelProgressStorage() {
	sql := `
	CREATE TABLE IF NOT EXISTS LevelProgress (
		EventId INTEGER NOT NULL,
		FromLevel INTEGER NOT NULL,
		ToLevel INTEGER NOT NULL,
		ToPercent FLOAT NOT NULL,
		Sequence INTEGER NOT NULL,
		FOREIGN KEY (EventId) REFERENCES Events(Id)
	);
	`

	if _, err := db.Exec(sql); err != nil {
		// TODO: Handle error.
		println(err.Error())
		os.Exit(1)
	}
}

func LoadUserLevelProgress(userId int) []core.LevelProgress {
	sql := `
	SELECT EventId, FromLevel, ToLevel, ToPercent, Sequence
	FROM Events
	LEFT JOIN LevelProgress ON LevelProgress.EventId = Events.Id
	WHERE UserId = ?  AND Type = 'level-progress' AND ProcessedOn IS NULL
	ORDER BY Sequence
	`

	rows, err := db.Query(sql, userId)
	if err != nil {
		// TODO: Handle error.
		println(err.Error())
		os.Exit(1)
	}
	defer rows.Close()

	var lps []core.LevelProgress

	for rows.Next() {
		var lp core.LevelProgress
		rows.Scan(&lp.EventId, &lp.FromLevel, &lp.ToLevel, &lp.ToPercent, &lp.Sequence);
		lps = append(lps, lp)
	}

	return lps
}

func InsertLevelProgress(lps []core.LevelProgress) {
	valueStrings := make([]string, 0, len(lps))
	valueArgs := make([]any, 0, len(lps) * 3)
	for _, lp := range lps {
		valueStrings = append(valueStrings, "(?, ?, ?, ?, ?)")
		valueArgs = append(valueArgs, lp.EventId)
		valueArgs = append(valueArgs, lp.FromLevel)
		valueArgs = append(valueArgs, lp.ToLevel)
		valueArgs = append(valueArgs, lp.ToPercent)
		valueArgs = append(valueArgs, lp.Sequence)
	}

	sql := "INSERT INTO LevelProgress (EventId, FromLevel, ToLevel, ToPercent, Sequence) VALUES %s"

	stmt := fmt.Sprintf(sql, strings.Join(valueStrings, ","))
	_, err := db.Exec(stmt, valueArgs...)
	if err != nil {
		// TODO: Handle error.
		println(err.Error())
		os.Exit(1)
	}
}
