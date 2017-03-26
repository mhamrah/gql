package validator

import (
	"errors"

	"github.com/mhamrah/gql"
	"github.com/mhamrah/gql/parser"
)

var (
	ErrNoQuery          = errors.New("Parsed schema has no operation type definitions nor a \"Query\" type")
	ErrInvalidQueryType = errors.New("Query definition is not an ObjectDefinition")
)

func BuildSchema(input parser.Schema) (*gql.Schema, error) {

	var query gql.ObjectDefinition

	if len(input.OperationTypeDefinitions) == 0 {
		q, ok := input.TypeDefinitions["Query"]
		if !ok {
			return nil, ErrNoQuery
		}
		delete(input.TypeDefinitions, "Query")
		od, ok := q.(gql.ObjectDefinition)
		if !ok {
			return nil, ErrInvalidQueryType
		}
		query = od
	}

	return &gql.Schema{QueryType: query, Types: input.TypeDefinitions}, nil
}
