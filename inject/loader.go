package inject

import (
	"caged/base"
	"caged/loaded"
)

func LoadModule(module *base.Module) *loaded.LoadedModule {
	loadedModule := loaded.CreateModule(module)
	loadedModule.Load()
	return loadedModule
}
