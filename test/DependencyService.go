package test

import (
	"caged/base"
)

type Dep struct {
	base.Injectable
	name string
}

func (d *Dep) New() {
	d.name = "Xy"
}
