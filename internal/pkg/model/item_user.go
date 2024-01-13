package model

// ItemsUsers undefined
type ItemsUsers struct {
	ItemId int64 `json:"item_id" gorm:"item_id"`
	UserId int64 `json:"user_id" gorm:"user_id"`
}

// TableName 表名称
func (*ItemsUsers) TableName() string {
	return "items_users"
}
