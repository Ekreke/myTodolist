package model

type CollectionsUsers struct {
	CollectionId int64 `json:"collection_id" gorm:"collection_id"`
	UserId       int64 `json:"user_id" gorm:"user_id"`
}

// TableName 表名称
func (*CollectionsUsers) TableName() string {
	return "collections_users"
}