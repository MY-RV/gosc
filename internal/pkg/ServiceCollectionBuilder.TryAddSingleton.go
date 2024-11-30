package pkg

import "github.com/MY-RV/gosc/scope"

func (serviceCollectionBuilder *ServiceCollectionBuilder) TryAddSingleton(registration any) error {
	return serviceCollectionBuilder._tryAddService(registration, scope.SINGLETON)
}
