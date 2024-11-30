package pkg

import "reflect"

type ServiceFactory[T any] struct {
	createService   func(*ServiceCollection) (T, error)
	serviceType     reflect.Type
	dependencyTypes []reflect.Type
}

func (serviceFactory *ServiceFactory[T]) CreateService(serviceCollection *ServiceCollection) (T, error) {
	return serviceFactory.createService(serviceCollection)
}

func (serviceFactory *ServiceFactory[T]) GetServiceType() reflect.Type {
	return serviceFactory.serviceType
}

func (serviceFactory *ServiceFactory[T]) GetDependencyTypes() []reflect.Type {
	return append([]reflect.Type{}, serviceFactory.dependencyTypes...)
}
