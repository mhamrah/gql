package gql

var _ TypeDefinition = InterfaceDefinition{}

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

func (i InterfaceDefinition) TypeKind() TypeKind {
	return TypeKind_INTERFACE
}
