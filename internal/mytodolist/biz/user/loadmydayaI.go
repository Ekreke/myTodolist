package user

import (
	"context"

	v1 "github.com/ekreke/myTodolist/pkg/api/mytodolist"
)

func (b *userBiz) LoadMydayAI(ctx context.Context, username string) (*v1.CommonResponseWizMsg, error) {
	return &v1.CommonResponseWizMsg{}, nil
}
