package pkg

import (
	"reflect"

	"github.com/MY-RV/gosc/enums/scopeStatus"
)

func NewServiceCollectionBuilder() *ServiceCollectionBuilder {
	return &ServiceCollectionBuilder{
		serviceCollection: ServiceCollection{
			activeServices: ServiceCollectionActiveServices{
				singleton: make(map[reflect.Type]*any),
				scoped:    make(map[reflect.Type]*any),
			},

			servicesDefinitions: make(map[reflect.Type]*ServiceDefinition[any]),
			scopeStatus:         scopeStatus.INACTIVE,
		},
	}
}
