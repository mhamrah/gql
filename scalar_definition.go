package gql

var _ TypeDefinition = ScalarDefinition{}

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

func (s ScalarDefinition) TypeKind() TypeKind {
	return TypeKind_SCALAR
}
