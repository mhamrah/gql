package parser

import (
	"testing"

	"github.com/mhamrah/gql"
	"github.com/stretchr/testify/assert"
)

func TestFoo(t *testing.T) {
	doc, err := ParseString(`
		query {
			human(id: 1000) {
			name
		}
		}
	`)
	assert.NoError(t, err)
	assert.NotNil(t, doc)
}

func TestQueryParse(t *testing.T) {
	doc, err := ParseString(simpleQuery)
	assert.NoError(t, err)
	assert.NotNil(t, doc)

	op := doc.Definitions[0].Operation
	assert.NotNil(t, op)
	assert.Equal(t, "FetchLukeQuery", op.Name)
	assert.Equal(t, gql.Query, op.OpType)
	assert.Len(t, op.SelectionSet, 1)
	assert.Equal(t, "human", op.SelectionSet[0].Field.Name)
	assert.Equal(t, "id", op.SelectionSet[0].Field.Arguments["id"].Name)
}

func TestMutationParse(t *testing.T) {
	doc, err := ParseString(mutation)
	assert.NoError(t, err)
	assert.NotNil(t, doc)
	op := doc.Definitions[0].Operation
	assert.NotNil(t, op)
	assert.Empty(t, op.Name)
	assert.Equal(t, gql.Mutation, op.OpType)
	assert.Equal(t, "likeStory", op.SelectionSet[0].Field.Name)
}

func TestKitchenSinkParse(t *testing.T) {
	doc, err := ParseString(kitchenSink)
	assert.NoError(t, err)
	assert.NotNil(t, doc)
	assert.True(t, len(doc.Definitions) > 1)
}

func TestFailureParse(t *testing.T) {
	t.Skip()
	doc, err := ParseString(failure)
	assert.Error(t, err)
	assert.Nil(t, doc)
}

func TestSchemaHelloWorld(t *testing.T) {
	doc, err := ParseString(schemaHelloWorld)
	assert.NoError(t, err)
	assert.NotNil(t, doc)
	schema := doc.Schema
	assert.NotNil(t, schema)
	assert.Nil(t, schema.OperationTypeDefinitions)
	assert.Len(t, schema.TypeDefinitions, 1)
	td := schema.TypeDefinitions["Hello"]
	assert.IsType(t, gql.ObjectDefinition{}, td)
	od, _ := td.(gql.ObjectDefinition)
	assert.Equal(t, "Hello", od.Name)
	assert.Len(t, od.Fields, 1)
	assert.Equal(t, "world", od.Fields[0].Name)

}

func TestSchemaKitchenSink(t *testing.T) {
	doc, err := ParseString(schemaKitchenSink)
	assert.NoError(t, err)
	assert.NotNil(t, doc)
	schema := doc.Schema
	assert.Len(t, schema.OperationTypeDefinitions, 2)
	assert.Len(t, schema.TypeDefinitions, 12)
}
