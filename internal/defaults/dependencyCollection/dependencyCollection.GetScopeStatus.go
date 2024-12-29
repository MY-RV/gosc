package dependencyCollection

import "github.com/MY-RV/gosc/enums/scopeStatus"

// GetScopeStatus implements definitions.DependencyCollection.
func (d *dependencyCollection) GetScopeStatus() scopeStatus.Enum {
	return d.scopeStatus
}
