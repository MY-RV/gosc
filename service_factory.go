package gosc

import "reflect"

type ServiceFactory[T any] interface {
	createService(s *ServiceCollection) *T
	getDefinitionType() reflect.Type
}
