package encoding

import (
	"golang.org/x/net/context"

	"github.com/mhamrah/gql/ast"
	"github.com/yarpc/yarpc-go"
	"github.com/yarpc/yarpc-go/transport"
)

type GqlFunc func(context.Context, yarpc.ReqMeta, *ast.Selection) (*Result, error)

type NamedLookup interface {
	ValueByName(string) interface{}
}
type Result struct {
	Body NamedLookup
	Meta yarpc.ResMeta
}

type Service interface {
	Handlers() map[string]GqlFunc
}

func Register(reg transport.Registry, service Service) {
	reg.Register("graphql", "graphql", &gqlHandler{handlers: service.Handlers()})
}
