package loaded

import (
	"caged/base"
	"reflect"
)

type LoadedInjectable struct {
	Injectable base.Injectable
}

func CreateInjectable(t reflect.Type) LoadedInjectable {
	injectable := LoadedInjectable{}
	return injectable
}
