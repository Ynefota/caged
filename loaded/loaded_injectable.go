package loaded

type LoadedInjectable struct {
}

func CreateInjectable() *LoadedInjectable {
	return new(LoadedInjectable)
}
