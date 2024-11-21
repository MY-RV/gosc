package gosc

import (
	"reflect"

	"github.com/MY-RV/gosc/scope"
	"github.com/MY-RV/gosc/scope_status"
)

type Service interface{ ~struct{} | interface{} }
type ServiceCollection struct {
	activeSingletonServices map[reflect.Type]*any
	activeScopedServices    map[reflect.Type]*any

	servicesDefinitions map[reflect.Type]ServiceDefinition
	scopeStatus         int
}

func GetService[T Service](serviceCollection *ServiceCollection) T {
	serviceType := goElemIfNonElem(reflect.TypeOf(new(T)))
	boxedService := serviceCollection.getService(serviceType)
	unboxedService, _ := (boxedService).(T)
	return unboxedService
}

func (serviceCollection *ServiceCollection) getService(reflectType reflect.Type) any {
	serviceDefinition := serviceCollection.servicesDefinitions[ensureElem(reflectType)]

	var service *any
	switch serviceDefinition.serviceScope {
	case scope.SINGLETON:
		service = serviceCollection._getServiceSingleton(reflectType)
	case scope.SCOPED:
		service = serviceCollection._getServiceScoped(reflectType)
	default:
		panic("invalid service scope")
	}

	if reflectType.Kind() == reflect.Ptr {
		return service
	}

	return *service
}

func (serviceCollection *ServiceCollection) _getServiceSingleton(reflectType reflect.Type) *any {
	return mapGetOrSet(serviceCollection.activeSingletonServices, reflectType, func() *any {
		serviceDefinition := serviceCollection.servicesDefinitions[reflectType]
		serviceFactory := serviceDefinition.serviceFactory

		boxedService := serviceFactory.createService(serviceCollection)
		return boxedService
	})
}

func (serviceCollection *ServiceCollection) _getServiceScoped(reflectType reflect.Type) *any {
	if serviceCollection.scopeStatus == scope_status.INACTIVE {
		panic("scope is not active")
	}

	return mapGetOrSet(serviceCollection.activeScopedServices, reflectType, func() *any {
		serviceDefinition := serviceCollection.servicesDefinitions[reflectType]
		serviceFactory := serviceDefinition.serviceFactory

		boxedService := serviceFactory.createService(serviceCollection)
		return boxedService
	})
}

func (serviceCollection *ServiceCollection) NewScope() *ServiceCollection {
	return &ServiceCollection{
		activeSingletonServices: serviceCollection.activeSingletonServices,
		activeScopedServices:    make(map[reflect.Type]*any),

		servicesDefinitions: serviceCollection.servicesDefinitions,
		scopeStatus:         scope_status.ACTIVE,
	}
}
