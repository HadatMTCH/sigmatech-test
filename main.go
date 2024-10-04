package main

import (
	"base-api/app"
	"base-api/constants"
	"fmt"
	"runtime"
	"time"
)

var (
	// application metadata
	appName    = "api"
	appVersion = "development"
	appCommit  = "xxxxxxx"
	goVersion  = runtime.Version()
	buildDate  = time.Now().UTC().Format("2006-01-02_15:04:05_UTC")
	buildArch  = fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH)
)

func getAppInfo() *constants.AppInfo {
	if appVersion == "" {
		appVersion = "0.0.1"
	}

	return &constants.AppInfo{
		AppName:        appName,
		AppVersion:     appVersion,
		AppCommit:      appCommit,
		BuildGoVersion: goVersion,
		BuildArch:      buildArch,
		BuildDate:      buildDate,
	}
}

func main() {
	app.Execute(getAppInfo())
}
