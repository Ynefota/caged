package test

type HelloWorldController struct {
}

func (controller HelloWorldController) Name(dep *Dep, name string) string {
	return "name " + dep.name
}
