package mytodolist

import "time"

type ItemCreateRequest struct {
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

type ItemRequesWizID struct {
	ItemID int64 `json:"item_id"`
}
type ItemInfoResponse struct {
	ItemName     string    `json:"item_name" gorm:"item_name"`
	Description  string    `json:"description" gorm:"description"`
	ProjectId    int64     `json:"project_id" gorm:"project_id"`
	Deadline     time.Time `json:"deadline" gorm:"deadline"`
	Important    int8      `json:"important" gorm:"important"`
	Done         int8      `json:"done" gorm:"done"`
	Myday        int8      `json:"myDay" gorm:"myDay"`
	CreatedTime  time.Time `json:"created_time" gorm:"created_time"`
	Node         int8      `json:"node" gorm:"node"`
	Checkpoint   int8      `json:"checkPoint" gorm:"checkPoint"`
	CollectionId int64     `json:"collection_id" gorm:"collection_id"`
}

type ItemUpdateRequest struct {
	ItemName     string    `json:"item_name" gorm:"item_name"`
	Description  string    `json:"description" gorm:"description"`
	ProjectId    int64     `json:"project_id" gorm:"project_id"`
	Deadline     time.Time `json:"deadline" gorm:"deadline"`
	Important    int8      `json:"important" gorm:"important"`
	Done         int8      `json:"done" gorm:"done"`
	Myday        int8      `json:"myDay" gorm:"myDay"`
	CreatedTime  time.Time `json:"created_time" gorm:"created_time"`
	Node         int8      `json:"node" gorm:"node"`
	Checkpoint   int8      `json:"checkPoint" gorm:"checkPoint"`
	CollectionId int64     `json:"collection_id" gorm:"collection_id"`
}
