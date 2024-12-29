package factories

import "github.com/MY-RV/gosc/definitions"

// Implements github.com/MY-RV/gosc/definitions.TypedFactory
type TypedFactory[T any] struct {
	factoryFunction func(definitions.DependencyCollection) T
}
