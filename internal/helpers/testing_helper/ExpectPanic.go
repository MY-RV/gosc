package testing_helper

import "testing"

func ExpectPanic(t *testing.T) {
	if r := recover(); r == nil {
		t.Errorf("expected panic, got success")
	}
}
