package definitions

import (
	"github.com/MY-RV/gosc/enums/scope"
)

type DependencyRecord interface {
	GetFactory() Factory
	GetScope() scope.Enum
}
