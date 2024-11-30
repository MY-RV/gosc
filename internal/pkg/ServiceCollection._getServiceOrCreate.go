package pkg

import (
	"fmt"
	"reflect"

	"github.com/MY-RV/gosc/internal/helpers/pointer_helper"
)

func (serviceCollection *ServiceCollection) _getServiceOrCreate(reflectType reflect.Type, registry map[reflect.Type]*any) (*any, error) {
	serviceDefinition := serviceCollection.servicesDefinitions[reflectType]
	if serviceDefinition == nil {
		return nil, fmt.Errorf("service definition not found for type %s", reflectType.String())
	}

	if servicePointer, exists := registry[reflectType]; exists {
		return servicePointer, nil
	}

	boxedService, err := serviceDefinition.factory.CreateService(serviceCollection)
	if err != nil {
		return nil, err
	}

	pointerValue := pointer_helper.ValueGoPointerIfNotPointer(reflect.ValueOf(boxedService))
	servicePointer := pointerValue.Interface().(*any)
	registry[reflectType] = servicePointer

	return servicePointer, nil
}
