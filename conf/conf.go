package conf

import (
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/subosito/gotenv"
)

// conf
var DB *gorm.DB

func DBInit() {
	drivername, args := getMysqlConf()
	db := myqlInit(drivername, args)
	DB = db
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
	err := gotenv.Load("conf/confs/mysql.env")
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
