package pkg

import "github.com/MY-RV/gosc/scope"

func (serviceCollectionBuilder *ServiceCollectionBuilder) AddScoped(registration any) {
	err := serviceCollectionBuilder._tryAddService(registration, scope.SCOPED)
	if err != nil {
		panic(err)
	}
}
