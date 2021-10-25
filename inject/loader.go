package inject

import (
	"caged/base"
	"caged/loaded"
	"caged/routing"
)

func LoadModule(module *base.Module) *loaded.LoadedModule {
	loadedModule := loaded.CreateModule(module)
	loadedModule.Load()
	return loadedModule
}
func LoadRouter(module *loaded.LoadedModule) *routing.Router {
	return routing.CreateRouter(module)
}
