package gql

type Selectable interface {
	ValueFromName(string) interface{}
}

type TypeDefinition interface {
	Selectable
	TypeName() string
	TypeKind() TypeKind
}

type OpType int

const (
	Query OpType = iota
	Mutation
	Subscription
)

type TypeFlags byte

const (
	List TypeFlags = 1 << iota
	Required
)

func IsRequired(input TypeFlags) bool {
	return input&Required != 0
}
