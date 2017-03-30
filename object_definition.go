package gql

import "fmt"

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

func (o ObjectDefinition) ValueFromName(input string) interface{} {
	switch input {
	case "kind":
		return o.TypeKind()
	case "name":
		return o.TypeName()
	case "description":
		return ""
	case "interfaces":
		return nil
	case "possibleTypes":
		return nil
	case "enumValues":
		return nil
	case "inputFields":
		return nil
	case "ofType":
		return nil
	}
	return fmt.Errorf("field %v is not a member of type %v", input, o.Name)
}

func (o ObjectDefinition) TypeKind() TypeKind {
	return TypeKind_OBJECT
}

func (o ObjectDefinition) AsType(isRequired, isList bool) Type {
	return Type{}
}
