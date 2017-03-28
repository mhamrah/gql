package gql

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringValue(t *testing.T) {
	input := Field{
		Arguments: map[string]Argument{
			"okay": Argument{
				Name:  "okay",
				Value: reflect.ValueOf("bar"),
			},
			"anint": Argument{
				Name:  "anint",
				Value: reflect.ValueOf(42),
			},
		},
	}

	result, err := input.StringValue("okay", "zzz", true)
	assert.NoError(t, err)
	assert.Equal(t, "bar", result)

	result, err = input.StringValue("okay", "zzz", false)
	assert.NoError(t, err)
	assert.Equal(t, "bar", result)

	result, err = input.StringValue("notpresent", "zzz", false)
	assert.NoError(t, err)
	assert.Equal(t, "zzz", result)

	result, err = input.StringValue("notpresent", "zzz", true)
	assert.Error(t, err)
	assert.Equal(t, "", result)

	result, err = input.StringValue("anint", "zzz", true)
	assert.Error(t, err)
	assert.Equal(t, "", result)

}
