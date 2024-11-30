package pkg

import "testing"

func Test_FactoryFunc(t *testing.T) {
	t.Run("Call with non {func} argument", func(t *testing.T) {
		_, err := NewFactoryFunc("not a function")
		if err == nil {
			t.Errorf("expected error, got nil")
		}
	})

	t.Run("{func} argument should have 1 return", func(t *testing.T) {
		_, err := NewFactoryFunc(func() (any, any) { return nil, nil })
		if err == nil {
			t.Errorf("expected error, got nil")
		}

		_, err = NewFactoryFunc(func() {})
		if err == nil {
			t.Errorf("expected error, got nil")
		}
	})

}
