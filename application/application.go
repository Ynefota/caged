package application

import (
	"caged/base"
	"caged/inject"
	"caged/loaded"
	"net/http"
	"strconv"
)

type Application struct {
	port         int
	module       *base.Module
	loadedModule *loaded.LoadedModule
}

func Create(module *base.Module) *Application {
	app := new(Application)
	app.module = module
	app.loadedModule = inject.LoadModule(module)
	return app
}

func (app *Application) Listen(port int) {
	app.port = port
	strPort := ":" + strconv.Itoa(app.port)
	http.ListenAndServe(strPort, nil)
}
