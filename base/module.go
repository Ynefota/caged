package base

import "reflect"

type Module struct {
	Controllers []reflect.Type
	Providers   []reflect.Type
	Imports     []reflect.Type
	Exports     []reflect.Type
}
