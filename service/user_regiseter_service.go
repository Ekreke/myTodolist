package service

import "github.com/ekreke/myTodolist/serializer"

// user register service

type UserRegisterService struct {
}

func (service UserRegisterService) Register() serializer.Response {
	return serializer.Response{
		Status: 200,
		Msg:    "success",
	}
}
