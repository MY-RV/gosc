package pointer_helper

import "reflect"

func ValueEnsureElem(value reflect.Value) reflect.Value {
	for value.Kind() == reflect.Ptr || value.Kind() == reflect.Interface {
		value = value.Elem()
	}

	return value
}