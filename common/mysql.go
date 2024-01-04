package common

import (
	"fmt"

	"github.com/ekreke/myTodolist/model"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

func InitMysql() *gorm.DB {
	driverName := "mysql"
	host := "127.0.0.1"
	port := "3306"
	database := "myTodolist"
	username := "root"
	password := "root"
	charset := "utf8"
	args := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=%s&parseTime=true",
		username,
		password,
		host,
		port,
		database,
		charset)
	db, err := gorm.Open(driverName, args)
	fmt.Printf("drivername:%v,args:%v", driverName, args)
	if err != nil {
		panic("failed to connect database,err:+" + err.Error())
	}
	db.AutoMigrate(&model.Users{})
	DB = db
	return db
}

func GetDB() *gorm.DB {
	return DB
}
