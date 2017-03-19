package gql

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/mhamrah/gql/ast"
	"github.com/mhamrah/gql/parser"
)

type GqlFunc func(context.Context, ast.Selection) (NamedLookup, error)

type Service interface {
	Handlers() map[string]GqlFunc
}

type NamedLookup interface {
	ValueFromName(string) interface{}
}

type gqlHandler struct {
	handlers map[string]GqlFunc
}

func NewQueryServer(s Service) http.Handler {
	return &gqlHandler{handlers: s.Handlers()}
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
		// no defs?
	}

	for _, op := range doc.Definitions {
		for _, selection := range op.Operation.SelectionSet {
			name := selection.Field.Name

			handler, ok := h.handlers[name]
			if !ok {
				http.Error(w, fmt.Sprintf("could not find handler for %s\n", name), http.StatusBadRequest)
				return
			}

			response, err := handler(r.Context(), selection)
			if err != nil {
				http.Error(w, fmt.Sprintf("impl returned error for %v", name), http.StatusBadRequest)
			}

			result := make(map[string]interface{})
			for _, s := range selection.Field.SelectionSet {
				f := s.Field

				result[f.Name] = response.ValueFromName(f.Name)
			}

			if err := json.NewEncoder(w).Encode(result); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
			}
		}
	}

}
