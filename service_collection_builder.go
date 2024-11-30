package gosc

import (
	"reflect"

	"github.com/MY-RV/gosc/scope"
	"github.com/MY-RV/gosc/scope_status"
)

type ServiceCollectionBuilder struct {
	serviceCollection ServiceCollection
}

func NewServiceCollectionBuilder() *ServiceCollectionBuilder {
	return &ServiceCollectionBuilder{
		serviceCollection: ServiceCollection{
			activeSingletonServices: make(map[reflect.Type]*any),
			activeScopedServices:    make(map[reflect.Type]*any),

			servicesDefinitions: make(map[reflect.Type]ServiceDefinition),
			scopeStatus:         scope_status.INACTIVE,
		},
	}
}

func (serviceCollectionBuilder *ServiceCollectionBuilder) AddSingleton(registration any) {
	serviceCollectionBuilder.addService(registration, scope.SINGLETON)
}

func (serviceCollectionBuilder *ServiceCollectionBuilder) AddScoped(registration any) {
	serviceCollectionBuilder.addService(registration, scope.SCOPED)
}

func (serviceCollectionBuilder *ServiceCollectionBuilder) addService(registration any, scope scope.Enum) {
	servicesDefinitions := serviceCollectionBuilder.serviceCollection.servicesDefinitions

	var serviceFactory ServiceFactory[any]
	var err error
	var ok bool

	if serviceFactory, ok = registration.(ServiceFactory[any]); ok {
		// do nothing
	} else if reflect.TypeOf(registration).Kind() == reflect.Func {
		serviceFactory, err = NewTypedFactoryFunc[any](registration)

		if err != nil {
			panic(err)
		}
	} else {
		panic("invalid registration")
	}

	definitionType := serviceFactory.getDefinitionType()
	servicesDefinitions[definitionType] = ServiceDefinition{
		serviceType:    definitionType,
		serviceFactory: serviceFactory,
		serviceScope:   scope,
	}
}

func (serviceCollectionBuilder *ServiceCollectionBuilder) Build() *ServiceCollection {
	serviceCollection := serviceCollectionBuilder.serviceCollection

	response := ServiceCollection{
		activeSingletonServices: serviceCollection.activeSingletonServices,
		activeScopedServices:    make(map[reflect.Type]*any),

		servicesDefinitions: serviceCollection.servicesDefinitions,
		scopeStatus:         scope_status.INACTIVE,
	}

	return &response
}
