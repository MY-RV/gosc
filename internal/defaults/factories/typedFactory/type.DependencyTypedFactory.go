package typedFactory

import (
	"reflect"

	"github.com/MY-RV/gosc/definitions"
)

// Implements github.com/MY-RV/gosc/definitions::DependencyTypedFactory
type DependencyTypedFactory[T any] struct {
	factoryFunction func(definitions.DependencyCollection) T
	factoryType     reflect.Type
}
