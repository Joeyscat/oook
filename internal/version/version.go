package version

import "fmt"

var (
	Version   = "unknown"
	GitCommit string
	GoVersion string
	BuildTime string
)

func FullVersion() string {
	return fmt.Sprintf("Version: %6s\nGit commit: %6s\nGo version: %6s\nBuild time: %6s",
		Version, GitCommit, GoVersion, BuildTime)
}
