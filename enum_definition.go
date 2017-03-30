package gql

var _ TypeDefinition = EnumDefinition{}

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

func (e EnumDefinition) TypeKind() TypeKind {
	return TypeKind_ENUM
}
