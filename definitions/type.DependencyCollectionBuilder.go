package definitions

import "github.com/MY-RV/gosc/enums/scope"

type DependencyCollectionBuilder interface {
	AddRegistry(scope.Enum, Factory) 
	AddSingleton(Factory)
	AddTransient(Factory)
	AddScoped(Factory)

	Build() DependencyCollection
}
