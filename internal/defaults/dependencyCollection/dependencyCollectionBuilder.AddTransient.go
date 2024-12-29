package dependencyCollection

import (
	"github.com/MY-RV/gosc/definitions"
	"github.com/MY-RV/gosc/enums/scope"
)

// AddTransient implements definitions.DependencyCollectionBuilder.
func (d *dependencyCollectionBuilder) AddTransient(factory definitions.Factory) {
	d.AddDependency(scope.SCOPED, factory)
}
