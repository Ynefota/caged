package loaded

import (
	"caged/ctx"
	"fmt"
	"reflect"
	"strings"
)

type LoadedController struct {
	name           string
	methods        map[string]func(ctx ctx.Context)
	controllerType reflect.Type
	Controller     *reflect.Value
}

func CreateController(t reflect.Type) *LoadedController {
	controller := new(LoadedController)
	controller.controllerType = t
	controller.methods = make(map[string]func(ctx ctx.Context))
	it := reflect.New(t)
	controller.Controller = &it

	// get controller name
	controllerName := controller.controllerType.Name()
	controllerName = strings.ToLower(controllerName)
	if strings.HasSuffix(controllerName, "controller") {
		controllerName = controllerName[0:strings.LastIndex(controllerName, "controller")]
	}
	controller.name = controllerName

	controller.loadMethods() // load methods

	return controller
}

func (controller *LoadedController) AutoWire(module *LoadedModule) {
	inj := controller.Controller.Type().Elem()
	autowireFieldNames := make([]string, 0)
	for i := 0; i < inj.NumField(); i++ {
		field := inj.Field(i)
		_, ok := field.Tag.Lookup("autowired")
		if ok {
			autowireFieldNames = append(autowireFieldNames, field.Name)
			fmt.Println(field.Name)
		}
	}
	for _, fieldName := range autowireFieldNames {
		fmt.Println("initialize in controller " + fieldName)
		field := controller.Controller.Elem().FieldByName(fieldName)
		field.Set(*module.GetInjectable(field.Type()))
	}
}

func (controller *LoadedController) AfterWire() {
	if object, ok := controller.Controller.Interface().(interface{ AfterWire() }); ok {
		object.AfterWire()
	}
}

func (controller *LoadedController) loadMethods() {
	for m := 0; m < controller.Controller.NumMethod(); m++ {
		method := controller.Controller.Method(m)
		paramsAmount := method.Type().NumIn()
		fmt.Println(paramsAmount)
		// TODO is controller.controllerType.Method(m).Name everytime the name of method
		controller.methods[controller.Controller.Type().Method(m).Name] = func(ctx ctx.Context) {
			var params []reflect.Value
			if paramsAmount == 1 {
				params = make([]reflect.Value, paramsAmount)
				params[0] = reflect.ValueOf(ctx)
			} else {
				params = nil
			}
			// get return value of controller
			fmt.Println(method.Call(params)) // TODO think about return value
		}
	}
}
