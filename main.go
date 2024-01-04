package main

import (
	"github.com/ekreke/myTodolist/conf"
	"github.com/ekreke/myTodolist/server"
)

func main() {
	conf.Init()
	r := server.NewRouter()
	r.Run(":3000")
}
