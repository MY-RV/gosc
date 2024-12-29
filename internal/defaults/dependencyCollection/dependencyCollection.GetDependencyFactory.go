package dependencyCollection

import (
	"reflect"

	"github.com/MY-RV/gosc/definitions"
)

// GetDependencyFactory implements definitions.DependencyCollection.
func (d *dependencyCollection) GetDependencyFactory(reflect.Type) (definitions.Factory, error) {
	panic("unimplemented")
}