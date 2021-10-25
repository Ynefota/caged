package routing

import (
	"caged/loaded"
	"github.com/valyala/fasthttp"
)

type Router struct {
	routes        []Route
	staticFolders []string
	module        *loaded.LoadedModule
}

func (router *Router) Handle(ctx *fasthttp.RequestCtx) {
	for i := 0; i < len(router.routes); i++ {
		route := router.routes[i]
		if route.Match(ctx.Path()) {

		}
	}
}

func (router *Router) UseStaticAssets(folder string) {
	router.staticFolders = append(router.staticFolders, folder)
}

func (router *Router) Setup() {

}

func CreateRouter(module *loaded.LoadedModule) *Router {
	router := &Router{module: module}

	return router
}
