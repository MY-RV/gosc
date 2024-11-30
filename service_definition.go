package gosc

import (
	"reflect"

	"github.com/MY-RV/gosc/scope"
)

type ServiceDefinition struct {
	serviceFactory ServiceFactory[any]
	serviceScope   scope.Enum
	serviceType    reflect.Type
}
