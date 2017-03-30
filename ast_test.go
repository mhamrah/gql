package gql

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTypeName(t *testing.T) {
	name := "InputTypeName"

	tests := []struct {
		input    TypeDefinition
		expected TypeKind
	}{
		{ObjectDefinition{Name: name}, TypeKind_OBJECT},
		{ScalarDefinition{Name: name}, TypeKind_SCALAR},
		{UnionDefinition{Name: name}, TypeKind_UNION},
		{InterfaceDefinition{Name: name}, TypeKind_INTERFACE},
		{InputObjectDefinition{Name: name}, TypeKind_INPUT_OBJECT},
		{EnumDefinition{Name: name}, TypeKind_ENUM},
	}

	for _, test := range tests {
		t.Run(reflect.TypeOf(test).Name(), func(t *testing.T) {
			assert.Equal(t, name, test.input.TypeName())
			assert.Equal(t, test.expected, test.input.TypeKind())
		})
	}

}
