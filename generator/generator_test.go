package generator

import (
	"testing"

	_ "github.com/kr/pretty"
	"github.com/mhamrah/gql/parser"
	"github.com/stretchr/testify/assert"
)

const simpleSchema = `
	schema {
		query: MyQueryRootType
		mutation: MyMutationRootType
	}

	type MyQueryRootType {
		someField: String
	}

	type MyMutationRootType {
		setSomeField(to: String): String
	}
`

const defaultSchema = `
	type Query {
		someField: String
	}`

func TestGenerateDefaultSchema(t *testing.T) {
	ast, err := parser.ParseString(defaultSchema)

	assert.NoError(t, err)
	assert.NotNil(t, ast)

	result, err := Generate(&ast.Schema, "test")
	assert.NoError(t, err)
	assert.NotNil(t, result)
}
