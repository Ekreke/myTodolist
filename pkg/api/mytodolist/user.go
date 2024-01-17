package v1

import "time"

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
