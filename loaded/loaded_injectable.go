package loaded

import (
	"fmt"
	"reflect"
)

type LoadedInjectable struct {
	Injectable *reflect.Value
}

func CreateInjectable(t reflect.Type) *LoadedInjectable {
	injectable := new(LoadedInjectable)
	in := reflect.New(t)
	injectable.Injectable = &in
	return injectable
}

func (injectable *LoadedInjectable) Init() {
	initMethod := injectable.Injectable.MethodByName("Init")
	fmt.Println(initMethod)
	initMethod.Call(nil)
}

func (injectable *LoadedInjectable) AutoWire(module *LoadedModule) {
	inj := injectable.Injectable.Type().Elem()
	for i := 0; i < inj.NumField(); i++ {
		field := inj.Field(i)
		_, ok := field.Tag.Lookup("autowired")
		if ok {

		}
		fmt.Println(field)
	}
}
