package serializer

import "github.com/ekreke/myTodolist/model"

type User struct {
	ID          int    `json:"id"`
	UserName    string `json:"username"`
	Password    string `json:"password"`
	ApartmentId int    `json:"apartment_id"`
	ProjectsId  string `json:"projects_id"`
	Link        string `json:"link"`
	Bio         string `json:"bio"`
	Avatar      string `json:"avatar"`
}

func BuildUser(user model.Users) User {
	return User{
		ID:          user.Id,
		UserName:    user.Username,
		Password:    user.Password,
		ApartmentId: user.ApartmentId,
		ProjectsId:  user.ProjectsId,
		Link:        user.Link,
		Bio:         user.Bio,
		Avatar:      user.Avatar,
	}
}
