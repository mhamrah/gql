package gql

type FieldDefinition struct {
	Name       string
	Arguments  map[string]InputValueDefinition
	Type       TypeDescription
	Directives []Directive
}
