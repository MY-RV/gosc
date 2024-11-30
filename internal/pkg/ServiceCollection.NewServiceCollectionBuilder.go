package pkg

import (
	"reflect"

	"github.com/MY-RV/gosc/enums/scope_status"
)

func NewServiceCollectionBuilder() *ServiceCollectionBuilder {
	return &ServiceCollectionBuilder{
		serviceCollection: ServiceCollection{
			activeServices: ServiceCollectionActiveServices{
				singleton: make(map[reflect.Type]*any),
				scoped:    make(map[reflect.Type]*any),
			},

			servicesDefinitions: make(map[reflect.Type]*ServiceDefinition[any]),
			scopeStatus:         scope_status.INACTIVE,
		},
	}
}
