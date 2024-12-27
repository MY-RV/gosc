package typedFactory

import "github.com/MY-RV/gosc/definitions"

// Implements github.com/MY-RV/gosc/definitions::TypedFactory.TypedCreate
func (typedFactory *DependencyTypedFactory[T]) TypedCreate(dependencyCollection definitions.DependencyCollection) T {
	return typedFactory.factoryFunction(dependencyCollection)
}
