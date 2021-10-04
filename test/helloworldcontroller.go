package test

import "caged/http"

type HelloWorldController struct {
}

func (controller HelloWorldController) Name(dep *Dep, response http.Response) string {
	return "name " + dep.name
}
