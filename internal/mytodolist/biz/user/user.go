package user

import (
	"context"

	"github.com/ekreke/myTodolist/internal/mytodolist/store"
	v1 "github.com/ekreke/myTodolist/pkg/api/mytodolist"
	"github.com/ekreke/myTodolist/pkg/token"
)

type UserBiz interface {
	CheckUserIfExist(username string) (bool, error)
	Login(ctx context.Context, r *v1.LoginRequest) (*v1.LoginResponse, error)
	Register(ctx context.Context, r *v1.RegisterRequest) (*v1.RegisterResponse, error)
	Info(ctx context.Context, username string) (*v1.InfoResponse, error)
	UpdateInfo(ctx context.Context, req *v1.UpdateInfoRequest, username string) (*v1.CommonResponseWizMsg, error)
	LoadImportantItems(ctx context.Context, req *v1.ImportantRequest, username string) (*v1.ImportantResponse, token.Token, error)
	UpdatePwd(ctx context.Context, username string, prepwd string, newpwd string) (*v1.CommonResponseWizMsg, error)
	GetCollctions(ctx context.Context, username string) (*v1.CollectionsResponse, error)
	// Get(ctx context.Context, username string, r *v1.GetRequest) (*v1.GetResponse, error)
	// Delete(ctx context.Context, username string, r *v1.DeleteRequest) (*v1.DeleteResponse, error)
}

type userBiz struct {
	ds store.Istore
}

var _ UserBiz = (*userBiz)(nil)

func New(ds store.Istore) *userBiz {
	return &userBiz{ds: ds}
}

func (u *userBiz) CheckUserIfExist(username string) (bool, error) {
	return u.ds.Users().CheckUserIfExist(username)
}
