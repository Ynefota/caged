package ctx

import "github.com/valyala/fasthttp"

type Context struct {
	context *fasthttp.RequestCtx
}

func (c *Context) Path() {

}
