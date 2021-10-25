package loaded

import (
	"caged/base"
	"reflect"
)

type LoadedModule struct {
	Module            *base.Module
	Controllers       []*LoadedController
	Providers         []*LoadedInjectable
	Imports           []*LoadedModule // TODO find a clean way to import other Modules
	currentInjectable reflect.Type
}

func CreateModule(module *base.Module) *LoadedModule {
	loadedModule := &LoadedModule{}
	loadedModule.Module = module
	return loadedModule
}

func (module *LoadedModule) GetController(name string) *LoadedController {
	for i := 0; i < len(module.Controllers); i++ {
		if module.Controllers[i].name == name {
			return module.Controllers[i]
		}
	}
	for i := 0; i < len(module.Imports); i++ {
		controller := module.Imports[i].GetController(name)
		if controller != nil {
			return controller
		}
	}
	return nil
}

func (module *LoadedModule) Load() {
	for i := 0; i < len(module.Module.Imports); i++ {
		newLoadedModule := module.LoadModule(module.Module.Imports[i])
		module.Imports = append(module.Imports, newLoadedModule)
	}
	for i := 0; i < len(module.Module.Providers); i++ {
		provider := module.Module.Providers[i]
		injectable, _ := module.LoadInjectable(provider)
		module.Providers = append(module.Providers, injectable)
	}
	for i := 0; i < len(module.Module.Controllers); i++ {
		controller := module.LoadController(module.Module.Controllers[i])
		module.Controllers = append(module.Controllers, controller)
	}
	for i := 0; i < len(module.Providers); i++ { // autowire injectables
		loadedProvider := module.Providers[i]
		loadedProvider.AutoWire(module)
	}
	for i := 0; i < len(module.Controllers); i++ { // autowire controllers
		loadedController := module.Controllers[i]
		loadedController.AutoWire(module)
	}
	for i := 0; i < len(module.Providers); i++ { // call after wire on injectables
		loadedProvider := module.Providers[i]
		loadedProvider.AfterWire()
	}
	for i := 0; i < len(module.Controllers); i++ { // call after wire on controllers
		loadedController := module.Controllers[i]
		loadedController.AfterWire()
	}
	// TODO load first all dependencies
}

func (module *LoadedModule) LoadModule(moduleType reflect.Type) *LoadedModule { // TODO what to with that thing
	// TODO implementation of loading other modules
	newModule := new(base.Module)
	newLoadedModule := CreateModule(newModule)
	return newLoadedModule
}

func (module *LoadedModule) GetInjectable(t reflect.Type) *reflect.Value {
	for i := 0; i < len(module.Imports); i++ {
		fromModule := module.Imports[i].GetInjectable(t)
		if fromModule != nil {
			return fromModule
		}
	}
	for i := 0; i < len(module.Providers); i++ {
		if module.Providers[i].Injectable.Type() == t {
			return module.Providers[i].Injectable
		}
	}
	return nil
}

func (module *LoadedModule) LoadInjectable(t reflect.Type) (*LoadedInjectable, error) { // TODO concentrate on everything because its little complex
	loadedInjectable := CreateInjectable(t)
	loadedInjectable.Init()
	return loadedInjectable, nil
}

func (module *LoadedModule) LoadController(classType reflect.Type) *LoadedController {
	loadedController := CreateController(classType)
	return loadedController
}

func (module *LoadedModule) IsLoaded(t reflect.Type) bool {
	for i := 0; i < len(module.Imports); i++ {
		if module.Imports[i].IsLoaded(t) {
			return true
		}
	}
	for i := 0; i < len(module.Providers); i++ {
		if module.Providers[i].Injectable.Type() == t {
			return true
		}
	}
	return false
}
