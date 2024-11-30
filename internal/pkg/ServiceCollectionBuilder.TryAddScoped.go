package pkg

import "github.com/MY-RV/gosc/enums/scope"

func (serviceCollectionBuilder *ServiceCollectionBuilder) TryAddScoped(registration any) error {
	return serviceCollectionBuilder._tryAddService(registration, scope.SCOPED)
}
