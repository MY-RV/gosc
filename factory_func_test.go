package gosc

import (
	"testing"
)

func TestFactoryFunc_NewFactoryFunc_WithNonFunction(t *testing.T) {
	_, err := NewFactoryFunc("not a function")
	if err == nil {
		t.Errorf("expected error, got nil")
	}
}

func TestFactoryFunc_NewFactoryFunc_NumOutShouldBe1(t *testing.T) {
	_, err := NewFactoryFunc(func() (any, any) { return nil, nil })
	if err == nil {
		t.Errorf("expected error, got nil")
	}

	_, err = NewFactoryFunc(func() {})
	if err == nil {
		t.Errorf("expected error, got nil")
	}
}

func TestFactoryFunc_FactoryFunc_NumOutShouldBe1(t *testing.T) {
	(func() {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("expected panic, got success")
			}
		}()

		FactoryFunc(func() (any, any) { return nil, nil })
	})()

	(func() {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("expected panic, got success")
			}
		}()

		FactoryFunc(func() {})
	})()
}
