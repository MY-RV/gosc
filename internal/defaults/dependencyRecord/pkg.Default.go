package dependencyRecord

import (
	"github.com/MY-RV/gosc/definitions"
	"github.com/MY-RV/gosc/enums/scope"
)

func Default(factory definitions.Factory, scope scope.Enum) definitions.DependencyRecord {
	return &dependencyRecord{
		factory: factory,
		scope:   scope,
	}
}