package dependencyCollection

import (
	"github.com/MY-RV/gosc/definitions"
	"github.com/MY-RV/gosc/enums/scope"
)

func (defaultBuilder *DefaultBuilder) AddSingleton(factory definitions.Factory) {
	defaultBuilder.AddRegistry(scope.SINGLETON, factory)
}
