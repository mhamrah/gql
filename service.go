package gql

import "context"

//go:generate mockgen  -destination=mocks/service.go --package=mocks github.com/mhamrah/gql Service

type HandlerFunc func(context.Context, Selection) (Selectable, error)

type Service interface {
	Handlers() map[string]HandlerFunc
	Schema() Schema
}
