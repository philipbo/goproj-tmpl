package main

import (
	_ "expvar"
	log "github.com/sirupsen/logrus"
	"goproj-tmpl/app"
	"goproj-tmpl/version"
	"goproj-tmpl/web"
	"net/http"
	_ "net/http/pprof"
)

var (
	BuildTime string
	GitHash   string
	GitBranch string
	GoVersion string
)

func main() {
	log.Printf("\nVersion:\nbuild time: %s\ngit hash: %s\ngit branch: %s\ngo version: %s",
		BuildTime, GitHash, GitBranch, GoVersion)

	version.NewVersion(BuildTime, GitHash, GitBranch, GoVersion)

	if err := app.GetApp().Init(); err != nil {
		log.Fatal(err)
	}

	http.Handle("/", web.Mount())
	addr := app.GetApp().Config.Addr
	log.Infof("start web server on %s", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
