package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/mhamrah/gql"
	"github.com/mhamrah/gql/parser"
)

type gqlHandler struct {
	handlers map[string]gql.HandlerFunc
}

func NewGqlHandler(s gql.Service) http.Handler {
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
		http.Error(w, "no query present", http.StatusBadRequest)
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

			renderSelection(w, response, selection)
		}
	}
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
	Data map[string]interface{} `json:"data"`
}
