package user

import (
	"context"

	v1 "github.com/ekreke/myTodolist/pkg/api/mytodolist"
)

func (b *userBiz) Info(ctx context.Context, username string) (*v1.InfoResponse, error) {
	inforesp, err := b.ds.Users().GetInfo(username)
	if err != nil {
		return nil, err
	}
	return inforesp, nil
}
