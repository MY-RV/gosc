package dependencyCollection

import (
	"github.com/MY-RV/gosc/definitions"
	"github.com/MY-RV/gosc/enums/scopeStatus"
)

// Implements github.com/MY-RV/gosc/definitions::DependencyCollectionBuilder.Build
func (defaultBuilder *DefaultBuilder) Build() definitions.DependencyCollection {
	return &DefaultDependencyCollection{
		registrations: defaultBuilder.defaultDependencyCollection.registrations,
		scopeStatus:   scopeStatus.ACTIVE,
	}
}
