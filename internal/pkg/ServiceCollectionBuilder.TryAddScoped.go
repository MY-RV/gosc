package pkg

import "github.com/MY-RV/gosc/scope"

func (serviceCollectionBuilder *ServiceCollectionBuilder) TryAddScoped(registration any) error {
	return serviceCollectionBuilder._tryAddService(registration, scope.SCOPED)
}
