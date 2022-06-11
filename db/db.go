package db

import (
	"fmt"
	"github.com/ZNotify/server/db/entity"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"os"
)

var DB *gorm.DB

func checkDBFile() {
	// Determining whether notify.db is exist
	va, err := os.Stat("notify.db")
	if err != nil {
		if os.IsNotExist(err) {
			// Create notify.db file
			fmt.Println("Creating notify.db file")
			file, err := os.Create("notify.db")
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			err = file.Close()
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		}
	} else {
		if va.IsDir() {
			fmt.Println("notify.db is directory.")
			os.Exit(1)
		}
	}
}

func InitDB() {
	checkDBFile()
	var err error
	DB, err = gorm.Open(sqlite.Open("notify.db"), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	err = DB.AutoMigrate(&entity.Message{}, &entity.FCMTokens{}, &entity.WebSubscription{})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
