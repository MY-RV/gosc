package pkg

import (
	"reflect"

	"github.com/MY-RV/gosc/enums/scopeStatus"
)

type ServiceCollection struct {
	activeServices      ServiceCollectionActiveServices
	servicesDefinitions map[reflect.Type]*ServiceDefinition[any]
	scopeStatus         scopeStatus.Enum
}
