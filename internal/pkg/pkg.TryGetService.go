package pkg

import (
	"fmt"
	"reflect"

	"github.com/MY-RV/gosc/internal/helpers/pointer_helper"
)

func TryGetService[T any](serviceCollection *ServiceCollection) (T, error) {
	serviceType := pointer_helper.TypeGoElem(reflect.TypeOf(new(T)))
	boxedService, err := serviceCollection.getService(serviceType)
	unboxedService, successUnboxing := (boxedService).(T)

	if err != nil {
		return unboxedService, err
	}

	if !successUnboxing {
		return unboxedService, fmt.Errorf("unboxing failed")
	}

	return unboxedService, nil
}
