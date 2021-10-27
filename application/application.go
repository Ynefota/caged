package application

// TODO get information on https://github.com/valyala/fasthttp#switching-from-nethttp-to-fasthttp

import (
	"caged/base"
	"caged/inject"
	"caged/loaded"
	"caged/routing"
	"github.com/valyala/fasthttp"
	"log"
	"strconv"
)

type Application struct {
	port         int
	server       *fasthttp.Server
	router       *routing.Router
	module       *base.Module
	loadedModule *loaded.LoadedModule
}

func Create(module *base.Module) *Application {
	app := &Application{}
	app.module = module
	app.loadedModule = inject.LoadModule(module)
	app.router = inject.LoadRouter(app.loadedModule)
	app.server = &fasthttp.Server{
		Handler: app.router.Handle,
	}
	return app
}
func (app *Application) Module() *loaded.LoadedModule {
	return app.loadedModule
}

func (app *Application) Listen(port int) {
	app.port = port
	app.router.Setup()
	strPort := ":" + strconv.Itoa(app.port)
	if err := app.server.ListenAndServe(strPort); err != nil {
		log.Fatalf("error in ListenAndServe: %s", err)
	}
}

func (app *Application) UseStaticAssets(folder string) {
	app.router.UseStaticAssets(folder)
}

func (app *Application) SetViewEngine(engine string) {
	app.router.SetViewEngine(engine)
}
