package gosc

func FactoryFunc(factoryFunction any) ServiceFactory[any] {
	result, err := NewFactoryFunc(factoryFunction)
	if err != nil {
		panic(err)
	}

	return result
}

func NewFactoryFunc(factoryFunction any) (ServiceFactory[any], error) {
	return NewTypedFactoryFunc[any](factoryFunction)
}
