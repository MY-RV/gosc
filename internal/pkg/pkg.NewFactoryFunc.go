package pkg

func NewFactoryFunc(factoryFunction any) (*ServiceFactory[any], error) {
	return NewTypedFactoryFunc[any](factoryFunction)
}
