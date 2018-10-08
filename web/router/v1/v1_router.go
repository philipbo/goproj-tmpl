package v1

import (
	"web-temp/web/router/base"
)

func NewV1Router() *V1Router {
	return (&V1Router{}).init()
}

type V1Router struct {
	base.BaseRouter
}

func (r *V1Router) init() *V1Router {
	r.Register("/example", NewExampleRouter())
	return r
}
