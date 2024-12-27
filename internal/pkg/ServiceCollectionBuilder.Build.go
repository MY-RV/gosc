package pkg

import (
	"reflect"

	"github.com/MY-RV/gosc/enums/scopeStatus"
)

func (serviceCollectionBuilder *ServiceCollectionBuilder) Build() *ServiceCollection {
	serviceCollection := serviceCollectionBuilder.serviceCollection

	response := ServiceCollection{
		activeServices: ServiceCollectionActiveServices{
			singleton: serviceCollection.activeServices.singleton,
			scoped:    make(map[reflect.Type]*any),
		},

		servicesDefinitions: serviceCollection.servicesDefinitions,
		scopeStatus:         scopeStatus.INACTIVE,
	}

	return &response
}
