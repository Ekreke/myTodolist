package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type Users struct {
	gorm.Model
	Id          int    `json:"id" sql:"id"`
	Username    string `json:"username" sql:"username"`
	Password    string `json:"password" sql:"password"`
	ApartmentId int    `json:"apartment_id" sql:"apartment_id"`
	ProjectsId  string `json:"projects_id" sql:"projects_id"`
	Link        string `json:"link" sql:"link"`
	Bio         string `json:"bio" sql:"bio"`
	Avatar      string `json:"avatar" sql:"avatar"`
}

func main() {
	db, err := gorm.Open("mysql", "root:root@tcp(127.0.0.1:3306)/mytodolist?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err)
		return
	}
	db.AutoMigrate(&Users{})
	u := &Users{}
	db.First(u)
	fmt.Println(u)
}
