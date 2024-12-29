package dependencyCollection

import "reflect"

// MustGetDependency implements definitions.DependencyCollection.
func (d *dependencyCollection) MustGetDependency(reflect.Type) any {
	panic("unimplemented")
}
