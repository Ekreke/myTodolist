package model

type Users struct {
	Id          int    `json:"id" sql:"id"`
	Username    string `json:"username" sql:"username"`
	Password    string `json:"password" sql:"password"`
	ApartmentId int    `json:"apartment_id" sql:"apartment_id"`
	ProjectsId  string `json:"projects_id" sql:"projects_id"`
	Link        string `json:"link" sql:"link"`
	Bio         string `json:"bio" sql:"bio"`
	Avatar      string `json:"avatar" sql:"avatar"`
}

func (user *Users) CheckPassword(password string) bool {
	return password == user.Password
}
