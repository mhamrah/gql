package gql

//go:generate mockgen  -destination=mocks/service.go --package=mocks github.com/mhamrah/gql Service

type Service interface {
	Handlers() map[string]HandlerFunc
	Schema() Schema
}
