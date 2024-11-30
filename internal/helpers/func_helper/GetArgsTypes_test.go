package func_helper

import (
	"reflect"
	"testing"

	"github.com/MY-RV/gosc/internal/helpers/testing_helper"
)

func Test_GetArgsTypes(t *testing.T) {
	t.Run("Call with non {func} argument", func(t *testing.T) {
		defer testing_helper.ExpectPanic(t)
		GetArgsTypes(reflect.TypeOf(""))
	})

	t.Run("{func} without arguments", func(t *testing.T) {
		got := GetArgsTypes(reflect.TypeOf(func() {}))
		if len(got) != 0 {
			t.Errorf("expected empty slice, got %v", got)
		}
	})

	t.Run("{func} with one argument", func(t *testing.T) {
		got := GetArgsTypes(reflect.TypeOf(func(arg1 string) {}))
		if len(got) != 1 {
			t.Errorf("expected 1 argument, got %d", len(got))
		}

		if got[0] != reflect.TypeOf("") {
			t.Errorf("expected string argument, got %v", got[0])
		}
	})

	t.Run("{func} with [n] argument", func(t *testing.T) {
		got := GetArgsTypes(reflect.TypeOf(func(arg1 int, arg2 testing_helper.Service1, arg3 *testing_helper.Service2) {}))
		expect := []reflect.Type{
			reflect.TypeOf(0),
			reflect.TypeOf(new(testing_helper.Service1)).Elem(),
			reflect.TypeOf(new(*testing_helper.Service2)).Elem(),
		}
		
		if len(got) != len(expect) {
			t.Errorf("expected 1 argument, got %d", len(got))
		}

		for i := 0; i < len(expect); i++ {
			if got[i] != expect[i] {
				t.Errorf("expected %v, got %v", expect[i], got[i])
			}
		}
	})
}
