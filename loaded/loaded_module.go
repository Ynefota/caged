package loaded

import (
	"caged/base"
	"caged/ctx"
	"errors"
	"fmt"
	"reflect"
	"strings"
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
	var loadedInjectable *LoadedInjectable
	if module.currentInjectable == nil {
		module.currentInjectable = t
	} else if module.currentInjectable == t {
		return nil, errors.New("makes loop") // TODO error handling
	}
	if !module.CanLoad(t) {
		return nil, errors.New("not found in module: " + t.Name())
	}
	if module.IsLoaded(t) {
		return module.GetInjectable(t), nil
	} else {
		loadedInjectable = CreateInjectable(t)
		injectableType := reflect.TypeOf(loadedInjectable.Injectable)
		newMethod, _ := injectableType.MethodByName("New")
		paramsAmount := newMethod.Type.NumIn()
		params := make([]reflect.Value, paramsAmount)
		for p := 0; p < paramsAmount; p++ {
			// TODO implementation
			paramType := newMethod.Type.In(p)
			loadedParam, err := module.LoadInjectable(paramType)
			if err != nil {
				return loadedParam, err
			}
			param := loadedParam.Injectable
			params[p] = reflect.ValueOf(param)
		}
	}

	if module.currentInjectable == t {
		module.currentInjectable = nil
	}
	return loadedInjectable, nil
}

func (module *LoadedModule) CanLoad(t reflect.Type) bool {
	return true // TODO implement
}

func (module *LoadedModule) LoadController(classType reflect.Type) *LoadedController {
	loadedController := CreateController()
	className := classType.Name()
	controllerValue := reflect.New(classType).Elem()
	for m := 0; m < controllerValue.NumMethod(); m++ {
		method := controllerValue.Method(m)
		paramsAmount := method.Type().NumIn()
		params := make([]reflect.Value, paramsAmount)
		ctxPosition := -1
		for p := 0; p < paramsAmount; p++ {
			paramType := method.Type().In(p)
			if paramType == classType {
				continue
			} else if paramType == reflect.TypeOf((*ctx.Context)(nil)).Elem() {
				ctxPosition = p
			} else if paramType.Implements(reflect.TypeOf((*base.Injectable)(nil)).Elem()) {
				param := module.GetInjectable(paramType)
				var paramValue reflect.Value
				if param == nil {
					paramValue = reflect.ValueOf(nil) // TODO cleanup & fix
				} else {
					paramValue = reflect.ValueOf(param.Injectable)
				}
				params[p] = paramValue
			}
		}
		loadedController.methods[classType.Method(m).Name] = func(ctx ctx.Context) {
			if ctxPosition != -1 {
				params[ctxPosition] = reflect.ValueOf(ctx)
			}
			// get return value of controller
			fmt.Println(method.Call(params)) // TODO think about return value
		}
		loadedController.methods[classType.Method(m).Name](ctx.Context{})
	}
	classNameLower := strings.ToLower(className)
	if strings.HasSuffix(classNameLower, "controller") {
		className = className[0:strings.LastIndex(classNameLower, "controller")]
	}
	return loadedController
}
