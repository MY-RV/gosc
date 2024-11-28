package gosc

import (
	"reflect"
	"testing"

	"github.com/MY-RV/gosc/scope"
	"github.com/MY-RV/gosc/scope_status"
)

type (
	singletonService struct {
		value int
	}
)

func TestServiceCollectionBuilder_AddSingleton(t *testing.T) {
	t.Run("Add Singleton Service", func(t *testing.T) {
		builder := NewServiceCollectionBuilder()

		factoryFunc := func() singletonService { return singletonService{  value: 1 } }
		builder.AddSingleton(factoryFunc)

		serviceDef := builder.serviceCollection.servicesDefinitions[reflect.TypeOf(new(singletonService))] // Expected definition type

		if serviceDef.serviceScope != scope.SINGLETON {
			t.Errorf("expected scope SINGLETON, got %d", serviceDef.serviceScope)
		}
		if serviceDef.serviceFactory == nil {
			t.Errorf("service factory should not be nil")
		}
	})
}

func TestServiceCollectionBuilder_AddScoped(t *testing.T) {
	t.Run("Add Scoped Service", func(t *testing.T) {
		builder := NewServiceCollectionBuilder()

		factoryFunc := func() serviceImplementation { return serviceImplementation{  value: 1 } }
		builder.AddScoped(factoryFunc)

		serviceDef := builder.serviceCollection.servicesDefinitions[reflect.TypeOf(new(serviceImplementation))]

		if serviceDef.serviceScope != scope.SCOPED {
			t.Errorf("expected scope SCOPED, got %d", serviceDef.serviceScope)
		}
		if serviceDef.serviceFactory == nil {
			t.Errorf("service factory should not be nil")
		}
	})
}

func TestServiceCollectionBuilder_Build(t *testing.T) {
	t.Run("Build Service Collection", func(t *testing.T) {
		builder := NewServiceCollectionBuilder()

		factoryFunc := func() string { return "service" }
		builder.AddSingleton(factoryFunc)

		serviceCollection := builder.Build()

		if serviceCollection.scopeStatus != scope_status.INACTIVE {
			t.Errorf("expected scope status INACTIVE, got %d", serviceCollection.scopeStatus)
		}

		if len(serviceCollection.servicesDefinitions) != 1 {
			t.Errorf("expected 1 service definition, got %d", len(serviceCollection.servicesDefinitions))
		}
	})
}
