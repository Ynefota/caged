package routing

import "fmt"

type Route struct {
}

func (r Route) Match(path []byte) bool {
	fmt.Println(path)
	return false
}
