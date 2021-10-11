package application

// TODO get information on https://github.com/valyala/fasthttp#switching-from-nethttp-to-fasthttp

import (
	"caged/base"
	"caged/inject"
	"caged/loaded"
	"caged/routing"
	"github.com/valyala/fasthttp"
	"strconv"
)

type Application struct {
	port         int
	router       *routing.Router
	module       *base.Module
	loadedModule *loaded.LoadedModule
}

func Create(module *base.Module) Application {
	app := Application{}
	app.module = module
	app.loadedModule = inject.LoadModule(module)
	app.router = routing.CreateRouter(app.loadedModule)
	return app
}

func (app *Application) Listen(port int) {
	app.port = port
	strPort := ":" + strconv.Itoa(app.port)
	_ = fasthttp.ListenAndServe(strPort, app.router.Handle)
}

func (app *Application) Test() {

}
