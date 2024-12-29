package dependencyCollection

import (
	"github.com/MY-RV/gosc/definitions"
)

// Implements github.com/MY-RV/gosc/definitions::DependencyCollectionBuilder
type dependencyCollectionBuilder struct {
	dependencyCollection dependencyCollection
}

// TYPE CHECKING
func init() {
	var _ definitions.DependencyCollectionBuilder = &dependencyCollectionBuilder{}
}
