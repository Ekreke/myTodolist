package collection

import (
	"github.com/ekreke/myTodolist/internal/pkg/log"
	v1 "github.com/ekreke/myTodolist/pkg/api/mytodolist"
)

func (c *collectionBiz) LoadItems(r *v1.CollectionLoadItemsWizIdRequest, username string) (*v1.CollectionLoadItemsResp, error) {
	resp , err := c.ds.Collection().LoadItems(r.CollectionId, username) 
	if err != nil {
		log.Errorw("load items from collection failed")
		return nil ,err
	} 
	return resp , nil
}
