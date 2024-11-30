package pointer_helper

import "reflect"

func TypeGoElem(t reflect.Type) reflect.Type {
	if t.Kind() == reflect.Ptr {
		return t.Elem()
	}

	return t
}
