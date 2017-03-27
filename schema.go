package gql

type Schema struct {
	Types        map[string]TypeDefinition // [__Type!]!
	QueryType    ObjectDefinition          // __Type!
	MutationType ObjectDefinition          // __Type
	Directives   map[string]Directive      //[__Directive!]!
}

func (s Schema) ValueFromName(field string) interface{} {
	return "foo"
}

func BuiltinDefinitions() map[string]TypeDefinition {
	builtin := make(map[string]TypeDefinition)

	builtin["Int"] = ScalarDefinition{Name: "Int"}
	builtin["Float"] = ScalarDefinition{Name: "Float"}
	builtin["String"] = ScalarDefinition{Name: "Float"}
	builtin["Boolean"] = ScalarDefinition{Name: "Boolean"}
	builtin["ID"] = ScalarDefinition{Name: "ID"}
	builtin["Float"] = ScalarDefinition{Name: "Float"}

	return builtin
}
