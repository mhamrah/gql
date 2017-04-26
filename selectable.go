package gql

import (
	"context"
	"errors"
	"fmt"
)

type Selectable interface {
	ValueFromName(string) interface{}
}

type HandlerFunc func(context.Context, Selection) (Selectable, error)

type serviceFunc func(context.Context, Selection) (interface{}, error)

type gqlService struct {
	schema   Schema
	handlers map[string]serviceFunc
}

func (s gqlService) Operation(context context.Context, op Operation) (map[string]interface{}, error) {

	result := make(map[string]interface{})

	for _, selection := range op.SelectionSet {
		name := selection.Field.Name

		handler, err := s.getHandler(name)
		if err != nil {
			return nil, err
		}

		response, err := handler(context, selection)
		if err != nil {
			return nil, err
		}

		// todo aliases!

		if len(selection.Field.SelectionSet) == 0 {
			result[name] = response
		} else {
			result[name] = s.selectionToMap(response, selection)
		}

	}

	return result, nil
}

func (s gqlService) getHandler(name string) (serviceFunc, error) {
	switch name {
	case "__schema":
		return s.schemaHandler, nil
	case "__type":
		return s.typeHandler, nil
	default:
		handler, ok := s.handlers[name]
		if !ok {
			return nil, fmt.Errorf("could not find handler for %s\n", name)
		}
		return handler, nil
	}
}

func (s gqlService) Select(context context.Context, selectionSet []Selection) error {

	if len(selectionSet) == 0 {
		return errors.New("no query present")
	}

	// for _, selection := range selectionSet {
	// 	name := selection.Field.Name

	// 	handler, ok := s.schema.Types[name]
	// 	if !ok {
	// 		return errors.New("no handler found")
	// 	}

	// 	result, err := handler(context, selection)
	// 	if err != nil {
	// 		return err
	// 	}

	// 	selectionToMap(result, selection)
	// }
	return nil
}

func (s gqlService) schemaHandler(context context.Context, operation Selection) (interface{}, error) {
	return s.schema, nil
}

func (s gqlService) typeHandler(context context.Context, operation Selection) (interface{}, error) {
	name, err := operation.Field.StringValue("name", "", true)
	if err != nil {
		return nil, err
	}

	success, ok := s.schema.Types[name]
	if !ok {
		return nil, fmt.Errorf("type %v not defined in schema", name)
	}
	return success, nil
}

func (s gqlService) selectionToMap(result interface{}, selection Selection) map[string]interface{} {
	r := make(map[string]interface{})

	for _, s := range selection.Field.SelectionSet {
		f := s.Field
		switch t := result.(type) {
		case Selectable:
			r[f.Name] = t.ValueFromName(f.Name)
		case string:
			r[f.Name] = t
		}

	}
	return r
}

type GraphqlResonse struct {
	Data   map[string]interface{} `json:"data"`
	Errors []Error                `json:"errors,omitempty"`
}

type Location struct {
	Line   int `json:",omitempty"`
	Column int `json:",omitempty"`
}

type Error struct {
	Message  string   `json:"message"`
	Location Location `json:"location,omitempty"`
}

// MLH: AsType() on (gen'd) struct gets you introspection, embed scalar/interface defs on struct.
//then valuebyname gets you a def, and you can blueprint the call. or make valuebyname return a objectdef etc, that has a Scalar() to get a concrete value. the objecdef etc is gen'd.
// experimient with reflect.value to avoid interface()
// you can add a string wrapper to go from string -> selectable.
// fragments
