package collection

import (
	"github.com/ekreke/myTodolist/internal/pkg/log"
	v1 "github.com/ekreke/myTodolist/pkg/api/mytodolist"
)

func (c *collectionBiz) Create(r *v1.CollectionCreateRequest, username string) (*v1.CommonResponseWizMsg, error) {
	resp, err := c.ds.Collection().Create(r.Icon, r.Name, username)
	if err != nil {
		log.Errorw("create collection failed")
		return nil, err
	}
	return resp, nil
}
