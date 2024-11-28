package gosc

import (
	"testing"
)

func TestTypedFactoryFunc_NewTypedFactoryFunc_WithNonFunction(t *testing.T) {
	_, err := NewTypedFactoryFunc[any]("not a function")
	if err == nil {
		t.Errorf("expected error, got nil")
	}
}

func TestTypedFactoryFunc_NewTypedFactoryFunc_NumOutShouldBe1(t *testing.T) {
	_, err := NewTypedFactoryFunc[any](func() (any, any) { return nil, nil })
	if err == nil {
		t.Errorf("expected error, got nil")
	}

	_, err = NewTypedFactoryFunc[any](func() {})
	if err == nil {
		t.Errorf("expected error, got nil")
	}
}

func TestTypedFactoryFunc_TypedFactoryFunc_NumOutShouldBe1(t *testing.T) {
	(func() {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("expected panic, got success")
			}
		}()

		TypedFactoryFunc[any](func() (any, any) { return nil, nil })
	})()

	(func() {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("expected panic, got success")
			}
		}()

		TypedFactoryFunc[any](func() {})
	})()
}
