package user

import (
	"github.com/ekreke/myTodolist/internal/mytodolist/biz"

	"github.com/ekreke/myTodolist/internal/mytodolist/store"
)

type UserController struct {
	// casbin
	// grpc
	b biz.IBiz
}

func New(ds store.Istore) *UserController {
	return &UserController{b: biz.NewBiz(ds)}
}

func Helloworld() {

}
