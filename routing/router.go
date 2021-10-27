package routing

import (
	"caged/loaded"
	"github.com/valyala/fasthttp"
)

type Router struct {
	routes       []Route
	viewEngine   *string
	staticRoutes []string
	module       *loaded.LoadedModule
}

func (router *Router) Handle(ctx *fasthttp.RequestCtx) {
	for i := 0; i < len(router.routes); i++ {
		route := router.routes[i]
		if route.DoRouting(ctx) {
			return
		}
	}
}

func (router *Router) SetViewEngine(engine string) {
	tmp := "." + engine
	router.viewEngine = &tmp
}

func (router *Router) UseStaticAssets(folder string) {
	router.staticRoutes = append(router.staticRoutes, folder)
}

func (router *Router) Setup() {
	for i := 0; i < len(router.staticRoutes); i++ {
		staticRoute := router.staticRoutes[i]
		route := &StaticRoute{static: staticRoute, viewEngine: router.viewEngine}
		route.Init()
		router.routes = append(router.routes, route)
	}
}

func CreateRouter(module *loaded.LoadedModule) *Router {
	router := &Router{module: module}
	router.routes = make([]Route, 0)
	router.staticRoutes = make([]string, 0)
	return router
}
