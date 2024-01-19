package item

import (
	"github.com/ekreke/myTodolist/internal/mytodolist/biz"
	"github.com/ekreke/myTodolist/internal/mytodolist/store"
)

type ItemController struct {
	// casbin
	// grpc
	b biz.IBiz
}

type IItemController interface {
}

var _ IItemController = (*ItemController)(nil)

func New(ds store.Istore) *ItemController {
	return &ItemController{b: biz.NewBiz(ds)}
}
