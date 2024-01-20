package user

import (
	"context"

	v1 "github.com/ekreke/myTodolist/pkg/api/mytodolist"
)

func (u *userBiz) GetCollctions(ctx context.Context, username string) (*v1.CollectionsResponse, error) {
	resp, err := u.ds.Users().GetCollections(username)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
