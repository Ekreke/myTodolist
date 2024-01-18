package user

import (
	"context"

	"github.com/ekreke/myTodolist/internal/pkg/errno"
	v1 "github.com/ekreke/myTodolist/pkg/api/mytodolist"
)

func (b *userBiz) UpdateInfo(ctx context.Context, req *v1.UpdateInfoRequest, username string) (v1.CommonResponseWizMsg, error) {
	if req.Username == "" || req.Username == username {
		return v1.CommonResponseWizMsg{Msg: "input username invalid"}, nil
	}
	flag, _ := b.CheckUserIfExist(req.Username)
	// if true , user exist
	if flag {
		return v1.CommonResponseWizMsg{Msg: errno.ErrUserAlreadyExist.Message}, nil
	}
	err := b.ds.Users().UpdateInfo(req, username)
	if err != nil {
		return v1.CommonResponseWizMsg{Msg: errno.ErrUserUpdateFailed.Message}, nil
	}
	return v1.CommonResponseWizMsg{Msg: "success"}, nil

}
