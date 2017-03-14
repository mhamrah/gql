package query

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQueryParse(t *testing.T) {
	doc, err := ParseString(simpleQuery)
	assert.NoError(t, err)
	assert.NotNil(t, doc)

	op := doc.Definitions[0].operation
	assert.NotNil(t, op)
	assert.Equal(t, "FetchLukeQuery", op.Name)
	assert.Equal(t, Query, op.OpType)
	assert.Len(t, op.SelectionSet, 1)
	assert.Equal(t, "human", op.SelectionSet[0].Field.Name)
	assert.Equal(t, "id", op.SelectionSet[0].Field.Arguments[0].Name)
}

func TestMutationParse(t *testing.T) {
	doc, err := ParseString(mutation)
	assert.NoError(t, err)
	assert.NotNil(t, doc)
	op := doc.Definitions[0].operation
	assert.NotNil(t, op)
	assert.Empty(t, op.Name)
	assert.Equal(t, Mutation, op.OpType)
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

func TestSchemaHelloWord(t *testing.T) {
	schema, err := ParseString(schemaHelloWorld)
	assert.NoError(t, err)
	assert.NotNil(t, schema)
	assert.NotNil(t, schema.Schema)

}

func TestSchemaKitchenSink(t *testing.T) {
	doc, err := ParseString(schemaKitchenSink)
	assert.NoError(t, err)
	assert.NotNil(t, doc)
	schema := doc.Schema
	assert.Len(t, schema.OperationTypeDefinitions, 2)
	assert.Len(t, schema.TypeDefinitions, 12)
}
