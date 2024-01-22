package item

import (
	"github.com/ekreke/myTodolist/internal/pkg/log"
	v1 "github.com/ekreke/myTodolist/pkg/api/mytodolist"
)

func (i *itemBiz) SetDone(itemid int, username string) (resp *v1.CommonResponseWizMsg, err error) {
	resp, err = i.ds.Items().SetDone(itemid, username)
	if err != nil {
		log.Errorw("ItemBiz SetDone failed")
		return nil, err
	} 
	return resp, nil 

}
