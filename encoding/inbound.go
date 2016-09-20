package encoding

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/mhamrah/gql/parser"
	"github.com/yarpc/yarpc-go/transport"
	"golang.org/x/net/context"
)

type gqlHandler struct {
	handlers map[string]GqlFunc
}

func (g *gqlHandler) Handle(ctx context.Context, opts transport.Options, req *transport.Request, rw transport.ResponseWriter) error {
	bytes, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return err
	}

	doc, err := parser.Parse(bytes)
	if err != nil {
		return err
	}

	reqMeta := FromTransportRequest(req)

	for _, op := range doc.Definitions {
		for _, selection := range op.SelectionSet {
			name := selection.Field.Name
			handler, ok := g.handlers[name]
			if !ok {
				return fmt.Errorf("could not find handler for %s\n", name)
			}

			response, err := handler(ctx, reqMeta, selection)
			if err != nil {
				return err
			}

			result := make(map[string]interface{})

			for _, s := range selection.Field.SelectionSet {
				f := s.Field

				result[f.Name] = response.Body.ValueByName(f.Name)
			}

			if err := json.NewEncoder(rw).Encode(result); err != nil {
				return fmt.Errorf("could not encode result as json")
			}

		}
	}

	return nil
}
