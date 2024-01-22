package item

import (
	"github.com/ekreke/myTodolist/internal/pkg/log"
	v1 "github.com/ekreke/myTodolist/pkg/api/mytodolist"
)

func (i *itemBiz) Delete(itemid int, username string) (resp *v1.CommonResponseWizMsg, err error) {
	resp, err = i.ds.Items().Delete(itemid, username)
	if err != nil {
		log.Errorw("item store delete failed")
		return nil, err
	}
	return resp, nil
}
