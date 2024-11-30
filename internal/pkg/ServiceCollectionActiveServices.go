package pkg

import (
	"reflect"
)

type ServiceCollectionActiveServices struct {
	singleton map[reflect.Type]*any
	scoped    map[reflect.Type]*any
}
