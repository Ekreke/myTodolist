package item

import (
	"github.com/ekreke/myTodolist/internal/pkg/log"
	"github.com/ekreke/myTodolist/internal/pkg/model"
	v1 "github.com/ekreke/myTodolist/pkg/api/mytodolist"
)

func (i *itemBiz) Update(request *v1.ItemUpdateRequest, username string) (resp *v1.CommonResponseWizMsg, err error) {
	it := &model.Items{
		ItemName:     request.ItemName,
		Description:  request.Description,
		ProjectId:    request.ProjectId,
		Deadline:     request.Deadline,
		Important:    request.Important,
		Done:         request.Done,
		Myday:        request.Myday,
		CreatedTime:  request.CreatedTime,
		Node:         request.Node,
		Checkpoint:   request.Checkpoint,
		CollectionId: request.CollectionId,
	}
	resp, err = i.ds.Items().Update(it, username)
	if err != nil {
		log.Errorw("Item Biz Update Failed")
		return nil, err

	}
	return resp, nil
}
