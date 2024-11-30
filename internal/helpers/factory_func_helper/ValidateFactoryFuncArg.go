package factory_func_helper

import (
	"fmt"
	"reflect"
)

func ValidateFactoryFuncArg(factoryFunction any) (reflect.Type, error) {
	factoryFuncType := reflect.TypeOf(factoryFunction)

	if factoryFuncType.Kind() != reflect.Func {
		return nil, fmt.Errorf("value of type %v is not a function", factoryFuncType)
	}

	if factoryFuncType.NumOut() != 1 {
		return nil, fmt.Errorf("function %v has %d returns, expected => 1", factoryFuncType, factoryFuncType.NumOut())
	}

	return factoryFuncType, nil
}