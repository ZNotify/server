package config

import (
	"context"
	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/messaging"
	"fmt"
	"google.golang.org/api/option"
	"io/ioutil"
	"net/http"
	"os"
)

var MiPushSecret string
var FCMCredential []byte

var MiPushClient *http.Client
var FCMClient *messaging.Client

func CheckMiPushSecret() {
	if len(os.Args) <= 1 {
		fmt.Println("MiPush Secret not set.")
		os.Exit(-1)
	}
	MiPushSecret = os.Args[1]
	initMiPushClient()
}

func CheckFCMCredential() {
	_, err := os.Stat("notify.json")
	if err != nil {
		fmt.Println("notify.json not found")
		os.Exit(1)
	}
	FCMCredential, err = ioutil.ReadFile("notify.json")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	initFCMClient()
}

func initMiPushClient() {
	MiPushClient = &http.Client{}
}

func initFCMClient() {
	opt := option.WithCredentialsJSON(FCMCredential)
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		fmt.Println(fmt.Errorf("error initializing app: %v", err))
		os.Exit(1)
	}
	FCMClient, err = app.Messaging(context.Background())
	if err != nil {
		fmt.Println(fmt.Errorf("error initializing app: %v", err))
		os.Exit(1)
	}
}
