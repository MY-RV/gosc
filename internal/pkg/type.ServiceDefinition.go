package pkg

import "github.com/MY-RV/gosc/enums/scope"

type ServiceDefinition[T any] struct {
	factory ServiceFactory[T]
	scope   scope.Enum
}
