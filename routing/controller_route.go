package routing

import (
	"github.com/valyala/fasthttp"
)

type ControllerRoute struct {
	Route
}

func (r *ControllerRoute) DoRouting(ctx *fasthttp.RequestCtx) bool {
	return false
}
