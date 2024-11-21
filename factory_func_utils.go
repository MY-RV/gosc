package gosc

import (
	"fmt"
	"reflect"
)

func validateNewFactoryFuncArg(factoryFunction any) (reflect.Type, error) {
	factoryFuncType := reflect.TypeOf(factoryFunction)

	if factoryFuncType.Kind() != reflect.Func {
		return nil, fmt.Errorf("value of type %v is not a function", factoryFuncType)
	}

	if factoryFuncType.NumOut() != 1 {
		return nil, fmt.Errorf("function %v has %d returns, expected => 1", factoryFuncType, factoryFuncType.NumOut())
	}

	return factoryFuncType, nil
}

func resolveFactoryDependencies(serviceCollection *ServiceCollection, argTypes []reflect.Type) ([]reflect.Value, error) {
	args := make([]reflect.Value, len(argTypes))
	for i, argType := range argTypes {
		service := serviceCollection.getService(argType)
		args[i] = reflect.ValueOf(service)
	}

	return args, nil
}
