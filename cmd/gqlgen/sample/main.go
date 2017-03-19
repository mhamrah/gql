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
	http.Handle("/", s)
	http.ListenAndServe(":8080", nil)
}
