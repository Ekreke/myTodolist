package model

// ProjectsUsers undefined
type ProjectsUsers struct {
	ProjectsId int64 `json:"projects_id" gorm:"projects_id"`
	UsersId    int64 `json:"users_id" gorm:"users_id"`
}

// TableName 表名称
func (*ProjectsUsers) TableName() string {
	return "projects_users"
}
