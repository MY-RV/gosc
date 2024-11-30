package factory_func_helper

import "testing"

func Test_ValidateFactoryFuncArg(t *testing.T) {
	t.Run("Call with non {func} argument", func(t *testing.T) {
		_, err := ValidateFactoryFuncArg("not a function")
		if err == nil {
			t.Errorf("expected error, got nil")
		}
	})

	t.Run("{func} argument should have 1 return", func(t *testing.T) {
		_, err := ValidateFactoryFuncArg(func() (any, any) { return nil, nil })
		if err == nil {
			t.Errorf("expected error, got nil")
		}

		_, err = ValidateFactoryFuncArg(func() {})
		if err == nil {
			t.Errorf("expected error, got nil")
		}
	})

}
