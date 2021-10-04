package loaded

import "caged/http"

type LoadedController struct {
	name    string
	methods map[string]func(res http.Response, req http.Request)
}

func CreateController() *LoadedController {
	controller := new(LoadedController)
	controller.methods = make(map[string]func(res http.Response, req http.Request))
	return controller
}
