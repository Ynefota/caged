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
	injectable.Injectable.MethodByName("Init").Call(nil)
}

func (injectable *LoadedInjectable) AutoWire(module *LoadedModule) {
	inj := injectable.Injectable.Type().Elem()
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
		fmt.Println("initialize " + fieldName)
		field := injectable.Injectable.Elem().FieldByName(fieldName)
		field.Set(*module.GetInjectable(field.Type()))
	}
}

func (injectable *LoadedInjectable) AfterWire() {
	if object, ok := injectable.Injectable.Interface().(interface{ AfterWire() }); ok {
		object.AfterWire()
	}
}
