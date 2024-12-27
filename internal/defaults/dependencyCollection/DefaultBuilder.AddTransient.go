package dependencyCollection

import (
	"github.com/MY-RV/gosc/definitions"
	"github.com/MY-RV/gosc/enums/scope"
)

func (defaultBuilder *DefaultBuilder) AddTransient(factory definitions.Factory) {
	defaultBuilder.AddRegistry(scope.TRANSIENT, factory)
}
