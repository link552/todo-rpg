package core

import "time"

type User struct {
	Id int
	Level int
	TotalExp int
}

type Task struct {
	Id int
	Title string
	Short int
	Long int
	Energy int
	Priority int
	CompletedOn time.Time
}

type Event struct {
	Id int
	UserId int
	Type string
	ProcessedOn time.Time
}

type LevelProgress struct {
	EventId int
	FromLevel int
	ToLevel int
	ToPercent float32
	Sequence int
}
