package model

import "time"

// Users undefined
type Users struct {
	ID        int64     `json:"id" gorm:"id"`
	Username  string    `json:"username" gorm:"username"`
	Password  string    `json:"password" gorm:"password"`
	Link      string    `json:"link" gorm:"link"`
	Bio       string    `json:"bio" gorm:"bio"`
	Avatar    string    `json:"avatar" gorm:"avatar"`
	CreatedAt time.Time `json:"created_at" gorm:"created_at"`
	UpdatedAt time.Time `json:"updated_at" gorm:"updated_at"`
	DeletedAt time.Time `json:"deleted_at" gorm:"deleted_at"`
	Root      int8      `json:"root" gorm:"root"`
}

// TableName 表名称
func (*Users) TableName() string {
	return "users"
}
