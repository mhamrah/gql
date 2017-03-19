package main

import (
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

const schema = `
type Query {
  me: User
}

type User {
  id: ID
  name: String
}`

const expected = `
	foo
`

func TestGenerate(t *testing.T) {
	err := Generate(os.TempDir(), "test", strings.NewReader(schema))
	assert.NoError(t, err)
}
