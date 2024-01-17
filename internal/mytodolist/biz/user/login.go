package user

import (
	"context"

	"github.com/ekreke/myTodolist/internal/pkg/errno"
	v1 "github.com/ekreke/myTodolist/pkg/api/mytodolist"
	"github.com/ekreke/myTodolist/pkg/token"
)

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
