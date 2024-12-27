package typedFactory

import (
	"reflect"

	"github.com/MY-RV/gosc/definitions"
)

func NewDependencyTypedFactory[T any](factoryFunction func(definitions.DependencyCollection) T) DependencyTypedFactory[T] {
	reflectType := reflect.TypeOf(factoryFunction)
	dependencyType := reflectType.Out(0)

	return DependencyTypedFactory[T]{
		factoryFunction: factoryFunction,
		factoryType:     dependencyType,
	}
}
