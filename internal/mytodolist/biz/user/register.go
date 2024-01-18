package user

import (
	"context"
	"time"

	"github.com/ekreke/myTodolist/internal/pkg/errno"
	"github.com/ekreke/myTodolist/internal/pkg/model"
	v1 "github.com/ekreke/myTodolist/pkg/api/mytodolist"
)

func (b *userBiz) Register(ctx context.Context, r *v1.RegisterRequest) (*v1.RegisterResponse, error) {
	u := &model.Users{Username: r.Username,
		Password:  r.Password,
		Bio:       r.Bio,
		Link:      r.Link,
		CreatedAt: time.Now(),
		DeletedAt: time.Now().AddDate(1, 0, 0),
	}
	flag, _ := b.ds.Users().CheckUserIfExist(r.Username)
	if flag {
		return nil, errno.ErrUserAlreadyExist
	}
	err := b.ds.Users().Create(ctx, u)
	if err != nil {
		return nil, errno.ErrUserCreateFailed
	}
	return &v1.RegisterResponse{Msg: "success"}, nil
}
