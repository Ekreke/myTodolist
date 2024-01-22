package item

import (
	"github.com/ekreke/myTodolist/internal/mytodolist/store"
	v1 "github.com/ekreke/myTodolist/pkg/api/mytodolist"
)

type itemBiz struct {
	ds store.Istore
}

type ItemBiz interface {
	Create(request *v1.ItemCreateRequest, username string) (resp *v1.CommonResponseWizMsg, err error)
	Delete(itemid int, username string) (resp *v1.CommonResponseWizMsg, err error)
	Info(itemid int, username string) (resp *v1.ItemInfoResponse, err error)
	SetDone(itemid int, username string) (resp *v1.CommonResponseWizMsg, err error)
	SetUnDone(itemid int, username string) (resp *v1.CommonResponseWizMsg, err error)
	Update(request *v1.ItemUpdateRequest, username string) (resp *v1.CommonResponseWizMsg, err error)
}

func New(ds store.Istore) *itemBiz {
	return &itemBiz{ds: ds}
}

var _ ItemBiz = (*itemBiz)(nil)
