package dependencyCollection

import (
	"github.com/MY-RV/gosc/definitions"
	"github.com/MY-RV/gosc/enums/scope"
)

// Implements github.com/MY-RV/gosc/definitions::DependencyRegistration
type FactoryRegistry struct {
	scope 	scope.Enum
	factory definitions.Factory
}