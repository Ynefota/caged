package base

type Module struct {
	Controllers []Controller
	Providers   []Injectable
	Imports     []Module
	Exports     []Injectable
}
