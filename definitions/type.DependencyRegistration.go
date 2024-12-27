package definitions

import (
	"github.com/MY-RV/gosc/enums/scope"
)

type DependencyRegistry interface {
	GetFactory() Factory
	GetScope() scope.Enum
}
