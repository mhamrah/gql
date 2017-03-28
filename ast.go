package gql

import "reflect"

type Selectable interface {
	ValueFromName(string) interface{}
}

type OpType int

const (
	Query OpType = iota
	Mutation
	Subscription
)

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

type TypeDefinition interface {
	Selectable
	TypeName() string
}

type ScalarDefinition struct {
	Name       string
	Directives []Directive
}

func (s ScalarDefinition) TypeName() string {
	return s.Name
}

func (s ScalarDefinition) ValueFromName(input string) interface{} {
	return s.Name
}

type FieldDefinition struct {
	Name       string
	Arguments  map[string]InputValueDefinition
	Type       TypeDescription
	Directives []Directive
}

type InputValueDefinition struct {
	Name       string
	Type       TypeDescription
	Default    reflect.Value
	Directives []Directive
}

func (i InputValueDefinition) TypeName() string {
	return i.Name
}

func (i InputValueDefinition) ValueFromName(input string) interface{} {
	return i.Name
}

type InterfaceDefinition struct {
	Name string

	Fields     []FieldDefinition
	Directives []Directive
}

func (i InterfaceDefinition) TypeName() string {
	return i.Name
}

func (i InterfaceDefinition) ValueFromName(input string) interface{} {
	return i.Name
}

type UnionDefinition struct {
	Name       string
	Types      []string
	Directives []Directive
}

func (u UnionDefinition) TypeName() string {
	return u.Name
}

func (u UnionDefinition) ValueFromName(input string) interface{} {
	return u.Name
}

type EnumDefinition struct {
	Name       string
	Values     []string
	Directives []Directive
}

func (e EnumDefinition) TypeName() string {
	return e.Name
}

func (e EnumDefinition) ValueFromName(input string) interface{} {
	return e.Name
}

type InputObjectDefinition struct {
	Name       string
	InputDefs  map[string]InputValueDefinition
	Directives []Directive
}

func (i InputObjectDefinition) TypeName() string {
	return i.Name
}

func (i InputObjectDefinition) ValueFromName(input string) interface{} {
	return i.Name
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

func IsRequired(input TypeFlags) bool {
	return input&Required != 0
}
