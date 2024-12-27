package definitions

import "reflect"

type Factory interface {
	GetFactoryType() reflect.Type
	Create(DependencyCollection) any
}
