package pkg

import (
	"fmt"
	"reflect"

	"github.com/MY-RV/gosc/scope"
)

func (serviceCollectionBuilder *ServiceCollectionBuilder) _tryAddService(registration any, scope scope.Enum) error {
	servicesDefinitions := serviceCollectionBuilder.serviceCollection.servicesDefinitions

	var serviceFactory *ServiceFactory[any]
	var err error
	var ok bool

	if serviceFactory, ok = registration.(*ServiceFactory[any]); !ok {
		if reflect.TypeOf(registration).Kind() == reflect.Func {
			serviceFactory, err = NewTypedFactoryFunc[any](registration)
			if err != nil {
				return err
			}
		} else {
			return fmt.Errorf("invalid registration")
		}
	}

	if serviceFactory == nil {
		return fmt.Errorf("service factory is nil")
	}

	definitionType := serviceFactory.GetServiceType()
	servicesDefinitions[definitionType] = &ServiceDefinition[any]{
		factory: *serviceFactory,
		scope:   scope,
	}

	return nil
}
