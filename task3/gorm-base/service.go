package main

import (
	_ "database/sql"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"
)

var dsn = "root:Ab123456@(127.0.0.1:3306)/blog"
var db *gorm.DB

func UserTestBase() {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	defer db.DB()
	if err != nil {
		fmt.Println(err)
	}
	u := User{}
	db.Where("id = ?", 1).Find(&u)
	fmt.Println(u)
}

func InitDb() {
	// init db config
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("failed to connect database: ", err)
	}
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(5)
	sqlDB.SetMaxOpenConns(3)
	sqlDB.SetConnMaxLifetime(time.Minute)

	// create tables
	err = db.AutoMigrate(&User{}, &Post{}, &Comment{})
	if err != nil {
		log.Fatalf("AutoMigrate error: %v", err)
	}

	fmt.Println("AutoMigrate OK.")

}

func FindByUserId(userID uint) (*User, error) {
	var u User = User{
		ID: userID,
	}
	e := db.Preload("Posts").
		Preload("Posts.Comments").
		First(&u).Error
	if e != nil {
		fmt.Printf("findByUserId error: %v", e)
	}
	return &u, nil
}
