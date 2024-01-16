package biz

import (
	"github.com/ekreke/myTodolist/internal/mytodolist/biz/user"
	"github.com/ekreke/myTodolist/internal/mytodolist/store"
)

type IBiz interface {
	Users() user.UserBiz
}

// 确保 biz 实现了 IBiz 接口.
var _ IBiz = (*biz)(nil)

// biz 是 IBiz 的一个具体实现.
type biz struct {
	ds store.Istore
}

func NewBiz(ds store.Istore) *biz {
	return &biz{ds: ds}
}

// Users 返回一个实现了 UserBiz 接口的实例.
func (b *biz) Users() user.UserBiz {
	return user.New(b.ds)
}
