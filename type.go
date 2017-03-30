package gql

import "fmt"

type Type struct {
	Kind          TypeKind
	Name          string
	Description   string
	Fields        []FieldDefinition
	Interfaces    []Type
	PossibleTypes []Type
	EnumValues    EnumDefinition
	InputFields   []InputValueDefinition
	OfType        *Type
}

func (t Type) FieldsF(includeDeprecated bool) Selectable {
	// if includeDeprecated {
	// 	return t.Fields
	// }

	// return t.Fields
	return nil
}

func (t Type) EnumValuesF(includeDeprecated bool) Selectable {
	return t.EnumValues
}

func (t Type) ValueFromName(input string) interface{} {
	switch input {
	case "kind":
		return t.Kind
	case "name":
		return t.Name
	case "description":
		return t.Description
	case "fields":
		return t.FieldsF
	case "interfaces":
		return t.Interfaces
	case "possibleTypes":
		return t.PossibleTypes
	case "enumValues":
		return t.EnumValuesF
	case "inputFields":
		return t.InputFields
	case "ofType":
		return t.OfType
	}

	return fmt.Errorf("field %v is not a member of type %v", input, t.Name)
}
