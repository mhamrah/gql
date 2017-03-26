package gql

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTypeName(t *testing.T) {
	name := "InputTypeName"

	types := []TypeDefinition{
		ObjectDefinition{Name: name},
		ScalarDefinition{Name: name},
		UnionDefinition{Name: name},
		InterfaceDefinition{Name: name},
		InputObjectDefinition{Name: name},
		EnumDefinition{Name: name},
		InputValueDefinition{Name: name},
	}

	for _, test := range types {
		t.Run(reflect.TypeOf(test).Name(), func(t *testing.T) {
			assert.Equal(t, name, test.TypeName())
		})
	}

}
