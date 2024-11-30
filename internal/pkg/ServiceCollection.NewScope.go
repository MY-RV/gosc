package pkg

import (
	"reflect"

	"github.com/MY-RV/gosc/scope_status"
)

func (serviceCollection *ServiceCollection) NewScope() *ServiceCollection {

	response := ServiceCollection{
		activeServices: ServiceCollectionActiveServices{
			singleton: serviceCollection.activeServices.singleton,
			scoped:    make(map[reflect.Type]*any),
		},

		servicesDefinitions: serviceCollection.servicesDefinitions,
		scopeStatus:         scope_status.ACTIVE,
	}

	return &response
}
