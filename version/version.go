package version

var (
	GVersion Version
)

type Version struct {
	GitHash   string `json:"git_hash"`
	GitBranch string `json:"git_branch"`
	GoVersion string `json:"go_version"`
	BuildTime string `json:"build_time"`
}

func NewVersion(buildTime, githash, gitbranch, goversion string) {
	GVersion = Version{
		GitHash:   githash,
		GitBranch: gitbranch,
		GoVersion: goversion,
		BuildTime: buildTime,
	}
}
