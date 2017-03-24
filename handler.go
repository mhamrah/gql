package gql

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/mhamrah/gql/ast"
	"github.com/mhamrah/gql/parser"
	"github.com/kr/pretty"
)

type GqlFunc func(context.Context, ast.Selection) (NamedLookup, error)

type Service interface {
	Handlers() map[string]GqlFunc
	Schema() ast.Schema
}

type NamedLookup interface {
	ValueFromName(string) interface{}
}

type gqlHandler struct {
	handlers map[string]GqlFunc
	schema ast.Schema
}

func NewQueryServer(s Service) http.Handler {
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

	pretty.Println(doc)
	if len(doc.Definitions) == 0 {
		// no defs?
	}

	for _, op := range doc.Definitions {
		for _, selection := range op.Operation.SelectionSet {
			name := selection.Field.Name

			if name == "__schema" {
				h.introspect(w, selection)
				return
			}

			handler, ok := h.handlers[name]
			if !ok {
				http.Error(w, fmt.Sprintf("could not find handler for %s\n", name), http.StatusBadRequest)
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

func renderSelection(w http.ResponseWriter, response NamedLookup, selection ast.Selection) {
		result := make(map[string]interface{})
			for _, s := range selection.Field.SelectionSet {
				f := s.Field

				result[f.Name] = response.ValueFromName(f.Name)
			}


			if err := json.NewEncoder(w).Encode(GraphqlResonse{Data:result}); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
			}
}

func  (h *gqlHandler) introspect(w http.ResponseWriter, selection ast.Selection) {
	json.NewEncoder(w).Encode(h.schema)
}

type GraphqlResonse struct {
	Data map[string]interface{} `json:"data"`
}