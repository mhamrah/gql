package gql

var _ TypeDefinition = InputObjectDefinition{}

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

func (i InputObjectDefinition) TypeKind() TypeKind {
	return TypeKind_INPUT_OBJECT
}
