package loaded

import (
	"caged/ctx"
)

type LoadedController struct {
	name    string
	methods map[string]func(ctx ctx.Context)
}

func CreateController() *LoadedController {
	controller := new(LoadedController)
	controller.methods = make(map[string]func(ctx ctx.Context))
	return controller
}
