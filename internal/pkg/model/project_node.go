package model

// ProjectsNodes undefined
type ProjectsNodes struct {
	ProjectId int64 `json:"project_id" gorm:"project_id"`
	ItemId    int64 `json:"item_id" gorm:"item_id"`
}

// TableName 表名称
func (*ProjectsNodes) TableName() string {
	return "projects_nodes"
}
