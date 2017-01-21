package query

import (
	"reflect"
)

type Document struct {
	Definitions []Definition
}

type OpType int

const (
	Query OpType = iota
	Mutation
	Subscription
)

type Definition struct {
	operation Operation
	fragment  Fragment
}

type Operation struct {
	OpType       OpType
	Name         string
	SelectionSet []Selection
	Variables    []Variable
	Directives   []Directive
}

type Selection struct {
	Field          Field
	FragmentSpread FragmentSpread
	InlineFragment Fragment
}

type Field struct {
	Name         string
	Arguments    []Argument
	SelectionSet []Selection
	Directives   []Directive
	Alias        string
	Ix           int
}

type Argument struct {
	Name  string
	Value reflect.Value
}

type Fragment struct {
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
