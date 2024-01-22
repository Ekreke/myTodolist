package item

import (
	"github.com/ekreke/myTodolist/internal/pkg/log"
	v1 "github.com/ekreke/myTodolist/pkg/api/mytodolist"
)

func (i *itemBiz) Info(itemid int, username string) (resp *v1.ItemInfoResponse, err error) {
	resp, err = i.ds.Items().Info(itemid, username)
	if err != nil {
		log.Errorw("ItemBiz Info Failed")
		return nil, err
	}
	return resp, nil
}
