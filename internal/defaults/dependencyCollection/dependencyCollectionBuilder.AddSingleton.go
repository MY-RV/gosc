package dependencyCollection

import (
	"github.com/MY-RV/gosc/definitions"
	"github.com/MY-RV/gosc/enums/scope"
)

// AddSingleton implements definitions.DependencyCollectionBuilder.
func (d *dependencyCollectionBuilder) AddSingleton(factory definitions.Factory) {
	d.AddDependency(scope.SCOPED, factory)
}
