package dependencyCollection

import (
	"github.com/MY-RV/gosc/definitions"
	"github.com/MY-RV/gosc/enums/scope"
)

// AddScoped implements definitions.DependencyCollectionBuilder.
func (d *dependencyCollectionBuilder) AddScoped(factory definitions.Factory) {
	d.AddDependency(scope.SCOPED, factory)
}
