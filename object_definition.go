package gql

var _ TypeDefinition = ObjectDefinition{}

type ObjectDefinition struct {
	Name        string
	Implements  []string
	Fields      []FieldDefinition
	Directives  []Directive
	IsOperation bool
}

func (o ObjectDefinition) TypeName() string {
	return o.Name
}
