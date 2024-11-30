package pkg

import (
	"reflect"

	"github.com/MY-RV/gosc/internal/helpers/pointer_helper"
)

func (serviceCollection *ServiceCollection) resolveServices(argTypes []reflect.Type) ([]reflect.Value, error) {
	args := make([]reflect.Value, len(argTypes))
	for i, argType := range argTypes {
		service, err := serviceCollection.getService(argType)
		if err != nil {
			return nil, err
		}

		args[i] = pointer_helper.ValueGoPointerIfNotPointer(reflect.ValueOf(service))
	}

	return args, nil
}
