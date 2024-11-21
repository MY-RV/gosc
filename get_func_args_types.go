package gosc

import "reflect"

func getFuncArgsTypes(factoryFuncType reflect.Type) ([]reflect.Type, error) {
	argTypes := make([]reflect.Type, factoryFuncType.NumIn())
	for i := 0; i < factoryFuncType.NumIn(); i++ {
		argTypes[i] = factoryFuncType.In(i)
	}

	return argTypes, nil
}