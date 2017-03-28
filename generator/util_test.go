package generator

import (
	"fmt"
	"testing"

	"github.com/mhamrah/gql"
	"github.com/stretchr/testify/assert"
)

func TestTypeMapping(t *testing.T) {
	tests := []struct {
		graphqlType string
		goType      string
	}{
		{"String", "string"},
		{"other", "other"},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("map %v to %v", test.graphqlType, test.goType), func(t *testing.T) {
			result := typeMap(test.graphqlType)
			assert.Equal(t, test.goType, result)
		})
	}
}

func TypeIntializerTest(t *testing.T) {
	tests := []struct {
		title    string
		ivd      gql.InputValueDefinition
		expected string
	}{
		{"string",
			gql.InputValueDefinition{
				Name: "p1",
				Type: gql.TypeDescription{
					Name: "String",
				},
			},
			"p1, err := operation.Field.StringValue(\"p1\", \"\", false)",
		},
	}

	for _, test := range tests {
		t.Run(test.title, func(t *testing.T) {
			result := typeInitializer(test.ivd)
			assert.Equal(t, test.expected, result)
		})
	}
}
