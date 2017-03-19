package main

import (
	"context"
	"net/http"

	"github.com/mhamrah/gql"
	"github.com/mhamrah/gql/ast"
)

/** gen'd **/
type Schema interface {
	Me(ctx context.Context) (User, error)
}

type User struct {
	ID   string
	Name string
}

func (u User) ValueFromName(name string) interface{} {
	switch name {
	case "id":
		return u.ID
	case "name":
		return u.Name
	}
	return nil
}

type service struct {
	impl Schema
}

func New(impl Schema) gql.Service {
	return service{impl: impl}
}

func (s service) Handlers() map[string]gql.GqlFunc {
	return map[string]gql.GqlFunc{
		"me": s.Me,
	}
}

func (s service) Me(ctx context.Context, operation ast.Selection) (gql.NamedLookup, error) {
	success, err := s.impl.Me(ctx)
	if err != nil {
		return nil, err
	}
	return success, nil
}

/* service impl */
type svc struct {
}

func (s *svc) Me(ctx context.Context) (User, error) {
	return User{ID: "1", Name: "foo"}, nil
}

func main() {
	svc := New(&svc{})
	s := gql.NewQueryServer(svc)
	http.Handle("/", s)
	http.ListenAndServe(":8080", nil)
}
