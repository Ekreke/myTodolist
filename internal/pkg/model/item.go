package model

import "time"

// Items undefined
type Items struct {
	ID int64 `json:"id" gorm:"id"`
	// editable no influence
	ItemName string `json:"item_name" gorm:"item_name"`
	// editable no influence
	Description string `json:"description" gorm:"description"`
	ProjectId   int64  `json:"project_id" gorm:"project_id"`
	// editeable no influence
	Deadline time.Time `json:"deadline" gorm:"deadline"`
	// editable no influence
	Important int8 `json:"important" gorm:"important"`
	// editable no influence
	Done int8 `json:"done" gorm:"done"`
	// editable
	Myday       int8      `json:"myDay" gorm:"myDay"`
	CreatedTime time.Time `json:"created_time" gorm:"created_time"`
	Node        int8      `json:"node" gorm:"node"`
	Checkpoint  int8      `json:"checkPoint" gorm:"checkPoint"`
	// editable
	CollectionId int64 `json:"collection_id" gorm:"collection_id"`
}

// TableName 表名称
func (*Items) TableName() string {
	return "items"
}
