package ast

import (
	"reflect"
)

type Document struct {
	Definitions []*Operation
}

type OpType int

const (
	Query OpType = iota
	Mutation
)

type Operation struct {
	OpType       OpType
	Name         string
	SelectionSet []*Selection
}

type Selection struct {
	Field *Field
}

type Field struct {
	Name         string
	Arguments    []*Argument
	SelectionSet []*Selection
}

type Argument struct {
	Name  string
	Value reflect.Value
}
