package gql

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestObjectDefinitionGenInterface(t *testing.T) {
	input := ObjectDefinition{
		Name:       "Query",
		Implements: nil,
		Fields: []FieldDefinition{
			FieldDefinition{
				Name:       "someField",
				Arguments:  nil,
				Type:       TypeDescription{Name: "String"},
				Directives: nil,
			},
		},
	}

	var b bytes.Buffer
	err := input.genInterface(&b)
	assert.NoError(t, err)
	result := b.String()
	assert.NotEmpty(t, result)
	assert.Equal(t, "\n\t\ttype Query interface {\n\t\t\tSomeField(ctx context.Context) (string, error)\n\t\t\t\n\t\t}\n\t", result)

	b.Reset()
	err = input.genService(&b)
	assert.NoError(t, err)
	result = b.String()
	assert.Equal(t, "\n\t\ttype query_impl struct {\n\t\t\timpl Query\n\t\t}\n\n\t\tfunc New(impl Query) gql.Service {\n\t\t\treturn query_impl{ impl: impl }\n\t\t}\n\n\t\tfunc (s query_impl) Handlers() map[string]gql.HandlerFunc {\n\t\t\treturn map[string]gql.HandlerFunc{\n\t\t\t\t\"someField\": s.SomeField,\n\t\t\t}\n\t\t}\n\n\t\t\n\t\t\tfunc (s query_impl) SomeField(ctx context.Context, operation gql.Selection) (gql.Selectable, error) {\n\t\t\t\t\n\n\t\t\t\tsuccess, err := s.impl.SomeField(ctx)\n\t\t\t\tif err != nil {\n\t\t\t\t\treturn nil, err\n\t\t\t\t}\n\t\t\t\treturn success, nil\n\t\t\t}\n\t\t\n\t", result)
}
