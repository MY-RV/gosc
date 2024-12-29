package dependencyCollection

import (
	"github.com/MY-RV/gosc/definitions"
	"github.com/MY-RV/gosc/enums/scope"
	"github.com/MY-RV/gosc/internal/defaults/dependencyRecord"
)

// AddDependency implements definitions.DependencyCollectionBuilder.
func (d *dependencyCollectionBuilder) AddDependency(scope scope.Enum, factory definitions.Factory) {
	dependencyRecord := dependencyRecord.Default(factory, scope)
	registry := d.dependencyCollection.registry
	registry[factory.GetFactoryType()] = dependencyRecord
}