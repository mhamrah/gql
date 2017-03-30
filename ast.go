package gql

import "reflect"

type Definition struct {
	Operation Operation
	Fragment  Fragment
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
	Arguments map[string]Argument
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

type OperationTypeDefinition struct {
	OpType OpType
	Name   string
}

type FieldDefinition struct {
	Name       string
	Arguments  map[string]InputValueDefinition
	Type       TypeDescription
	Directives []Directive
}

type TypeDescription struct {
	Name  string
	Flags TypeFlags
}
