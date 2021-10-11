package routing

import (
	"caged/loaded"
	"github.com/valyala/fasthttp"
)

type Router struct {
	routes []Route
	module *loaded.LoadedModule
}

func (r *Router) Handle(ctx *fasthttp.RequestCtx) {

}

type Route struct {
}

func CreateRouter(module *loaded.LoadedModule) *Router {
	router := &Router{module: module}
	return router
}
