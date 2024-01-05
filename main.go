package main

import (
	"github.com/ekreke/myTodolist/conf"
	"github.com/ekreke/myTodolist/server"
)

// @title 这里写标题
// @version 1.0
// @description 这里写描述信息
// @termsOfService  http://swagger.io/terms/

// @contact.name 这里写联系人名字
// @contact.email 这里写联系人邮箱

func main() {
	conf.DBInit()
	r := server.NewRouter()
	r.Run(":3000")
}
