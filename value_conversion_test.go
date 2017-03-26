package gql

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)


func TestValueAsInt(t *testing.T) {
	test := 42
	result, err := ValueAsInt(reflect.ValueOf(test))
	assert.NoError(t, err)
	assert.Equal(t, test, result)

	result, err = ValueAsInt(reflect.ValueOf("42"))
	assert.Error(t, err)
	assert.Equal(t, 0, result)
}

func TestValueAsString(t *testing.T) {
	test := "foobar"
	result, err := ValueAsString(reflect.ValueOf(test))
	assert.NoError(t, err)
	assert.Equal(t, test, result)

	result, err = ValueAsString(reflect.ValueOf(42))
	assert.Error(t, err)
	assert.Equal(t, "", result)
}
