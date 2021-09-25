package test

type HelloWorldController struct {
}

func (controller HelloWorldController) Name(dep *Dep) string {
	return "name " + dep.name
}
