package test

import (
	"caged/ctx"
	"fmt"
)

type HelloWorldController struct {
	Dep *Dep `autowired:""`
}

func (controller *HelloWorldController) Name(ctx ctx.Context) string {
	return "name " + controller.Dep.name
}

func (controller *HelloWorldController) AfterWire() {
	fmt.Println(controller.Dep.name)
}
