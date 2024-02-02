package collection

import (
	"github.com/ekreke/myTodolist/internal/pkg/log"
	v1 "github.com/ekreke/myTodolist/pkg/api/mytodolist"
)

func (c *collectionBiz) Delete(r *v1.DeleteCollectionWizIDRequest, username string) (*v1.CommonResponseWizMsg, error) {
	resp, err := c.ds.Collection().Delete(r.CollectionId, username)
	if err != nil {
		log.Errorw("delete collection failed")
		return nil, err
	}
	return resp, nil
}
