package app

import (
	"base-api/app/http"
	"base-api/constants"
	"base-api/infra/http_server"
	"os"
)

var (
	appMeta      *constants.AppInfo
	httpCommands = constants.HTTPCommands
)

func init() {
	httpCommands.RunE = http_server.New().RunHTTP
	constants.RootCommands.AddCommand(constants.VersionCommands)
	constants.RootCommands.AddCommand(constants.HTTPCommands)
	constants.RootCommands.AddCommand(http.ServeGRPC())
}

// Execute run root command
func Execute(appInfo *constants.AppInfo) {
	appMeta = appInfo
	if err := constants.RootCommands.Execute(); err != nil {
		os.Exit(1)
	}
}

// GetAppInfo return application information
func GetAppInfo() *constants.AppInfo {
	return appMeta
}
