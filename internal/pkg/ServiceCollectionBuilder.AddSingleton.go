package pkg

import "github.com/MY-RV/gosc/enums/scope"

func (serviceCollectionBuilder *ServiceCollectionBuilder) AddSingleton(registration any) {
	err := serviceCollectionBuilder._tryAddService(registration, scope.SINGLETON)
	if err != nil {
		panic(err)
	}
}
