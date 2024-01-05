package model

import "time"

type Item struct {
	ID            int
	ItemName      string
	Description   string
	FromProjectID int
	Deadline      time.Time
	IsImportant   bool
}
