package test

import (
	"caged/ctx"
)

type HelloWorldController struct {
}

func (controller *HelloWorldController) Name(dep *Dep, ctx ctx.Context) string {
	return "name " + dep.name
}
