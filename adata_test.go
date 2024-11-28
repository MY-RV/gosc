
package gosc

type (
	serviceInterface interface {
		Get() int
	}

	serviceImplementation struct {
		value int
	}
)
