package dependencyCollection

import (
	"reflect"

	"github.com/MY-RV/gosc/definitions"
	"github.com/MY-RV/gosc/enums/scopeStatus"
)

// Implements github.com/MY-RV/gosc/definitions::DependencyCollection
type dependencyCollection struct {
	singletonServices map[reflect.Type]any
	scopedServices    map[reflect.Type]any

	registry    map[reflect.Type]definitions.DependencyRecord
	scopeStatus scopeStatus.Enum
}

// TYPE CHECKING
func init() {
	var _ definitions.DependencyCollection = &dependencyCollection{}
}
