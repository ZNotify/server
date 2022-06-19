package entity

import (
	"fmt"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"notify-api/db"
	"os"
)

func checkDBFile() {
	// Determining whether notify.db is existed
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

func Init() {
	checkDBFile()
	var err error
	db.DB, err = gorm.Open(sqlite.Open("notify.db"), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	err = db.DB.AutoMigrate(&Message{}, &FCMTokens{}, &WebSubscription{})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
