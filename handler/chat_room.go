package handler

import (
	"time"
)

type chatRoom struct {
	Id          int
	UUID        []byte
	Name        string
	LastMessage string
	TimeStamp   time.Time
}
