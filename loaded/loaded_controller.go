package loaded

type LoadedController struct {
	name    string
	methods map[string]func()
}

func CreateController() *LoadedController {
	controller := new(LoadedController)
	controller.methods = make(map[string]func())
	return controller
}
