package loaded

import (
	"caged/base"
	"reflect"
)

type LoadedInjectable struct {
	Injectable *base.Injectable
}

func CreateInjectable(t reflect.Type) *LoadedInjectable {
	injectable := new(LoadedInjectable)
	injectable.Injectable = reflect.New(t).Elem().Interface().(*base.Injectable)
	return injectable
}

func (injectable *LoadedInjectable) Init() {
	injectableValue := reflect.ValueOf(injectable.Injectable)
	initMethod := injectableValue.MethodByName("Init")
	initMethod.Call([]reflect.Value{})
}
func (injectable *LoadedInjectable) AutoWire(module *LoadedModule) {

}
