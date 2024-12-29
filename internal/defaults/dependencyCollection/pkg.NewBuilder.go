package dependencyCollection

import (
	"reflect"

	"github.com/MY-RV/gosc/definitions"
	"github.com/MY-RV/gosc/enums/scopeStatus"
)

func NewBuilder() definitions.DependencyCollectionBuilder {
	return &dependencyCollectionBuilder{
		dependencyCollection: dependencyCollection{
			singletonServices: make(map[reflect.Type]any),
			scopedServices:    make(map[reflect.Type]any),
			registry:          make(map[reflect.Type]definitions.DependencyRecord),
			scopeStatus:       scopeStatus.INACTIVE,
		},
	}
}
