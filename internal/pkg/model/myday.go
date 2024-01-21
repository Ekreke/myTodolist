package model

type Myday struct {
	User_id int64 `json:"user_id" gorm:"user_id"`
	Item_id int64 `json:"item_id" gorm:"item_id"`
}
