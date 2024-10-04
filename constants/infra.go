package constants

import (
	"fmt"
	"log"
	"strings"

	"github.com/spf13/cobra"
)

// AppInfo application info structure
type AppInfo struct {
	AppName        string
	AppVersion     string
	AppCommit      string
	BuildGoVersion string
	BuildArch      string
	BuildDate      string
}

var (
	// meta
	AppMeta *AppInfo

	// root command
	RootCommands = &cobra.Command{
		Use:   "startup-engine-go-sdk",
		Short: "Startup Engine Go SDK",
		Long:  "Startup Enginer is Epic!",
	}

	// version sub command
	VersionCommands = &cobra.Command{
		Use:   "version",
		Short: "Print version info",
		Long:  "Print version information of api",
		Run: func(command *cobra.Command, args []string) {
			infoStr := strings.Builder{}
			infoStr.WriteString(fmt.Sprintf("%s - api version info:\n", AppMeta.AppName))
			infoStr.WriteString(fmt.Sprintf("Version:\t%s\n", AppMeta.AppVersion))
			infoStr.WriteString(fmt.Sprintf("Commit Hash:\t%s\n", AppMeta.AppCommit))
			infoStr.WriteString(fmt.Sprintf("Go Version:\t%s\n", AppMeta.BuildGoVersion))
			infoStr.WriteString(fmt.Sprintf("Arch:\t\t%s\n", AppMeta.BuildArch))
			infoStr.WriteString(fmt.Sprintf("Build:\t\t%s\n", strings.Replace(AppMeta.BuildDate, "_", " ", -1)))
			log.Println(infoStr.String())
		},
	}

	HTTPCommands = &cobra.Command{
		Use:   "serve-http",
		Short: "Run http server",
		Long:  "API",
	}
)
