package inject

import (
	"caged/base"
	"caged/loaded"
)

func LoadModule(module *base.Module) *loaded.LoadedModule {
	loadedModule := new(loaded.LoadedModule)
	loadedModule.Module = module
	loadedModule.Load()
	return loadedModule
}
