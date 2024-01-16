package store

import (
	"context"

	"github.com/ekreke/myTodolist/internal/pkg/model"
	"gorm.io/gorm"
)

// UserStore 定义了 user 模块在 store 层所实现的方法.
type UserStore interface {
	Create(ctx context.Context, user *model.Users) error
	Get(ctx context.Context, username string) (*model.Users, error)
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
