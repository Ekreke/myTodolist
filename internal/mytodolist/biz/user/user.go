package user

import (
	"context"

	"github.com/ekreke/myTodolist/internal/mytodolist/store"
	v1 "github.com/ekreke/myTodolist/pkg/api/mytodolist"
)

type UserBiz interface {
	Login(ctx context.Context, r *v1.LoginRequest) (*v1.LoginResponse, error)
	Register(ctx context.Context, r *v1.RegisterRequest) (*v1.RegisterResponse, error)
	// Get(ctx context.Context, username string, r *v1.GetRequest) (*v1.GetResponse, error)
	// Delete(ctx context.Context, username string, r *v1.DeleteRequest) (*v1.DeleteResponse, error)
	// Update(ctx context.Context, username string, r *v1.UpdateRequest) (*v1.UpdateResponse, error)
}

type userBiz struct {
	ds store.Istore
}

var _ UserBiz = (*userBiz)(nil)

func New(ds store.Istore) *userBiz {
	return &userBiz{ds: ds}
}
