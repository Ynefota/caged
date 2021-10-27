package routing

import (
	"github.com/valyala/fasthttp"
	"os"
	"strings"
)

type StaticRoute struct {
	Route
	static     string
	viewEngine *string
}

func (s *StaticRoute) Init() {
	if strings.HasSuffix(s.static, "/") {
		s.static = strings.TrimSuffix(s.static, "/")
	}
}

func (s *StaticRoute) DoRouting(ctx *fasthttp.RequestCtx) bool {
	filename := s.static + string(ctx.Path())
	if s.routePath(ctx, filename) {
		return true
	} else if s.viewEngine != nil {
		return s.alternativePath(ctx, filename)
	}
	return false
}

func (s *StaticRoute) routePath(ctx *fasthttp.RequestCtx, filename string) bool {
	_, err := os.Stat(filename)
	exists := !os.IsNotExist(err)
	if exists {
		ctx.SendFile(filename)
	}
	return exists
}
func (s *StaticRoute) alternativePath(ctx *fasthttp.RequestCtx, filename string) bool {
	if !strings.HasSuffix(filename, *s.viewEngine) {
		filename += *s.viewEngine
		return s.routePath(ctx, filename)
	}
	return false
}
