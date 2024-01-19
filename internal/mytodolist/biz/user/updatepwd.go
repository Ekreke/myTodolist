package user

import (
	"context"

	"github.com/ekreke/myTodolist/internal/pkg/errno"
	"github.com/ekreke/myTodolist/internal/pkg/log"
	v1 "github.com/ekreke/myTodolist/pkg/api/mytodolist"
)

func (b *userBiz) UpdatePwd(ctx context.Context, username string, prepwd string, newpwd string) (*v1.CommonResponseWizMsg, error) {
	flag, err := b.ds.Users().CheckUserIfExist(username)
	if err != nil {
		log.Infow("check user if exist failed")
		return nil, err
	}
	if flag == false {
		return nil, errno.ErrUserNotFound
	}
	if prepwd == newpwd {
		return nil, errno.ErrPwdDuplicate
	}
	err = b.ds.Users().UpdatePwd(username, newpwd)
	if err != nil {
		log.C(ctx).Infow("biz update pwd run failed")
		return nil, errno.ErrUpdatePwdFailed

	}
	return &v1.CommonResponseWizMsg{Msg: "update pwd success"}, nil
}
