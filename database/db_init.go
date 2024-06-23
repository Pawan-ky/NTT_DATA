package database

import (
	"fmt"

	// "gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// var logger = zapLog.GetLogger
var Instance *gorm.DB
var dbError error

// func Init() *gorm.DB {
// 	dbURI := "root:root@123@tcp(localhost:3306)/demo_db?charset=utf8&parseTime=True&loc=Local"
// 	fmt.Println(dbURI)
// 	dbInstance, err := gorm.Open(mysql.Open(dbURI), &gorm.Config{
// 	})
// 	if err != nil {
// 		panic("Failed to connect to the database")
// 	}
// 	Instance = dbInstance
// 	return Instance
// }

func Init() *gorm.DB {

    db, err := gorm.Open(sqlite.Open("demo.db"), &gorm.Config{})
    if err != nil {
        panic("failed to connect database")
    }
	Instance = db
	return Instance
}

func Migrate() {
	err := Instance.AutoMigrate(&ExoplanetType{}, &Exoplanet{})
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Database Migration Completed!")
}