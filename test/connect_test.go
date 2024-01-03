package main_test

import (
	"fmt"
	"testing"

	"github.com/ekreke/myTodolist/common"
	"github.com/ekreke/myTodolist/model"
	_ "github.com/go-sql-driver/mysql"
)

func TestInitMysql(t *testing.T) {
	db := common.InitMysql()

	// Check if the returned value is not nil
	if db == nil {
		t.Errorf("InitMysql() returned nil")
	}

	// Check if the connection was successfully established
	if err := db.Error; err != nil {
		t.Errorf("Failed to connect to database: %v", err)
	}

	// Check if AutoMigrate was called
	if !db.HasTable(&model.Users{}) {
		t.Errorf("AutoMigrate was not called for model.User")
	}

	fmt.Println(db)
	// // Check if the global variable DB was set correctly
	// if DB != db {
	// 	t.Errorf("DB was not set correctly")
	// }
	user := &model.Users{}
	d := db.First(user)
	if d != nil {
		t.Errorf("Failed to find user: %v", d)
	}
	fmt.Println("the user is:", d)
}
func TestInitMysql2(t *testing.T) {
	fmt.Println("This is a test")

	// ... rest of your test code ...
}
