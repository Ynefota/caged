package routing

import "github.com/valyala/fasthttp"

type Route interface {
	DoRouting(ctx *fasthttp.RequestCtx) bool
}
