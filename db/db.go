package db

import (
	"fmt"
	"github.com/ZNotify/server/db/entity"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"os"
)

var DB *gorm.DB

func checkDB() {
	// Determining whether notify.db is directory
	va, err := os.Stat("notify.db")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if va.IsDir() {
		fmt.Println("notify.db is directory.")
		os.Exit(1)
	}
}

func InitDB() {
	checkDB()
	DB, err := gorm.Open(sqlite.Open("notify.db"), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	err = DB.AutoMigrate(&entity.Message{}, &entity.FCMTokens{})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
