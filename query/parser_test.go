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
