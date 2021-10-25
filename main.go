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
			reflect.TypeOf(test.HelloWorldController{}),
		},
		Providers: []reflect.Type{
			reflect.TypeOf(test.Dep{}),
		},
	}
	app := application.Create(&module)

	pApp := app
	pApp.Listen(8000)

	/*t := reflect.TypeOf(test.Dep{})
	elm := reflect.New(t)
	fmt.Println(elm)
	elm.MethodByName("Init").Call(nil)*/
}
