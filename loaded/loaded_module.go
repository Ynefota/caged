package loaded

import (
	"caged/base"
	"caged/http"
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

func (module *LoadedModule) GetDependency(t reflect.Type) *LoadedInjectable {
	if module.IsLoaded(t) {
		fmt.Print("is loaded: ")
		// injectable := module.LoadInjectable(t)
		// TODO load from module
		// TODO check loading cycle and load dependent dependencies
	} else {
		fmt.Print("isn't loaded: ")
	}
	fmt.Println(t)
	// TODO load dependency of t and add to module (also return)
	return nil
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
	for methodCounter := 0; methodCounter < classType.NumMethod(); methodCounter++ {
		method := classType.Method(methodCounter)
		paramsAmount := method.Type.NumIn()
		params := make([]reflect.Value, method.Type.NumIn())
		for paramCounter := 0; paramCounter < paramsAmount; paramCounter++ {
			paramType := method.Type.In(paramCounter)
			if paramType == classType {
				params[paramCounter] = controllerValue
			} else if paramType.Implements(reflect.TypeOf((*http.Response)(nil)).Elem()) {
				params[paramCounter] = reflect.ValueOf((*http.Response)(nil))
			} else if paramType.Implements(reflect.TypeOf((*http.Request)(nil)).Elem()) {
				params[paramCounter] = reflect.ValueOf((*http.Request)(nil))
			} else if paramType.Implements(reflect.TypeOf((*base.Injectable)(nil)).Elem()) {
				params[paramCounter] = reflect.ValueOf(module.GetDependency(paramType))
			}
		}
		loadedController.methods[method.Name] = func(res http.Response, req http.Request) {
			for i := 0; i < paramsAmount; i++ { // TODO inject response and requests

			}
			fmt.Println(method.Func.Call(params)) // TODO think about injecting ResponseWriter and Request
		}
	}
	classNameLower := strings.ToLower(className)
	if strings.HasSuffix(classNameLower, "controller") {
		className = className[0:strings.LastIndex(classNameLower, "controller")]
	}
	return loadedController
}

func (module *LoadedModule) LoadInjectable(t reflect.Type) *LoadedInjectable {
	return nil // TODO implementation
}
