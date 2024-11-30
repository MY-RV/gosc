package pkg

import (
	"fmt"
	"reflect"

	"github.com/MY-RV/gosc/scope_status"
)

func (serviceCollection *ServiceCollection) _getServiceScoped(reflectType reflect.Type) (*any, error) {
	if serviceCollection.scopeStatus == scope_status.INACTIVE {
		return nil, fmt.Errorf("scope is not active")
	}

	return serviceCollection._getServiceOrCreate(reflectType, serviceCollection.activeServices.scoped)
}
