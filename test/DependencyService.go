package test

import (
	"caged/base"
)

type Dep struct {
	base.Injectable
	Dep  *Dep `autowired:""`
	name string
}

func (d *Dep) Init() {
	d.name = "Xy"
}
