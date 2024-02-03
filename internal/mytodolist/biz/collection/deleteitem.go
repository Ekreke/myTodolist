package collection

import (
	"github.com/ekreke/myTodolist/internal/pkg/log"
	v1 "github.com/ekreke/myTodolist/pkg/api/mytodolist"
)

func (c *collectionBiz) DeleteItem(r *v1.CollectionDeleteItemRequest, username string) (*v1.CommonResponseWizMsg, error) {
	resp, err := c.ds.Collection().DeleteItem(r.ItemId, r.CollectionId, username)
	if err != nil {
		log.Errorw("delete item from collection failed")
		return nil, err
	}
	return resp, nil
}
