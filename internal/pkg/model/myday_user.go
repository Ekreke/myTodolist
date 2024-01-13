package model

// MydayUsers undefined
type MydayUsers struct {
	ID     int64 `json:"id" gorm:"id"`
	UserId int64 `json:"user_id" gorm:"user_id"`
	ItemId int64 `json:"item_id" gorm:"item_id"`
}

// TableName 表名称
func (*MydayUsers) TableName() string {
	return "myday_users"
}
