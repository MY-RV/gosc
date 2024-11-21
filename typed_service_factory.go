package gosc

import "reflect"

type TypedServiceFactory[T any] struct {
	factoryFunction ServiceFactoryFunction[T]
	definitionType  reflect.Type
}

func (f *TypedServiceFactory[T]) createService(s *ServiceCollection) *T {
	return f.factoryFunction(s)
}

func (f *TypedServiceFactory[T]) getDefinitionType() reflect.Type {
	return f.definitionType
}
