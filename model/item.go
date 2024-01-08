package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Item struct {
	ID            int
	ItemName      string
	Description   string
	FromProjectID int
	Deadline      time.Time
	IsImportant   bool
}

func GetItemByItemId(db *gorm.DB, id int) Item {
	var item Item
	db.Where("id =?", id).Find(item)
	return item
}
