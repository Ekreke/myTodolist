package user

import (
	"context"

	"github.com/ekreke/myTodolist/internal/mytodolist/store"
	"github.com/ekreke/myTodolist/internal/pkg/errno"
	v1 "github.com/ekreke/myTodolist/pkg/api/mytodolist"
	"github.com/ekreke/myTodolist/pkg/token"
)

type UserBiz interface {
	Login(ctx context.Context, r *v1.LoginRequest) (*v1.LoginResponse, error)
	// Create(ctx context.Context, username string, r *v1.CreateRequest) (*v1.CreateResponse, error)
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

func (b *userBiz) Login(ctx context.Context, r *v1.LoginRequest) (*v1.LoginResponse, error) {
	user, err := b.ds.Users().Get(ctx, r.Username)
	if err != nil {
		return nil, errno.ErrUserNotFound
	}
	if user.Password != r.Password {
		return nil, errno.ErrUserNotFound
	}
	token, err := token.Sign(r.Username)
	if err != nil {
		return nil, errno.ErrSignToken
	}
	return &v1.LoginResponse{Token: token}, nil
}
