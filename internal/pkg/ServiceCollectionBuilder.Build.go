package pkg

import (
	"reflect"

	"github.com/MY-RV/gosc/enums/scope_status"
)

func (serviceCollectionBuilder *ServiceCollectionBuilder) Build() *ServiceCollection {
	serviceCollection := serviceCollectionBuilder.serviceCollection

	response := ServiceCollection{
		activeServices: ServiceCollectionActiveServices{
			singleton: serviceCollection.activeServices.singleton,
			scoped:    make(map[reflect.Type]*any),
		},

		servicesDefinitions: serviceCollection.servicesDefinitions,
		scopeStatus:         scope_status.INACTIVE,
	}

	return &response
}
