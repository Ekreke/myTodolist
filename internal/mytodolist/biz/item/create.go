package item

import (
	"time"

	"github.com/ekreke/myTodolist/internal/pkg/model"
	v1 "github.com/ekreke/myTodolist/pkg/api/mytodolist"
)

func (i *itemBiz) Create(request *v1.ItemCreateRequest, username string) (resp *v1.CommonResponseWizMsg, err error) {
	it := &model.Items{
		ItemName:     request.ItemName,
		Description:  request.Description,
		ProjectId:    request.ProjectId,
		Deadline:     time.Now(),
		Important:    request.Important,
		Done:         request.Done,
		Myday:        request.Myday,
		CreatedTime:  time.Now(),
		Node:         request.Node,
		Checkpoint:   request.Checkpoint,
		CollectionId: request.CollectionId,
	}

	resp, err = i.ds.Items().Create(it, username)
	i.ds.Users()
	if err != nil {
		return nil, err
	}
	return resp, nil
}
