package loaded

import (
	"caged/base"
	"fmt"
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
	for i := 0; i < len(module.Providers); i++ {
		loadedProvider := module.Providers[i]
		loadedProvider.AutoWire(module)
	}
	// TODO load first all dependencies
	for i := 0; i < len(module.Module.Controllers); i++ {
		controller := module.LoadController(module.Module.Controllers[i])
		module.Controllers = append(module.Controllers, controller)
	}
}

func (module *LoadedModule) LoadModule(moduleType reflect.Type) *LoadedModule { // TODO what to with that thing
	// TODO implementation of loading other modules
	newModule := new(base.Module)
	newLoadedModule := CreateModule(newModule)
	return newLoadedModule
}

func (module *LoadedModule) GetInjectable(t reflect.Type) *LoadedInjectable {
	if module.IsLoaded(t) {
		fmt.Print("is loaded: ")
		// injectable := module.LoadInjectable(t)
		// TODO get dependency from module
		// TODO check loading cycle and load dependent dependencies
	} else {
		fmt.Print("isn't loaded: ")
	}
	fmt.Println(t)
	// TODO just get dependency from module
	return nil
}

func (module *LoadedModule) IsLoaded(t reflect.Type) bool {
	for i := 0; i < len(module.Imports); i++ {
		if module.Imports[i].IsLoaded(t) {
			return true
		}
	}
	for i := 0; i < len(module.Providers); i++ {
		if reflect.TypeOf(module.Providers[i].Injectable) == t {
			return true
		}
	}
	return false
}

func (module *LoadedModule) LoadInjectable(t reflect.Type) (*LoadedInjectable, error) { // TODO concentrate on everything because its little complex
	loadedInjectable := CreateInjectable(t)
	loadedInjectable.Init()
	return loadedInjectable, nil
}

func (module *LoadedModule) CanLoad(t reflect.Type) bool { // check if type is in module
	return true // TODO implement
}

func (module *LoadedModule) LoadController(classType reflect.Type) *LoadedController {
	loadedController := CreateController(classType)
	loadedController.Load(module)
	return loadedController
}
