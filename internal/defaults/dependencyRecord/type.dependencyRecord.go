package dependencyRecord

import (
	"github.com/MY-RV/gosc/definitions"
	"github.com/MY-RV/gosc/enums/scope"
)

type dependencyRecord struct {
	factory definitions.Factory
	scope   scope.Enum
}

// TYPE CHECKING
func init() {
	var _ definitions.DependencyRecord = &dependencyRecord{}
}
