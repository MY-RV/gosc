package pointer_helper

import "reflect"

func TypeGoPointerIfNotPointer(t reflect.Type) reflect.Type {
	if t.Kind() == reflect.Ptr {
		return t
	}

	return reflect.PointerTo(t)
}

