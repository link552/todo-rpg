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
