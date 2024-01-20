package mytodolist

import (
	"time"

	"github.com/ekreke/myTodolist/internal/pkg/model"
)

var pagination string

type LoginRequest struct {
	Username string `json:"username" form:"username" valid:"alphanum,required,stringlength(1|255)"`
	Password string `json:"password" form:"password" valid:"required,stringlength(6|18)"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type InfoRequest struct {
	Token string `json:"token"`
}

type RegisterRequest struct {
	Username string `form:"username" valid:"alphanum,required,stringlength(1|255)"`
	Password string `form:"password" valid:"required,stringlength(6|18)"`
	Bio      string `form:"bio" valid:"stringlength(0|20)"`
	Link     string `form:"link" valid:"stringlength(0|20)"`
}

type RegisterResponse struct {
	Msg string `json:"msg"`
}

type InfoResponse struct {
	Username   string    `json:"username"`
	Bio        string    `json:"bio"`
	Link       string    `json:"link"`
	Avatar     string    `json:"avatar"`
	Root       int       `json:"root"`
	Created_At time.Time `json:"created_at"`
}

type UpdateInfoRequest struct {
	Username string `form:"username"`
	Bio      string `form:"bio"`
	Avatar   string `form:"avatar"`
	Link     string `form:"link"`
}

type UpdateInfoResponse struct {
	Msg string `json:"msg"`
}

type ImportantRequest struct {
	Pagination string `form:"pagination"`
}

type ImportantResponse struct {
	ImportantItems []ItemInfo `json:"items"`
	CursorToken    string     `json:"cursor_token"`
}

type ItemInfo struct {
	ID          int64     `json:"id" gorm:"id"`
	ItemName    string    `json:"item_name" gorm:"item_name"`
	Description string    `json:"description" gorm:"description"`
	ProjectId   int64     `json:"project_id" gorm:"project_id"`
	Deadline    time.Time `json:"deadline" gorm:"deadline"`
	Done        int8      `json:"done" gorm:"done"`
	CreatedTime time.Time `json:"created_time" gorm:"created_time"`
}

type UpdatepwdRequest struct {
	Prepwd string `form:"prepwd"`
	Newpwd string `form:"newpwd"` // 新密码
}

type CollectionsResponse struct {
	Collections []model.Collections `json:"collections"`
}
