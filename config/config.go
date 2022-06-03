package config

import (
	"fmt"
	"os"
)

var MiPushSecret string
var FCMCredential []byte
var VAPIDPublicKey string
var VAPIDPrivateKey string

func CheckMiPushSecret() {
	MiPushSecret = os.Getenv("MiPushSecret")
	if MiPushSecret == "" {
		fmt.Println("MiPushSecret not set")
		os.Exit(1)
	}
}

func CheckFCMCredential() {
	FCMCredential = []byte(os.Getenv("FCMCredential"))
	if len(FCMCredential) == 0 {
		fmt.Println("FCMCredential not set")
		os.Exit(1)
	}
}

func CheckWebPushCert() {
	VAPIDPublicKey = os.Getenv("VAPIDPublicKey")
	VAPIDPrivateKey = os.Getenv("VAPIDPrivateKey")
	if VAPIDPublicKey == "" || VAPIDPrivateKey == "" {
		fmt.Println("VAPIDPublicKey or VAPIDPrivateKey not set")
		os.Exit(1)
	}
}
