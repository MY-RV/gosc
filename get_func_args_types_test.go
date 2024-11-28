package gosc

import (
	"reflect"
	"testing"
)

func TestGetFuncArgsTypes(t *testing.T) {
	tests := []struct {
		name         string
		inputFunc    interface{}
		expectedArgs []reflect.Type
		expectError  bool
	}{
		{
			name: "No arguments",
			inputFunc: func() {},
			expectedArgs: []reflect.Type{},
			expectError: false,
		},
		{
			name: "Single argument",
			inputFunc: func(a int) {},
			expectedArgs: []reflect.Type{reflect.TypeOf(0)},
			expectError: false,
		},
		{
			name: "Multiple arguments",
			inputFunc: func(a int, b string, c bool) {},
			expectedArgs: []reflect.Type{
				reflect.TypeOf(0),
				reflect.TypeOf(""),
				reflect.TypeOf(true),
			},
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			factoryFuncType := reflect.TypeOf(tt.inputFunc)

			argTypes, err := getFuncArgsTypes(factoryFuncType)

			if tt.expectError {
				if err == nil {
					t.Errorf("expected error but got nil")
				}
				return
			}

			if err != nil {
				t.Errorf("unexpected error: %v", err)
				return
			}

			if len(argTypes) != len(tt.expectedArgs) {
				t.Errorf("expected %d argument types, got %d", len(tt.expectedArgs), len(argTypes))
			}

			for i, expectedType := range tt.expectedArgs {
				if argTypes[i] != expectedType {
					t.Errorf("expected argument type %v at index %d, got %v", expectedType, i, argTypes[i])
				}
			}
		})
	}
}
