package model

type Collections struct {
	ID   int64  `json:"id" gorm:"id"`
	Name string `json:"name" gorm:"name"`
	Icon int64  `json:"icon" gorm:"icon"`
}

// TableName 表名称
func (*Collections) TableName() string {
	return "collections"
}
