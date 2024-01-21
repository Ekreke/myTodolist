package collection

import (
	"github.com/ekreke/myTodolist/internal/mytodolist/biz"
	"github.com/ekreke/myTodolist/internal/mytodolist/store"
)

type CollectionController struct {
	// casbin
	// grpc
	b biz.IBiz
}

type ICollectinoController interface {
}

var _ ICollectinoController = (*CollectionController)(nil)

func New(ds store.Istore) *CollectionController {
	return &CollectionController{b: biz.NewBiz(ds)}
}
