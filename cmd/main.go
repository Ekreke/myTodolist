package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
)

const MySQLDSN = "root:root@(localhost:3306)/mytodolist?charset=utf8mb4&parseTime=True&loc=Local"

func main() {
	db, err := gorm.Open(mysql.Open(MySQLDSN))
	if err != nil {
		panic(fmt.Errorf("cannot establish db connection: %w", err))
	}
	g := gen.NewGenerator(gen.Config{
		OutPath:          "internal/pkg/model",
		Mode:             gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // generate mode
		FieldWithTypeTag: true,
	})
	g.UseDB(db)
	g.ApplyBasic(
		g.GenerateAllTable()...,
	)
	g.Execute()

}
