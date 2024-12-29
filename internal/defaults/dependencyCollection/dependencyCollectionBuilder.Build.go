package dependencyCollection

import (
	"reflect"

	"github.com/MY-RV/gosc/definitions"
	"github.com/MY-RV/gosc/enums/scopeStatus"
	"github.com/MY-RV/gosc/internal/helpers/map_helper"
)

// Build implements definitions.DependencyCollectionBuilder.
func (d *dependencyCollectionBuilder) Build() definitions.DependencyCollection {
	return &dependencyCollection{
		singletonServices: d.dependencyCollection.singletonServices,
		scopedServices:    make(map[reflect.Type]any),
		registry:          map_helper.CopyMap(d.dependencyCollection.registry),
		scopeStatus:       scopeStatus.ACTIVE,
	}
}