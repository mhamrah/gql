package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/mhamrah/gql"
	"github.com/mhamrah/gql/parser"
)

type gqlHandler struct {
	handlers map[string]gql.HandlerFunc
	schema   gql.Schema
}

func NewGqlHandler(s gql.Service) http.Handler {
	return &gqlHandler{handlers: s.Handlers(), schema: s.Schema()}
}

func (h *gqlHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	doc, err := parser.ParseBytes(b)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if len(doc.Definitions) == 0 {
		http.Error(w, "no query present", http.StatusBadRequest)
	}

	for _, op := range doc.Definitions {
		for _, selection := range op.Operation.SelectionSet {
			name := selection.Field.Name

			handler, err := h.getHandler(name)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			response, err := handler(r.Context(), selection)
			if err != nil {
				http.Error(w, fmt.Sprintf("impl returned error for %v", name), http.StatusBadRequest)
			}

			renderSelection(w, response, selection)
		}
	}
}

func (h *gqlHandler) getHandler(name string) (gql.HandlerFunc, error) {
	switch name {
	case "__schema":
		return h.schemaHandler, nil
	case "__type":
		return h.typeHandler, nil
	default:
		handler, ok := h.handlers[name]
		if !ok {
			return nil, fmt.Errorf("could not find handler for %s\n", name)
		}
		return handler, nil
	}
}

func (h *gqlHandler) schemaHandler(context context.Context, operation gql.Selection) (gql.Selectable, error) {
	return h.schema, nil
}

func (h *gqlHandler) typeHandler(context context.Context, operation gql.Selection) (gql.Selectable, error) {
	name := ""
	name, err := argumentAsString(operation.Field, "name", name)
	if err != nil {
		return nil, err
	}

	success, ok := h.schema.Types[name]
	if !ok {
		return nil, fmt.Errorf("type %v not defined in schema", name)
	}
	return success, nil
}

func argumentAsString(field gql.Field, arg string, current string) (string, error) {
	result := current
	if input, ok := field.Arguments[arg]; ok {
		var err error
		if !input.Value.IsValid() {
			return result, fmt.Errorf("%v does not contain a valid value", arg)
		}
		result, err = gql.ValueAsString(input.Value)
		if err != nil {
			return result, err
		}
	}
	return result, nil
}

func renderSelection(w http.ResponseWriter, result gql.Selectable, selection gql.Selection) {
	r := make(map[string]interface{})
	for _, s := range selection.Field.SelectionSet {
		f := s.Field

		r[f.Name] = result.ValueFromName(f.Name)
	}

	if err := json.NewEncoder(w).Encode(GraphqlResonse{Data: r}); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
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
