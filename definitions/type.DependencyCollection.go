package definitions

import (
	"reflect"

	"github.com/MY-RV/gosc/enums/scopeStatus"
)

type DependencyCollection interface {
	GetDependencyFactory(reflect.Type) (Factory, error)
	GetRegisteredDependencies() []reflect.Type
	GetScopeStatus() scopeStatus.Enum

	GetDependency(reflect.Type) (any, error)
	MustGetDependency(reflect.Type) any
}
