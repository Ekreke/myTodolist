package conf

import (
	"fmt"
	"log"
	"os"

	"github.com/ekreke/myTodolist/model"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/subosito/gotenv"
)

// conf
var DB *gorm.DB

func Init() {
	drivername, args := getMysqlConf()
	db := myqlInit(drivername, args)
	defer db.Close()
	if db != nil {
		fmt.Println("connect mysql success")
	}
	u := &model.Users{}
	db.First(u)
	fmt.Println(u.Username)

}

func myqlInit(drivername string, args string) *gorm.DB {
	db, err := gorm.Open(drivername, args)
	if err != nil {
		log.Fatal("connect mysql failed, err:", err)
	}
	DB = db
	fmt.Println("args :", args)
	return db
}

func getMysqlConf() (string, string) {
	err := gotenv.Load("./confs/mysql.env")
	if err != nil {
		log.Fatal("load env failed, err:", err)
	}
	drivername := os.Getenv("driverName")
	host := os.Getenv("host")
	port := os.Getenv("port")
	database := os.Getenv("database")
	username := os.Getenv("mysqlusername")
	password := os.Getenv("password")
	charset := os.Getenv("charset")
	args := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=%s&parseTime=true",
		username,
		password,
		host,
		port,
		database,
		charset)
	return drivername, args
}
