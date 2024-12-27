package dependencyCollection

import "github.com/MY-RV/gosc/definitions"

// Implements github.com/MY-RV/gosc/definitions::DependencyRegistration.GetFactory
func (factoryRegistration *FactoryRegistry) GetFactory() definitions.Factory {
	return factoryRegistration.factory
}