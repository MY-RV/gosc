package pointer_helper

import "reflect"

func ValueGoPointerIfNotPointer(value reflect.Value) reflect.Value {
	if value.Kind() == reflect.Ptr {
		return value
	}

	ptr := reflect.New(value.Type())
	ptr.Elem().Set(value)
	return ptr
}
