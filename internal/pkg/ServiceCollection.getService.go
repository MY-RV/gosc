package pkg

import (
	"fmt"
	"reflect"

	"github.com/MY-RV/gosc/internal/helpers/pointer_helper"
	"github.com/MY-RV/gosc/scope"
)

func (serviceCollection *ServiceCollection) getService(reflectType reflect.Type) (any, error) {
	reflectTypePointer := pointer_helper.TypeGoPointerIfNotPointer(reflectType)
	serviceDefinition := serviceCollection.servicesDefinitions[reflectTypePointer]

	var service *any
	var err error

	switch serviceDefinition.scope {
	case scope.SINGLETON:
		service, err = serviceCollection._getServiceSingleton(reflectTypePointer)
	case scope.SCOPED:
		service, err = serviceCollection._getServiceScoped(reflectTypePointer)
	default:
		return nil, fmt.Errorf("invalid service scope")
	}

	if err != nil {
		return nil, err
	}

	response := *service

	if reflectType.Kind() == reflect.Ptr || reflectType.Kind() == reflect.Interface {
		return response, nil
	}

	elem := pointer_helper.ValueEnsureElem(reflect.ValueOf(response))
	return elem.Interface(), nil
}
