package project

import (
	"github.com/ekreke/myTodolist/internal/mytodolist/store"
	v1 "github.com/ekreke/myTodolist/pkg/api/mytodolist"
)

type ProjectBiz interface {
	Join(userid int64, projectid int64, pwd string) (*v1.CommonResponseWizMsg, error)
	Myprojects(userid int64) (*v1.MyprojectsResponse, error)
	Quit(userid int64, projectid int64) (*v1.CommonResponseWizMsg, error)
}

type projectBiz struct {
	ds store.Istore
}

var _ ProjectBiz = (*projectBiz)(nil)

func New(ds store.Istore) *projectBiz {
	return &projectBiz{ds: ds}
}
