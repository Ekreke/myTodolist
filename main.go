package main

import (
	"github.com/ekreke/myTodolist/conf"
	"github.com/ekreke/myTodolist/server"
)

// @title myTodolist
// @version 1.0
// @description myTodoList's api docs
// @contact.name ekreke
// should execute swag init everytimes when update annotation
func main() {
	conf.DBInit()
	r := server.NewRouter()
	r.Run(":3000")
}
