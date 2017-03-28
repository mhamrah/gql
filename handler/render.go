package handler

import (
	"encoding/json"
	"net/http"

	"github.com/mhamrah/gql"
)

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
