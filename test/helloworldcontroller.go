package test

import (
	"caged/ctx"
)

type HelloWorldController struct {
	dep *Dep
}

func (controller *HelloWorldController) Name(ctx ctx.Context) string {
	return "name " + controller.dep.name
}
