package version

var (
	GVersion Version
)

type Version struct {
	Who       string `json:"who"`
	Build     string `json:"build"`
	Branch    string `json:"branch"`
	GoVersion string `json:"go_version"`
	Time      string `json:"time"`
}

func NewVersion(who, buildTime, build, branch, goVersion string) {
	GVersion = Version{
		Who:       who,
		Build:     build,
		Branch:    branch,
		GoVersion: goVersion,
		Time:      buildTime,
	}
}
