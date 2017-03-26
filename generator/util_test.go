package generator

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTypeMapping(t *testing.T) {
	testCases := []struct {
		graphqlType string
		goType      string
	}{
		{"String", "string"},
	}

	for _, test := range testCases {
		t.Run(fmt.Sprintf("map %v to %v", test.graphqlType, test.goType), func(t *testing.T) {
			result := typeMap(test.graphqlType)
			assert.Equal(t, test.goType, result)
		})
	}
}
