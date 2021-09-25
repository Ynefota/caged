package main

import (
	"caged/application"
	"caged/base"
	"caged/test"
)

func main() {
	module := base.Module{Controllers: []base.Controller{test.HelloWorldController{}}}
	app := application.Create(&module)

	pApp := &app
	pApp.Listen(3000)
}
