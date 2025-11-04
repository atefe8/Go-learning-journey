package entity

import "time"

type Game struct {
	ID          uint
	CategoryID  uint
	QuestionIDs []uint
	PlayersIDs  []uint
	StartTime   time.Time
}
