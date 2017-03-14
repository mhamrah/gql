package query

import (
	"reflect"
)

type Document struct {
	Definitions []Definition
	Schema      Schema
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
	Type    TypeDescription
	Default reflect.Value
}

type FragmentSpread struct {
	Name       string
	Directives []Directive
}

type ObjectField struct {
	Key   string
	Value reflect.Value
}

type Schema struct {
	OperationTypeDefinitions []OperationTypeDefinition
	TypeDefinitions          []TypeDefinition
}

type OperationTypeDefinition struct {
	OpType OpType
	Name   string
}

type TypeDefinition interface {
}

type ScalarDefinition struct {
	Name string
	Directives []Directive
}

type ObjectDefinition struct {
	Name       string
	Implements []string
	Fields []FieldDefinition
	Directives []Directive
}

type FieldDefinition struct {
	Name      string
	Arguments []InputValueDefinition
	Type      TypeDescription
	Directives []Directive
}

type InputValueDefinition struct {
	Name    string
	Type    TypeDescription
	Default reflect.Value
	Directives []Directive
}

type InterfaceDefinition struct {
	Name string

	Fields []FieldDefinition
	Directives []Directive

}

type UnionDefinition struct {
	Name  string
	Types []string
	Directives []Directive
}

type EnumDefinition struct {
	Name   string
	Values []string
	Directives []Directive
}

type InputObjectDefinition struct {
	Name      string
	InputDefs []InputValueDefinition
	Directives []Directive
}

type TypeFlags byte

const (
	List TypeFlags = 1 << iota
	Required
)

type TypeDescription struct {
	Name  string
	Flags TypeFlags
}
