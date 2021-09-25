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

func (module *LoadedModule) Load() {
	for i := 0; i < len(module.Module.Controllers); i++ {
		controller := module.LoadController(module.Module.Controllers[i])
		module.Controllers = append(module.Controllers, controller)
	}
}

func (module *LoadedModule) LoadDependency(t reflect.Type) {
	fmt.Println(t)
}

func (module *LoadedModule) LoadController(controller base.Controller) *LoadedController {
	loadedController := CreateController()
	classType := reflect.TypeOf(controller)
	className := classType.Name()
	for i := 0; i < classType.NumMethod(); i++ {
		method := classType.Method(i)
		params := make([]reflect.Value, method.Type.NumIn())
		for j := 0; j < method.Type.NumIn(); j++ {
			paramType := method.Type.In(j)
			if paramType == classType {
				params[j] = reflect.ValueOf(controller)
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
	fmt.Println(className)
	return loadedController
}
func (module *LoadedModule) LoadInjectable(injectable base.Injectable) *LoadedInjectable {
	loadedInjectable := new(LoadedInjectable)
	return loadedInjectable
}
