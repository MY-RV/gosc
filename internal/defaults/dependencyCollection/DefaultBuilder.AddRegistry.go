package dependencyCollection

import (
	"github.com/MY-RV/gosc/definitions"
	"github.com/MY-RV/gosc/enums/scope"
)

func (defaultBuilder *DefaultBuilder) AddRegistry(scope scope.Enum, factory definitions.Factory) {
	dependencyRegistry := &FactoryRegistry{
		factory: factory,
		scope:   scope,
	}

	registrations := defaultBuilder.defaultDependencyCollection.registrations
	// TODO: use Right RegistrationType
	registrations[factory.GetFactoryType()] = dependencyRegistry
}
