package dependencyCollection

import (
	"github.com/MY-RV/gosc/definitions"
	"github.com/MY-RV/gosc/enums/scope"
)

func (defaultBuilder *DefaultBuilder) AddScoped(factory definitions.Factory) {
	defaultBuilder.AddRegistry(scope.SCOPED, factory)
}
