package model

// CollectionsItems undefined
type CollectionsItems struct {
	CollectionId int64 `json:"collection_id" gorm:"collection_id"`
	ItemsId      int64 `json:"items_id" gorm:"items_id"`
}

// TableName 表名称
func (*CollectionsItems) TableName() string {
	return "collections_items"
}
