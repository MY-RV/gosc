package pkg

import "github.com/MY-RV/gosc/scope"

type ServiceDefinition[T any] struct {
	factory ServiceFactory[T]
	scope   scope.Enum
}
