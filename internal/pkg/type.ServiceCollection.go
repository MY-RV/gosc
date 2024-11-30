package pkg

import (
	"reflect"

	"github.com/MY-RV/gosc/enums/scope_status"
)

type ServiceCollection struct {
	activeServices      ServiceCollectionActiveServices
	servicesDefinitions map[reflect.Type]*ServiceDefinition[any]
	scopeStatus         scope_status.Enum
}
