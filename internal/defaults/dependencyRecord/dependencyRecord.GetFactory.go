package dependencyRecord

import "github.com/MY-RV/gosc/definitions"

// GetFactory implements definitions.DependencyRecord.
func (d *dependencyRecord) GetFactory() definitions.Factory {
	return d.factory
}
