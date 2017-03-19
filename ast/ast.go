package ast

import (
	"io"
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

type Field struct {
	Name         string
	Arguments    map[string]Argument
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

type Schema struct {
	OperationTypeDefinitions []OperationTypeDefinition
	TypeDefinitions          map[string]TypeDefinition
}

type OperationTypeDefinition struct {
	OpType OpType
	Name   string
}

type TypeDefinition interface {
	Generate(w io.Writer) error
	NamedType() string
}

type ScalarDefinition struct {
	Name       string
	Directives []Directive
}

func (s ScalarDefinition) Generate(w io.Writer) error {
	return nil
}

func (s ScalarDefinition) NamedType() string {
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

func (i InputValueDefinition) Generate(w io.Writer) error {
	return nil
}

func (i InputValueDefinition) NamedType() string {
	return i.Name
}

type InterfaceDefinition struct {
	Name string

	Fields     []FieldDefinition
	Directives []Directive
}

func (i InterfaceDefinition) Generate(w io.Writer) error {
	return nil
}

func (i InterfaceDefinition) NamedType() string {
	return i.Name
}

type UnionDefinition struct {
	Name       string
	Types      []string
	Directives []Directive
}

func (u UnionDefinition) Generate(w io.Writer) error {
	return nil
}

func (u UnionDefinition) NamedType() string {
	return u.Name
}

type EnumDefinition struct {
	Name       string
	Values     []string
	Directives []Directive
}

func (e EnumDefinition) Generate(w io.Writer) error {
	return nil
}

func (e EnumDefinition) NamedType() string {
	return e.Name
}

type InputObjectDefinition struct {
	Name       string
	InputDefs  map[string]InputValueDefinition
	Directives []Directive
}

func (i InputObjectDefinition) Generate(w io.Writer) error {
	return nil
}

func (i InputObjectDefinition) NamedType() string {
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
