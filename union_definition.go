package gql

var _ TypeDefinition = UnionDefinition{}

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

func (u UnionDefinition) TypeKind() TypeKind {
	return TypeKind_UNION
}
