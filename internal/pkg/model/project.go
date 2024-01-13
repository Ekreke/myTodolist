package model

import "time"

// Projects undefined
type Projects struct {
	ID          int64     `json:"id" gorm:"id"`
	Name        string    `json:"name" gorm:"name"`
	Description string    `json:"description" gorm:"description"`
	CreatedTime time.Time `json:"created_time" gorm:"created_time"`
	EndTime     time.Time `json:"end_time" gorm:"end_time"`
	Password    string    `json:"password" gorm:"password"`
	AdminId     int64     `json:"admin_id" gorm:"admin_id"`
}

// TableName 表名称
func (*Projects) TableName() string {
	return "projects"
}
