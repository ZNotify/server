package db

import (
	"fmt"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"notify-api/db/entity"
	"os"
)

var DB *gorm.DB

func checkDBFile() {
	fa, err := os.Stat("data")
	if err != nil {
		if os.IsNotExist(err) {
			err = os.Mkdir("data", 0777)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		} else {
			fmt.Println(err)
			os.Exit(1)
		}
	} else {
		if !fa.IsDir() {
			fmt.Println("data is not a directory")
			os.Exit(1)
		}
	}

	// Determining whether notify.db is existed
	va, err := os.Stat("data/notify.db")
	if err != nil {
		if os.IsNotExist(err) {
			// Create notify.db file
			fmt.Println("Creating notify.db file")
			file, err := os.Create("data/notify.db")
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
			fmt.Println("data/notify.db is directory.")
			os.Exit(1)
		}
	}
}

func Init() {
	checkDBFile()
	var err error
	DB, err = gorm.Open(sqlite.Open("data/notify.db"), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	err = DB.AutoMigrate(
		&entity.Message{},
		&entity.FCMTokens{},
		&entity.WebSubscription{},
	)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
