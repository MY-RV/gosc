package gosc

import (
	//"fmt"
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
	reflectTypePointer := goPointerIfNonPointer(reflectType)
	serviceDefinition := serviceCollection.servicesDefinitions[reflectTypePointer]
	
	var service *any
	switch serviceDefinition.serviceScope {
	case scope.SINGLETON:
		service = serviceCollection._getServiceSingleton(reflectTypePointer)
	case scope.SCOPED:
		service = serviceCollection._getServiceScoped(reflectTypePointer)
	default:
		panic("invalid service scope")
	}

	response := *service

	if reflectType.Kind() == reflect.Ptr || reflectType.Kind() == reflect.Interface {
		return response
	}

	elem := goValueElemIfPointer(reflect.ValueOf(response))
	return elem.Interface()

	/* fmt.Println()
	fmt.Printf("|response: %v, %v\n", reflectType, reflectType)
	fmt.Printf("|response: %v, %v\n", response, reflect.TypeOf(response))
	fmt.Printf("|response: %v, %v\n", &response, reflect.TypeOf(&response))
	fmt.Println() */

	/* if reflectType.Kind() != reflect.Ptr {
		return &response
	} */

	return response
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
