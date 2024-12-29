package dependencyCollection

import "reflect"

// GetDependency implements definitions.DependencyCollection.
func (d *dependencyCollection) GetDependency(reflect.Type) (any, error) {
	panic("unimplemented")
}
