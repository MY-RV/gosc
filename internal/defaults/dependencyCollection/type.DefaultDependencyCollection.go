package dependencyCollection

import (
	"reflect"

	"github.com/MY-RV/gosc/definitions"
	"github.com/MY-RV/gosc/enums/scopeStatus"
)

// Implements github.com/MY-RV/gosc/definitions::DependencyCollection
type DefaultDependencyCollection struct {
	// singletonServices map[reflect.Type]any
	// scopedServices    map[reflect.Type]any

	registrations map[reflect.Type]definitions.DependencyRegistry
	scopeStatus   scopeStatus.Enum
}



func (ddc *DefaultDependencyCollection) GetDependencyFactory(reflect.Type) (definitions.Factory, error) {
	panic("Not Implemented!")
}

func (ddc *DefaultDependencyCollection) GetRegisteredDependencies() []reflect.Type {
	panic("Not Implemented!")
}

func (ddc *DefaultDependencyCollection) GetScopeStatus() scopeStatus.Enum {
	panic("Not Implemented!")
}

func (ddc *DefaultDependencyCollection) GetDependency(reflect.Type) (any, error) {
	panic("Not Implemented!")
}

func (ddc *DefaultDependencyCollection) MustGetDependency(reflect.Type) any {
	panic("Not Implemented!")
}