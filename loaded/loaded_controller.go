package loaded

import (
	"caged/base"
	"caged/ctx"
	"fmt"
	"reflect"
	"strings"
)

type LoadedController struct {
	name           string
	methods        map[string]func(ctx ctx.Context)
	controllerType reflect.Type
}

func CreateController(t reflect.Type) *LoadedController {
	controller := new(LoadedController)
	controller.controllerType = t
	controller.methods = make(map[string]func(ctx ctx.Context))
	return controller
}
func (controller *LoadedController) Load(module *LoadedModule) {
	controllerName := controller.controllerType.Name()
	controllerValue := reflect.New(controller.controllerType).Elem()
	for m := 0; m < controllerValue.NumMethod(); m++ {
		method := controllerValue.Method(m)
		paramsAmount := method.Type().NumIn()
		params := make([]reflect.Value, paramsAmount)
		ctxPosition := -1
		for p := 0; p < paramsAmount; p++ {
			paramType := method.Type().In(p)
			if paramType == controller.controllerType {
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
		// TODO is controller.controllerType.Method(m).Name everytime the name of method
		controller.methods[controller.controllerType.Method(m).Name] = func(ctx ctx.Context) {
			if ctxPosition != -1 {
				params[ctxPosition] = reflect.ValueOf(ctx)
			}
			// get return value of controller
			fmt.Println(method.Call(params)) // TODO think about return value
		}
		controller.methods[controller.controllerType.Method(m).Name](ctx.Context{})
	}
	controllerName = strings.ToLower(controllerName)
	if strings.HasSuffix(controllerName, "controller") {
		controllerName = controllerName[0:strings.LastIndex(controllerName, "controller")]
	}
	fmt.Println(controllerName)
	controller.name = controllerName
}
