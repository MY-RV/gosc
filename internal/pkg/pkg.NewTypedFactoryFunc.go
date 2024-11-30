package pkg

import (
	"fmt"
	"reflect"

	"github.com/MY-RV/gosc/internal/helpers/factory_func_helper"
	"github.com/MY-RV/gosc/internal/helpers/func_helper"
)

func NewTypedFactoryFunc[T any](factoryFunction any) (*ServiceFactory[T], error) {
	factoryFuncType, err := factory_func_helper.ValidateFactoryFuncArg(factoryFunction)
	if err != nil {
		return nil, err
	}

	argTypes := func_helper.GetArgsTypes(factoryFuncType)
	factoryFuncValue := reflect.ValueOf(factoryFunction)

	factory := &ServiceFactory[T]{
		dependencyTypes: argTypes,
		serviceType:     factoryFuncType.Out(0),
		createService: func(serviceCollection *ServiceCollection) (T, error) {
			args, _ := serviceCollection.resolveServices(argTypes)
			callResults := factoryFuncValue.Call(args)
			result := callResults[0]

			value := result.Interface()
			response, successUnboxing := value.(T)
			if !successUnboxing {
				return response, fmt.Errorf("factory function must return a %s", factoryFuncType.Out(0).String())
			}

			return response, nil
		},
	}

	return factory, nil
}
