package pkg

func TypedFactoryFunc[T any](factoryFunction any) *ServiceFactory[T] {
	typedFactoryFunc, err := NewTypedFactoryFunc[T](factoryFunction)
	if err != nil {
		panic(err)
	}

	return typedFactoryFunc
}
