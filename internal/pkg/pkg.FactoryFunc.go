package pkg

func FactoryFunc(factoryFunction any) *ServiceFactory[any] {
	typedFactoryFunc, err := NewTypedFactoryFunc[any](factoryFunction)
	if err != nil {
		panic(err)
	}

	return typedFactoryFunc
}
