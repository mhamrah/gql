package ast

import (
	"reflect"
)

type Document struct {
	Definitions []Operation
}

type OpType int

const (
	Query OpType = iota
	Mutation
)

type Operation struct {
	OpType       OpType
	Name         string
	SelectionSet []Selection
}

type Selection struct {
	Field *Field
}

type Field struct {
	Name         string
	Arguments    []Argument
	SelectionSet []Selection
	Ix           int
}

type Argument struct {
	Name  string
	Value reflect.Value
}

type FragmentDefinition struct {
	Name          string
	TypeCondition string
	Directives    []Directive
	SelectionSet  []Selection
}

type Directive struct {
	Name      string
	Arguments []Argument
}

type Variable struct {
	Name    string
	Type    string
	Default reflect.Value
}

type FragmentSpread struct {
	Name       string
	Directives []Directive
}
