package gql

import (
	"context"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Human struct {
	Name string
}

func (h Human) ValueFromName(field string) interface{} {
	return h.Name
}

func TestSelectableReturnsTypes(t *testing.T) {
	s := gqlService{
		schema: Schema{},
		handlers: map[string]serviceFunc{
			"human": func(ctx context.Context, s Selection) (interface{}, error) {
				return Human{Name: "han"}, nil
			},
			"hello": func(ctx context.Context, s Selection) (interface{}, error) {
				return "world", nil
			},
		},
	}

	input := Selection{
		Field: Field{
			Name: "human",
			Arguments: map[string]Argument{
				"name": Argument{
					Name:  "name",
					Value: reflect.ValueOf("luke"),
				},
			},
			SelectionSet: []Selection{
				Selection{
					Field: Field{
						Name: "name",
					},
				},
			},
		},
	}

	op := Operation{
		SelectionSet: []Selection{
			input,
		},
	}

	opHello := Operation{
		SelectionSet: []Selection{
			Selection{
				Field: Field{
					Name: "hello",
				},
			},
		},
	}

	result, err := s.Operation(context.Background(), op)
	assert.NoError(t, err)
	assert.NotNil(t, result)

	assert.Contains(t, result, "human")

	result, err = s.Operation(context.Background(), opHello)
	assert.NoError(t, err)
	assert.NotNil(t, result)

	assert.Contains(t, result, "hello")
	assert.Equal(t, "world", result["hello"])

}
