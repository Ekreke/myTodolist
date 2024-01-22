package item

import (
	"github.com/ekreke/myTodolist/internal/pkg/log"
	v1 "github.com/ekreke/myTodolist/pkg/api/mytodolist"
)

func (i *itemBiz) SetUnDone(itemid int, username string) (resp *v1.CommonResponseWizMsg, err error) {
	resp, err = i.ds.Items().SetUnDone(itemid, username)
	if err != nil {
		log.Errorw("ItemBiz SetUndone Failed")
		return nil, err
	}
	return resp, nil
}
