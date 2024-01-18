package store

import (
	"context"

	"github.com/ekreke/myTodolist/internal/pkg/model"
	v1 "github.com/ekreke/myTodolist/pkg/api/mytodolist"
	"gorm.io/gorm"
)

// UserStore 定义了 user 模块在 store 层所实现的方法.
type UserStore interface {
	Create(ctx context.Context, user *model.Users) error
	Get(ctx context.Context, username string) (*model.Users, error)
	GetInfo(username string) (*v1.InfoResponse, error)
	CheckUserIfExist(username string) (bool, error)
	UpdateInfo(req *v1.UpdateInfoRequest, username string) error
}

// UserStore 接口的实现.
type users struct {
	db *gorm.DB
}

// 确保 users 实现了 UserStore 接口.
var _ UserStore = (*users)(nil)

func newUsers(db *gorm.DB) *users {
	return &users{db}
}

// Create 插入一条 user 记录.
func (u *users) Create(ctx context.Context, user *model.Users) error {
	return u.db.Create(&user).Error
}

// Get 根据用户名查询指定user的数据库记录
func (u *users) Get(ctx context.Context, username string) (*model.Users, error) {
	var user model.Users
	if err := u.db.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *users) GetInfo(username string) (*v1.InfoResponse, error) {
	var user model.Users
	if err := u.db.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	resp := &v1.InfoResponse{
		Username:   user.Username,
		Bio:        user.Bio,
		Link:       user.Link,
		Avatar:     user.Avatar,
		Root:       int(user.Root),
		Created_At: user.CreatedAt,
	}
	return resp, nil
}

// if exist , return true else return false
func (u *users) CheckUserIfExist(username string) (bool, error) {
	var uname string
	if err := u.db.Debug().Where("username = ?", username).Select(&uname).Error; err != nil {
		return false, err
	}
	return true, nil
}

func (u *users) UpdateInfo(req *v1.UpdateInfoRequest, username string) error {
	// user := &model.Users{}
	user := &model.Users{}
	err := u.db.Debug().Where("username = ?", username).First(&user).Error
	if err != nil {
		return err
	}
	if req.Username != "" {
		user.Username = req.Username
	}
	if req.Bio != "" {
		user.Bio = req.Bio
	}
	if req.Avatar != "" {
		user.Avatar = req.Avatar
	}
	if req.Link != "" {
		user.Link = req.Link
	}
	return u.db.Save(&user).Error

}
