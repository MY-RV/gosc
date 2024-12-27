package gosc

import "github.com/MY-RV/gosc/definitions"

/* import "github.com/MY-RV/gosc/internal/pkg"

type ServiceCollectionBuilder pkg.ServiceCollectionBuilder
type ServiceCollection pkg.ServiceCollection
type ServiceFactory[T any] pkg.ServiceFactory[T]

func NewServiceCollectionBuilder() *pkg.ServiceCollectionBuilder {
	return pkg.NewServiceCollectionBuilder()
}

func TypedFactoryFunc[T any](factoryFunction any) *pkg.ServiceFactory[T] {
	return pkg.TypedFactoryFunc[T](factoryFunction)
}

func NewTypedFactoryFunc[T any](factoryFunction any) (*pkg.ServiceFactory[T], error) {
	return pkg.NewTypedFactoryFunc[T](factoryFunction)
}

func FactoryFunc(factoryFunction any) *pkg.ServiceFactory[any] {
	return pkg.FactoryFunc(factoryFunction)
}

func NewFactoryFunc(factoryFunction any) (*pkg.ServiceFactory[any], error) {
	return pkg.NewFactoryFunc(factoryFunction)
} */

func NewDependencyCollectionBuilder() definitions.DependencyCollectionBuilder {
	return nil
}
