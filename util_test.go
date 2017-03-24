package gql

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetValue(t *testing.T) {
	dest := "bar"
	foo := reflect.ValueOf("foo")
	bar := reflect.ValueOf(dest)

	result, err := setValue(foo, bar)

	assert.NoError(t, err)
	assert.Equal(t, result.String(), foo.String())
}
