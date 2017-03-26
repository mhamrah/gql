package generator

import (
	"bytes"
	"testing"

	_ "github.com/kr/pretty"
	"github.com/mhamrah/gql/parser"
	"github.com/mhamrah/gql/validator"
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

const noQuery = `
	type Foo {
		nana: String
	}
`

// const invalidQuery = `
// 	scalar Query
// `

type verifyGenCode func(t *testing.T, input map[string]bytes.Buffer)

func verifyNotNil(t *testing.T, input map[string]bytes.Buffer) {
	assert.NotNil(t, input)
}

func verifyNil(t *testing.T, input map[string]bytes.Buffer) {
	assert.Nil(t, input)
}

func TestGenerateDefaultSchema(t *testing.T) {
	testCases := []struct {
		title  string
		schema string
		err    error
		verify verifyGenCode
	}{
		{"simple", simpleSchema, nil, verifyNotNil},
		{"default", defaultSchema, nil, verifyNotNil},
		{"noQuery", noQuery, validator.ErrNoQuery, verifyNil},
		//	{"invalidQuery", invalidQuery, validator.ErrInvalidQueryType, verifyNil},
	}

	for _, test := range testCases {
		t.Run(test.title, func(t *testing.T) {
			ast, err := parser.ParseString(test.schema)

			assert.NoError(t, err)
			assert.NotNil(t, ast)

			s, _ := validator.BuildSchema(ast.Schema)
			result, err := Generate(s, "test")
			assert.Equal(t, test.err, err)
			test.verify(t, result)
		})
	}

}
