package v1

type LoginRequest struct {
	Username string `json:"username" form:"username" valid:"alphanum,required,stringlength(1|255)"`
	Password string `json:"password" form:"password" valid:"required,stringlength(6|18)"`
}

type LoginResponse struct {
	Token string `json:"token"`
}
