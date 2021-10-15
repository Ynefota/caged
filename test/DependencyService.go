package test

import (
	"caged/base"
	"fmt"
)

type Dep struct {
	base.Injectable
	name string
}

func (d *Dep) Init() {
	d.name = "Xy"
	fmt.Println("hi")
}
