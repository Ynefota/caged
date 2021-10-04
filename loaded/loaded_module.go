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
	Imports     []*LoadedModule // TODO find a clean way to import other Modules
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

func (module *LoadedModule) GetInjectable(t reflect.Type) *base.Injectable {
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
	for m := 0; m < classType.NumMethod(); m++ {
		method := classType.Method(m)
		paramsAmount := method.Type.NumIn()
		params := make([]reflect.Value, method.Type.NumIn())
		fmt.Println(reflect.TypeOf(params))
		resPosition := -1
		reqPosition := -1
		for p := 0; p < paramsAmount; p++ {
			paramType := method.Type.In(p)
			if paramType == classType {
				params[p] = controllerValue
			} else if paramType == reflect.TypeOf((*http.Response)(nil)).Elem() {
				resPosition = p
			} else if paramType == reflect.TypeOf((*http.Request)(nil)).Elem() {
				reqPosition = p
			} else if paramType.Implements(reflect.TypeOf((*base.Injectable)(nil)).Elem()) {
				params[p] = reflect.ValueOf(module.GetInjectable(paramType))
			}
		}
		loadedController.methods[method.Name] = func(res http.Response, req http.Request) {
			if reqPosition != -1 {
				params[reqPosition] = reflect.ValueOf(req)
			}
			if resPosition != -1 {
				params[resPosition] = reflect.ValueOf(res)
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
