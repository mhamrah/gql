package validator

import (
	"testing"

	"github.com/mhamrah/gql"
	"github.com/mhamrah/gql/parser"
	"github.com/stretchr/testify/assert"
)

var (
	queryOp = gql.OperationTypeDefinition{
		OpType: gql.Query,
		Name:   "SomeQuery",
	}

	queryType = gql.ObjectDefinition{
		Name: "Query",
		Fields: []gql.FieldDefinition{
			gql.FieldDefinition{
				Name: "Foo",
				Type: gql.TypeDescription{Name: "String"},
			},
		},
	}

	someQueryType = gql.ObjectDefinition{
		Name: "SomeQuery",
		Fields: []gql.FieldDefinition{
			gql.FieldDefinition{
				Name: "Foo",
				Type: gql.TypeDescription{Name: "String"},
			},
		},
	}

	scalarQuery = gql.ScalarDefinition{Name: "Query"}
)

func TestBuildSchema(t *testing.T) {

	testCases := []struct {
		title    string
		opDefs   []gql.OperationTypeDefinition
		typeDefs []gql.TypeDefinition
	}{
		{"queryType", nil, []gql.TypeDefinition{queryType}},
	}

	for _, test := range testCases {
		t.Run(test.title, func(t *testing.T) {
			schema := parser.Schema{
				OperationTypeDefinitions: test.opDefs,
				TypeDefinitions:          test.typeDefs,
			}
			result, err := BuildSchema(schema)
			assert.NoError(t, err)
			assert.Equal(t, test.typeDefs[0], result.QueryType)
		})
	}

}

func TestBuildSchemaErrors(t *testing.T) {
	testCases := []struct {
		title       string
		opDefs      []gql.OperationTypeDefinition
		typeDefs    []gql.TypeDefinition
		expectedErr error
	}{
		{"queryType", nil, []gql.TypeDefinition{scalarQuery}, ErrInvalidQueryType},
	}

	for _, test := range testCases {
		t.Run(test.title, func(t *testing.T) {
			schema := parser.Schema{
				OperationTypeDefinitions: test.opDefs,
				TypeDefinitions:          test.typeDefs,
			}
			result, err := BuildSchema(schema)
			assert.Equal(t, test.expectedErr, err)
			assert.Nil(t, result)

		})
	}

}
