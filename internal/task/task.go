package task

import "time"

type Task struct {
	Id int
	Title string
	Short int
	Long int
	Energy int
	Priority int
	CompletedOn time.Time
}
