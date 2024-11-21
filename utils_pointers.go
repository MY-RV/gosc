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

func ensureElem(t reflect.Type) reflect.Type {
	for t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	return t
}
