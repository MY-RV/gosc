package typedFactory

import "reflect"

// Implements github.com/MY-RV/gosc/definitions::TypedFactory.GetFactoryType
func (typedFactory *DependencyTypedFactory[T]) GetFactoryType() reflect.Type {
	return typedFactory.factoryType
}
