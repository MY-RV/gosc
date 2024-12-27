package pkg

import (
	"reflect"

	"github.com/MY-RV/gosc/enums/scopeStatus"
)

func (serviceCollection *ServiceCollection) NewScope() *ServiceCollection {

	response := ServiceCollection{
		activeServices: ServiceCollectionActiveServices{
			singleton: serviceCollection.activeServices.singleton,
			scoped:    make(map[reflect.Type]*any),
		},

		servicesDefinitions: serviceCollection.servicesDefinitions,
		scopeStatus:         scopeStatus.ACTIVE,
	}

	return &response
}
