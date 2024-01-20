package t

import (
	"github.com/ekreke/myTodolist/internal/mytodolist/store"
)

type TestController struct {
}

func New(ds store.Istore) *TestController {
	return &TestController{}
}
