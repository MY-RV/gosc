package pkg

import (
	"reflect"
)

func (serviceCollection *ServiceCollection) _getServiceSingleton(reflectType reflect.Type) (*any, error) {
	return serviceCollection._getServiceOrCreate(reflectType, serviceCollection.activeServices.singleton)
}
