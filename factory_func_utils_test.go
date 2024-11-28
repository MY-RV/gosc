package gosc

import (
	"reflect"
	"testing"
)

func TestFactoryFuncUtils_validateNewFactoryFuncArg_factoryFuncTypeKindIsNotFunc(t *testing.T) {
	_, err := NewTypedFactoryFunc[any]("not a function")

	if err == nil {
		t.Errorf("expected error, got nil")
	}
}

func TestFactoryFuncUtils_validateNewFactoryFuncArg_factoryFuncTypeNumOutShouldBe1(t *testing.T) {
	_, err := NewTypedFactoryFunc[any](func() (any, any) { return nil, nil })
	if err == nil {
		t.Errorf("expected error, got nil")
	}

	_, err = NewTypedFactoryFunc[any](func() {})
	if err == nil {
		t.Errorf("expected error, got nil")
	}

	_, err = NewTypedFactoryFunc[any](func() any { return nil })
	if err != nil {
		t.Errorf("expected nil, got error: %v", err)
	}
}

func (s *serviceImplementation) Get() int {
	return s.value
}

func createServiceCollection() *ServiceCollection {
	serviceCollectionBuilder := NewServiceCollectionBuilder()

	serviceCollectionBuilder.AddSingleton(
		func() serviceInterface {
			return &serviceImplementation{value: 1}
		},
	)
	serviceCollectionBuilder.AddSingleton(
		func() serviceImplementation {
			return serviceImplementation{value: 1}
		},
	)

	return serviceCollectionBuilder.Build()
}

func TestFactoryFuncUtils_resolveFactoryDependencies_TestWithInterfaceDependency(t *testing.T) {
	serviceCollection := createServiceCollection()

	factoryFuncArgsTypes, _ := getFuncArgsTypes(reflect.TypeOf(
		func(service serviceInterface) { },
	))

	dependencies, _ := resolveFactoryDependencies(serviceCollection, factoryFuncArgsTypes)
	if len(dependencies) != 1 {
		t.Errorf("expected 1 dependency, got %d", len(dependencies))
	}

	dependency := dependencies[0].Interface()
	unboxed, successUnbox := dependency.(serviceInterface)
	if !successUnbox {
		t.Errorf("expected dependency to be of type %v, got %v", factoryFuncArgsTypes[0], reflect.TypeOf(dependency))
	} else if got := unboxed.Get(); got != 1 {
		t.Errorf("expected value to be 1, got %d", got)
	}
}

func TestFactoryFuncUtils_resolveFactoryDependencies_TestWithInterfacePointerDependency(t *testing.T) {
	serviceCollection := createServiceCollection()

	factoryFuncArgsTypes, _ := getFuncArgsTypes(reflect.TypeOf(
		func(service *serviceInterface) { },
	))

	dependencies, _ := resolveFactoryDependencies(serviceCollection, factoryFuncArgsTypes)
	if len(dependencies) != 1 {
		t.Errorf("expected 1 dependency, got %d", len(dependencies))
	}

	dependency := dependencies[0].Interface()
	unboxed, successUnbox := dependency.(serviceInterface)
	if !successUnbox {
		t.Errorf("expected dependency to be of type %v, got %v", factoryFuncArgsTypes[0], reflect.TypeOf(dependency))
	} else if got := unboxed.Get(); got != 1 {
		t.Errorf("expected value to be 1, got %d", got)
	}
}

func TestFactoryFuncUtils_resolveFactoryDependencies_TestWithStructDependency(t *testing.T) {
	serviceCollection := createServiceCollection()

	factoryFuncArgsTypes, _ := getFuncArgsTypes(reflect.TypeOf(
		func(service serviceImplementation) { },
	))

	dependencies, _ := resolveFactoryDependencies(serviceCollection, factoryFuncArgsTypes)
	if len(dependencies) != 1 {
		t.Errorf("expected 1 dependency, got %d", len(dependencies))
	}

	dependency := dependencies[0].Interface()
	unboxed, successUnbox := dependency.(serviceImplementation)
	if !successUnbox {
		t.Errorf("expected dependency to be of type %v, got %v", factoryFuncArgsTypes[0], reflect.TypeOf(dependency))
	} else if got := (unboxed).Get(); got != 1 {
		t.Errorf("expected value to be 1, got %d", got)
	}
}

func TestFactoryFuncUtils_resolveFactoryDependencies_TestWithStructPointerDependency(t *testing.T) {
	serviceCollection := createServiceCollection()

	factoryFuncArgsTypes, _ := getFuncArgsTypes(reflect.TypeOf(
		func(service *serviceImplementation) { },
	))

	dependencies, _ := resolveFactoryDependencies(serviceCollection, factoryFuncArgsTypes)
	if len(dependencies) != 1 {
		t.Errorf("expected 1 dependency, got %d", len(dependencies))
	}

	dependency := dependencies[0].Interface()
	unboxed, successUnbox := dependency.(*serviceImplementation)
	if !successUnbox {
		t.Errorf("expected dependency to be of type %v, got %v", factoryFuncArgsTypes[0], reflect.TypeOf(dependency))
	} else if got := (unboxed).Get(); got != 1 {
		t.Errorf("expected value to be 1, got %d", got)
	}
}
