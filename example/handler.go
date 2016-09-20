package main

import (
	"github.com/mhamrah/gql/example/starwars"
	yarpc "github.com/yarpc/yarpc-go"
	"golang.org/x/net/context"
)

type handler struct {
	humans map[string]*starwars.Human
}

func (h *handler) Droid(ctx context.Context, reqMeta yarpc.ReqMeta, id *starwars.IdRequest) (*starwars.Droid, yarpc.ResMeta, error) {
	return nil, nil, nil
}
func (h *handler) Hero(ctx context.Context, reqMeta yarpc.ReqMeta, episode starwars.Episode) (*starwars.Character, yarpc.ResMeta, error) {
	return nil, nil, nil
}
func (h *handler) Human(ctx context.Context, reqMeta yarpc.ReqMeta, id *starwars.IdRequest) (*starwars.Human, yarpc.ResMeta, error) {
	return h.humans[id.ID], nil, nil
}

func getHumans() map[string]*starwars.Human {
	luke := &starwars.Human{
		ID:   "1000",
		Name: sp("Luke Skywalker"),
		// Friends:    []int{1002, 1003, 2000, 2001},
		// AppearsIn:  []int{4, 5, 6},
		HomePlanet: sp("Tatooine"),
	}
	vader := &starwars.Human{
		ID:   "1001",
		Name: sp("Darth Vader"),
		// Friends:    []int{1002, 1003, 2000, 2001},
		// AppearsIn:  []int{4, 5, 6},
		HomePlanet: sp("Tatooine"),
	}
	han := &starwars.Human{
		ID:   "1002",
		Name: sp("Han Solo"),
		// Friends:    []int{1002, 1003, 2000, 2001},
		// AppearsIn:  []int{4, 5, 6},
	}
	leia := &starwars.Human{
		ID:   "1003",
		Name: sp("Leia Organa"),
		// Friends:    []int{1002, 1003, 2000, 2001},
		// AppearsIn:  []int{4, 5, 6},
		HomePlanet: sp("Alderaan"),
	}
	tarkin := &starwars.Human{
		ID:   "1004",
		Name: sp("Wilhuff Tarkin"),
		// Friends:    []int{1002, 1003, 2000, 2001},
		// AppearsIn:  []int{4, 5, 6},
	}

	return map[string]*starwars.Human{luke.ID: luke, vader.ID: vader, han.ID: han, leia.ID: leia, tarkin.ID: tarkin}
}

func sp(input string) *string {
	var out string
	out = input
	return &out
}
