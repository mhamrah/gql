package gql

type ParsedSchema struct {
	OperationTypeDefinitions []OperationTypeDefinition
	TypeDefinitions          map[string]TypeDefinition
}

type Schema struct {
	OperationTypeDefinitions []OperationTypeDefinition
	TypeDefinitions          map[string]TypeDefinition
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
