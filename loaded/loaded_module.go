package loaded

import (
	"caged/base"
	"fmt"
	"reflect"
	"strings"
)

type LoadedModule struct {
	Module      *base.Module
	Controllers []*LoadedController
	Providers   []*LoadedInjectable
	Imports     []*LoadedModule
}

func CreateModule(module *base.Module) *LoadedModule {
	loadedModule := &LoadedModule{}
	loadedModule.Module = module
	return loadedModule
}

func (module *LoadedModule) Load() {
	for i := 0; i < len(module.Module.Imports); i++ {
		newLoadedModule := module.LoadModule(module.Module.Imports[i])
		module.Imports = append(module.Imports, newLoadedModule)
	}
	// TODO load first all dependencies
	for i := 0; i < len(module.Module.Controllers); i++ {
		controller := module.LoadController(module.Module.Controllers[i])
		module.Controllers = append(module.Controllers, controller)
	}
}
func (module *LoadedModule) LoadModule(moduleType reflect.Type) *LoadedModule {
	// TODO implementation of loading other modules
	newModule := new(base.Module)
	newLoadedModule := CreateModule(newModule)
	return newLoadedModule
}

func (module *LoadedModule) LoadDependency(t reflect.Type) {
	if t.Implements(reflect.TypeOf((*base.Injectable)(nil)).Elem()) {
		if !module.IsLoaded(t) {
			fmt.Println("is children of injectable")
			// injectable := module.LoadInjectable(t)
			// TODO load from module
			// TODO check loading cycle and load dependent dependencies
		} else {
			fmt.Println("isn't a injectable")
		}
		fmt.Println(t)
	}
	// TODO load dependency of t and add to module (also return)
}

func (module *LoadedModule) IsLoaded(t reflect.Type) bool {
	for i := 0; i < len(module.Imports); i++ {
	}
	for i := 0; i < len(module.Providers); i++ {
		if reflect.TypeOf(module.Providers[i].Injectable) == t {
			return true
		}
	}
	return false
}

func (module *LoadedModule) LoadController(classType reflect.Type) *LoadedController {
	loadedController := CreateController()
	className := classType.Name()
	controllerValue := reflect.New(classType)
	for i := 0; i < classType.NumMethod(); i++ {
		method := classType.Method(i)
		params := make([]reflect.Value, method.Type.NumIn())
		for j := 0; j < method.Type.NumIn(); j++ {
			paramType := method.Type.In(j)
			if paramType == classType {
				params[j] = controllerValue
			} else {
				module.LoadDependency(paramType)
			}
		}
		loadedController.methods[method.Name] = func() {
			fmt.Println(method.Func.Call(params))
		}
	}
	classNameLower := strings.ToLower(className)
	if strings.HasSuffix(classNameLower, "controller") {
		className = className[0:strings.LastIndex(classNameLower, "controller")]
	}
	return loadedController
}
func (module *LoadedModule) LoadInjectable(t reflect.Type) *LoadedInjectable {
	loadedInjectable := new(LoadedInjectable)
	return loadedInjectable
}
