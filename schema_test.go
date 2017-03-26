package gql

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuiltinSchema(t *testing.T) {
	builtins := BuiltinDefinitions()

	assert.NotNil(t, builtins["String"])

}
