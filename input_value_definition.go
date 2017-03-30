package gql

import "reflect"

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
