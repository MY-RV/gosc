package typedFactory

import "github.com/MY-RV/gosc/definitions"

// Implements github.com/MY-RV/gosc/definitions::TypedFactory.Create
func (typedFactory *DependencyTypedFactory[T]) Create(dependencyCollection definitions.DependencyCollection) any {
	return typedFactory.factoryFunction(dependencyCollection)
}
