package mytodolist

import "github.com/ekreke/myTodolist/internal/pkg/model"

type MyprojectsResponse struct {
	Projects []model.Projects `json:"projects"`
}

type ProjectQuitRequest struct {
	ProjectId int64 `form:"project_id"`
}
type JoinProjectRequest struct {
	ProjectId int64  `form:"project_id"`
	Password  string `form:"password"`
}
