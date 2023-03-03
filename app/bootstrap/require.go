package bootstrap

import (
	"net/http"
	"os"
	"runtime"
	"strconv"
	"strings"

	"go.uber.org/zap"

	"github.com/ZNotify/server/app/global"
)

func RequireNetwork() {
	if !global.IsProd() {
		zap.S().Info("Skip connection check in non-production mode")
		return
	}

	go func() {
		_, err := http.Get("https://www.google.com/robots.txt")
		if err != nil {
			zap.S().Panicf("Failed to connect to internet: %v", err)
		}
	}()
}

func RequireX64() {
	if strconv.IntSize != 64 {
		zap.S().Panic("Only 64-bit platform is supported")
	}
}

func RequireRootOnLinux() {
	portStr := strings.Split(global.App.Config.Server.Address, ":")[1]
	port, err := strconv.Atoi(portStr)
	if err != nil {
		zap.S().Panicf("Failed to parse port: %v", err)
	}
	//goland:noinspection GoBoolExpressions
	if runtime.GOOS != "linux" {
		return
	}
	if port > 1024 {
		return
	}
	euid := os.Geteuid()
	if euid == -1 {
		zap.S().Panic("Get euid -1 on linux platform %s", runtime.GOOS)
	}
	if euid != 0 {
		zap.S().Panic("Please run with root privilege if you want to listen on port < 1024")
	}
}

func checkRequirements() {
	RequireNetwork()
	RequireX64()
	RequireRootOnLinux()
}
