package gqlqueryserver

import (
	"fmt"
	"reflect"

	"github.com/mhamrah/gql/ast"
	"github.com/mhamrah/gql/encoding"
	"github.com/mhamrah/gql/example/starwars"
	"github.com/mhamrah/gql/example/starwars/yarpc/queryserver"
	yarpc "github.com/yarpc/yarpc-go"
	"golang.org/x/net/context"
)

type service struct {
	impl queryserver.Interface
}

func New(impl queryserver.Interface) encoding.Service {
	return &service{impl: impl}
}

func (s service) Handlers() map[string]encoding.GqlFunc {
	return map[string]encoding.GqlFunc{
		"Droid": s.Droid,
		"Hero":  s.Hero,
		"Human": s.Human}
}

func (s service) Droid(ctx context.Context, reqMeta yarpc.ReqMeta, selection *ast.Selection) (*encoding.Result, error) {
	fmt.Println("!droid")
	return nil, nil
}

func (s service) Hero(ctx context.Context, reqMeta yarpc.ReqMeta, selection *ast.Selection) (*encoding.Result, error) {
	fmt.Println("!hero")
	return nil, nil
}

func (s service) Human(ctx context.Context, reqMeta yarpc.ReqMeta, selection *ast.Selection) (*encoding.Result, error) {
	args := &starwars.IdRequest{}

	r := reflect.ValueOf(args).Elem()
	for _, a := range selection.Field.Arguments {
		f := r.FieldByName(a.Name)
		if f.IsValid() && f.CanSet() && f.Type() == a.Value.Type() {
			f.Set(a.Value)
		} else {
			return nil, fmt.Errorf("cannot set field %s", a.Name)
		}
	}

	success, resMeta, err := s.impl.Human(ctx, reqMeta, args)
	if err != nil {
		return nil, err
	}

	return &encoding.Result{Body: success, Meta: resMeta}, nil
}
