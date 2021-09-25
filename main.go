package main

import (
	"caged/application"
	"caged/base"
	"caged/test"
	"reflect"
)

func main() {
	module := base.Module{
		Controllers: []reflect.Type{
			reflect.TypeOf(&test.HelloWorldController{}),
		},
	}
	app := application.Create(&module)

	pApp := &app
	pApp.Listen(3000)
}
