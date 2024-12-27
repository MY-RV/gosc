package dependencyCollection

import "github.com/MY-RV/gosc/enums/scope"


// Implements github.com/MY-RV/gosc/definitions::DependencyRegistration.GetFactory
func (factoryRegistration *FactoryRegistry) GetScope() scope.Enum {
	return factoryRegistration.scope
}