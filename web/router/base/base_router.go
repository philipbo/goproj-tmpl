package base

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type Router interface {
	Register(name string, child Router)
	Load(rg *gin.RouterGroup)
}

type BaseRouter struct {
	Name  string
	Child map[string]Router
}

func (r *BaseRouter) Register(name string, child Router) {
	if r.Child == nil {
		r.Child = make(map[string]Router)
	} else {
		_, ok := r.Child[name]
		if ok {
			log.Fatalf("router name are already registered for child router '%s'", name)
		}
	}

	r.Child[name] = child
}

func (r *BaseRouter) Load(rg *gin.RouterGroup) {
	for name, router := range r.Child {
		group := rg.Group(name)
		router.Load(group)
	}
}
