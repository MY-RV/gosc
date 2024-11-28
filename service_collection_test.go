package gosc

import (
	"testing"

	"github.com/MY-RV/gosc/scope_status"
)

func TestServiceCollection_GetService(t *testing.T) {
	t.Run("Retrieve Singleton Service", func(t *testing.T) {
		expectedValue := 3
		builder := NewServiceCollectionBuilder()
		factoryFunc := func() serviceImplementation { return serviceImplementation{value: expectedValue} }
		builder.AddSingleton(factoryFunc)
		serviceCollection := builder.Build()

		service := GetService[serviceImplementation](serviceCollection)

		if service.value != expectedValue {
			t.Errorf("expected %d, got %d", expectedValue, service.value)
		}
	})

	t.Run("Retrieve Scoped Service", func(t *testing.T) {
		expectedValue := 4
		builder := NewServiceCollectionBuilder()
		factoryFunc := func() serviceImplementation { return serviceImplementation{value: expectedValue} }
		builder.AddScoped(factoryFunc)
		serviceCollection := builder.Build()

		(func() {
			defer func() {
				if r := recover(); r == nil {
					t.Errorf("expected panic, got success")
				}
			}()

			GetService[serviceImplementation](serviceCollection)
		})()

		serviceCollection = serviceCollection.NewScope()
		service := GetService[serviceImplementation](serviceCollection)

		if service.value != expectedValue {
			t.Errorf("expected %d, got %d", expectedValue, service.value)
		}
	})
}

func TestServiceCollection_NewScope(t *testing.T) {
	t.Run("Create New Scope", func(t *testing.T) {
		builder := NewServiceCollectionBuilder()
		factoryFunc := func() string { return "scoped service" }
		builder.AddScoped(factoryFunc)
		serviceCollection := builder.Build()

		newScope := serviceCollection.NewScope()

		if newScope.scopeStatus != scope_status.ACTIVE {
			t.Errorf("expected scope status ACTIVE, got %d", newScope.scopeStatus)
		}

		if len(newScope.activeScopedServices) != 0 {
			t.Errorf("expected 0 active scoped services, got %d", len(newScope.activeScopedServices))
		}
	})
}

/* func TestServiceCollection_getServiceSingleton(t *testing.T) {
	t.Run("Retrieve Singleton Service", func(t *testing.T) {
		serviceCollection := &ServiceCollection{
			activeSingletonServices: make(map[reflect.Type]*any),
			servicesDefinitions: map[reflect.Type]ServiceDefinition{
				reflect.TypeOf(""): {
					serviceScope:   scope.SINGLETON,
					serviceFactory: FactoryFunc(func() {}),
				},
			},
		}

		service := serviceCollection._getServiceSingleton(reflect.TypeOf(""))

		if service == nil || *service != "singleton service" {
			t.Errorf("expected 'singleton service', got '%v'", service)
		}
	})
} */
