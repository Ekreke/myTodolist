package collection

import (
	"github.com/ekreke/myTodolist/internal/pkg/log"
	v1 "github.com/ekreke/myTodolist/pkg/api/mytodolist"
)

func (c *collectionBiz) AddItem(r *v1.CollectionAddItemRequest, username string) (*v1.CommonResponseWizMsg, error) {
	resp, err := c.ds.Collection().AddItem(r.ItemId, r.CollectionId, username)
	if err != nil {
		log.Errorw("add item failed")
		return nil, err

	}
	return resp, nil
}
