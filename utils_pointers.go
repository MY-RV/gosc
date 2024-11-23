package gosc

import "reflect"

func goPointerIfNonPointer(t reflect.Type) reflect.Type {
	if t.Kind() == reflect.Ptr {
		return t
	}

	return reflect.PointerTo(t)
}

func goElemIfNonElem(t reflect.Type) reflect.Type {
	if t.Kind() == reflect.Ptr {
		return t.Elem()
	}

	return t
}

/* func ensureElem(t reflect.Type) reflect.Type {
	for t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	return t
} */

func goValuePointerIfElem(value reflect.Value) reflect.Value {
	if value.Kind() == reflect.Ptr {
		return value
	}

	ptr := reflect.New(value.Type())
	ptr.Elem().Set(value)
	return ptr
}

func goValueElemIfPointer(value reflect.Value) reflect.Value {
	for value.Kind() == reflect.Ptr || value.Kind() == reflect.Interface {
		value = value.Elem()
	}

	return value
}
