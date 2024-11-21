package gosc

import (
	"reflect"
)

func TypedFactoryFunc[T any](factoryFunction any) ServiceFactory[T] {
	result, err := NewTypedFactoryFunc[T](factoryFunction)
	if err != nil {
		panic(err)
	}

	return result
}

func NewTypedFactoryFunc[T any](factoryFunction any) (ServiceFactory[T], error) {
	factoryFuncType, err := validateNewFactoryFuncArg(factoryFunction)
	if err != nil {
		return nil, err
	}

	argTypes, err := getFuncArgsTypes(factoryFuncType)
	if err != nil {
		return nil, err
	}

	factoryFuncValue := reflect.ValueOf(factoryFunction)

	factory := &TypedServiceFactory[T]{
		definitionType: goPointerIfNonPointer(factoryFuncType.Out(0)),
		factoryFunction: func(serviceCollection *ServiceCollection) *T {
			args, _ := resolveFactoryDependencies(serviceCollection, argTypes)
			callResults := factoryFuncValue.Call(args)
			result := callResults[0]

			if result.Kind() != reflect.Ptr {
				value := result.Interface().(T)
				return &value
			}

			return result.Interface().(*T)
		},
	}

	return factory, nil
}