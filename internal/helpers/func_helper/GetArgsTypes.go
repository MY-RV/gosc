package func_helper

import "reflect"

func GetArgsTypes(factoryFuncType reflect.Type) []reflect.Type {
	argTypes := make([]reflect.Type, factoryFuncType.NumIn())
	for i := 0; i < factoryFuncType.NumIn(); i++ {
		argTypes[i] = factoryFuncType.In(i)
	}

	return argTypes
}