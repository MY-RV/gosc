package gosc

import "reflect"

type ServiceDefinition struct {
	serviceFactory ServiceFactory[any]
	serviceScope   int
	serviceType    reflect.Type
}
