package dependencyCollection

import (
	"reflect"

	"github.com/MY-RV/gosc/definitions"
	"github.com/MY-RV/gosc/enums/scopeStatus"
)

func NewDefaultDependencyBuilder() definitions.DependencyCollectionBuilder {
	return &DefaultBuilder{
		defaultDependencyCollection: DefaultDependencyCollection{
			registrations: make(map[reflect.Type]definitions.DependencyRegistry),
			scopeStatus:   scopeStatus.INACTIVE,
		},
	}
}
