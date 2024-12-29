package dependencyRecord

import "github.com/MY-RV/gosc/enums/scope"

// GetScope implements definitions.DependencyRecord.
func (d *dependencyRecord) GetScope() scope.Enum {
	return d.scope
}
