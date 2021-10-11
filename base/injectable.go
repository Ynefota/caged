package base

type Injectable interface {
	Init()
	AfterWire()
}
