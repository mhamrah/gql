package gql

type Schema struct {
	Types        map[string]TypeDefinition // [__Type!]!
	QueryType    ObjectDefinition          // __Type!
	MutationType ObjectDefinition          // __Type
	Directives   map[string]Directive      //[__Directive!]!
}

func (s Schema) ValueFromName(field string) interface{} {
	switch field {
	case "types":
		return s.Types
	case "queryType":
		return s.QueryType
	case "mutationType":
		return s.MutationType
	case "directives":
		return s.Directives
	}
	return nil
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

type TypeKind int

const (
	TypeKind_SCALAR TypeKind = iota
	TypeKind_OBJECT
	TypeKind_INTERFACE
	TypeKind_UNION
	TypeKind_ENUM
	TypeKind_INPUT_OBJECT
	TypeKind_LIST
	TypeKind_NON_NULL
)

type DirectiveLocation int

const (
	DirectiveLocation_QUERY DirectiveLocation = iota
	DirectiveLocation_MUTATION
	DirectiveLocation_FIELD
	DirectiveLocation_FRAGMENT_DEFINITION
	DirectiveLocation_FRAGMENT_SPREAD
	DirectiveLocation_INLINE_FRAGMENT
	DirectiveLocation_SCHEMA
	DirectiveLocation_SCALAR
	DirectiveLocation_OBJECT
	DirectiveLocation_FIELD_DEFINITION
	DirectiveLocation_ARGUMENT_DEFINITION
	DirectiveLocation_INTERFACE
	DirectiveLocation_UNION
	DirectiveLocation_ENUM
	DirectiveLocation_ENUM_VALUE
	DirectiveLocation_INPUT_OBJECT
	DirectiveLocation_INPUT_FIELD_DEFINITION
)

/*
type __Schema {
  types: [__Type!]!
  queryType: __Type!
  mutationType: __Type
  directives: [__Directive!]!
}

type __Type {
  kind: __TypeKind!
  name: String
  description: String

  # OBJECT and INTERFACE only
  fields(includeDeprecated: Boolean = false): [__Field!]

  # OBJECT only
  interfaces: [__Type!]

  # INTERFACE and UNION only
  possibleTypes: [__Type!]

  # ENUM only
  enumValues(includeDeprecated: Boolean = false): [__EnumValue!]

  # INPUT_OBJECT only
  inputFields: [__InputValue!]

  # NON_NULL and LIST only
  ofType: __Type
}

type __Field {
  name: String!
  description: String
  args: [__InputValue!]!
  type: __Type!
  isDeprecated: Boolean!
  deprecationReason: String
}

type __InputValue {
  name: String!
  description: String
  type: __Type!
  defaultValue: String
}

type __EnumValue {
  name: String!
  description: String
  isDeprecated: Boolean!
  deprecationReason: String
}

enum __TypeKind {
  SCALAR
  OBJECT
  INTERFACE
  UNION
  ENUM
  INPUT_OBJECT
  LIST
  NON_NULL
}

type __Directive {
  name: String!
  description: String
  locations: [__DirectiveLocation!]!
  args: [__InputValue!]!
}

enum __DirectiveLocation {
  QUERY
  MUTATION
  FIELD
  FRAGMENT_DEFINITION
  FRAGMENT_SPREAD
  INLINE_FRAGMENT
  SCHEMA
  SCALAR
  OBJECT
  FIELD_DEFINITION
  ARGUMENT_DEFINITION
  INTERFACE
  UNION
  ENUM
  ENUM_VALUE
  INPUT_OBJECT
  INPUT_FIELD_DEFINITION
}
*/
