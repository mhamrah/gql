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
		title   string
		input   gql.ObjectDefinition
		typeDef string
	}{
		{"queryType", queryType, "Query"},
	}

	for _, test := range testCases {
		t.Run(test.title, func(t *testing.T) {
			schema := parser.Schema{
				TypeDefinitions: map[string]gql.TypeDefinition{
					test.typeDef: test.input,
				},
			}
			result, err := BuildSchema(schema)
			assert.NoError(t, err)
			assert.Equal(t, test.input, result.QueryType)
		})
	}

}

// func TestBuildSchemaErrors(t *testing.T) {
// 	testCases := []struct {
// 		title   string
// 		input   gql.ObjectDefinition
// 		typeDef string
// 	}{
// 		{"queryType", queryType, "Query"},
// 	}

// 	for _, test := range testCases {
// 		t.Run(test.title, func(t *testing.T) {
// 			schema := parser.Schema{
// 				TypeDefinitions: map[string]gql.TypeDefinition{
// 					test.typeDef: test.input,
// 				},
// 			}
// 			result, err := BuildSchema(schema)
// 			assert.NoError(t, err)
// 			assert.Equal(t, test.input, result.QueryType)
// 		})
// 	}

// }
