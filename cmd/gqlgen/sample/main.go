package main

import (
	"context"
	"net/http"

	"github.com/mhamrah/gql"
	"github.com/mhamrah/gql/cmd/gqlgen/sample/sample"
)

type svc struct {
}

func (s *svc) Human(ctx context.Context, id string) (sample.Human, error) {
	return sample.Human{Id: id, Name: "foo", HomePlanet: "Tattoine"}, nil
}

func main() {
	svc := sample.New(&svc{})
	s := gql.NewQueryServer(svc)
	fs := http.FileServer(http.Dir("public"))

	http.Handle("/graphql", s)
	http.Handle("/", fs)
	http.ListenAndServe(":8080", nil)
}
