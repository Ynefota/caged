package main

import (
	"caged/application"
	"caged/base"
	"caged/test"
)

func main() {
	module := base.Module{Controllers: []base.Controller{test.HelloWorldController{}}}
	application.Create(&module).Listen(4000)

}
