package mytodolist

import (
	"time"

	"github.com/ekreke/myTodolist/internal/pkg/model"
)

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

type ProjectInfoRequest struct {
	ProjectId int64 `form:"project_id"`
}

type ProjectInfoResponse struct {
	Projects model.Projects `json:"projects"`
	Users    []model.Users  `json:"users"`
}

type ProjectCreateRequest struct {
	Name        string `form:"name"`
	Description string `form:"description"`
	// endtime input with unix time
	Endtime int64 `form:"endtime"`
}

type IcreatedResponse struct {
	Projects []model.Projects `json:"projects"`
}

type ProjectUpdateRequest struct {
	ID          int    `form:"id"`
	Name        string `form:"name"`
	Description string `form:"description"`
	// endtime input with unix time
	Endtime int64 `form:"endtime"`
}

type ProjectUpdateNodeRequest struct {
	ItemName     string    `json:"item_name" form:"item_name"`
	Description  string    `json:"description" form:"description"`
	ProjectId    int64     `json:"project_id" form:"project_id"`
	Deadline     int64     `json:"deadline" form:"deadline"`
	Important    int8      `json:"important" form:"important"`
	Done         int8      `json:"done" form:"done"`
	Myday        int8      `json:"myDay" form:"myDay"`
	CreatedTime  time.Time `json:"created_time" form:"created_time"`
	Node         int8      `json:"node" form:"node"`
	Checkpoint   int8      `json:"checkPoint" form:"checkPoint"`
	CollectionId int64     `json:"collection_id" form:"collection_id"`
}

type ProjectNodeInfoResponse struct {
	NodeInfo model.Items `json:"node_info"`
}

type ProjectNodes struct {
	Nodes []model.Items `json:"nodes"`
}

type AllProjectsResponse struct {
	Projects []model.Projects `json:"projects"`
}
