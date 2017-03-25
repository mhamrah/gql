package gql

import (
	"fmt"
	"reflect"
	"strings"

	"text/template"
)

var funcMap = template.FuncMap{
	"title":           strings.Title,
	"typeMap":         typeMap,
	"lower":           strings.ToLower,
	"typeInitializer": typeInitializer,
}

func typeMap(input string) string {
	switch input {
	case "String":
		return "string"
	default:
		return input
	}
}

func typeInitializer(input InputValueDefinition) callInit {
	rhs := "\"\""
	f := ""

	switch input.Type.Name {
	case "String":
		if input.Default.IsValid() {
			rhs = input.Default.String()
		}
		f = "GetString"
	case "Int":
		d := int64(0)
		if input.Default.IsValid() {
			d = input.Default.Int()
		}
		rhs = fmt.Sprintf("%v", d)
		f = "GetInt"
	default:
		rhs = fmt.Sprintf("&%v{}", input.Type.Name)
		//f = "copyStruct"
	}

	return callInit{
		Init:   fmt.Sprintf("%v := %v", input.Name, rhs),
		Setter: fmt.Sprintf("%v, err = gql.%v(input.Value)", input.Name, f),
	}
}

type callInit struct {
	Init   string
	Setter string
}

func GetInt(input reflect.Value) (int, error) {
	if input.Kind() != reflect.Int {
		return 0, fmt.Errorf("%v not an int", input.Kind())
	}
	return int(input.Int()), nil
}

func GetString(input reflect.Value) (string, error) {
	if input.Kind() != reflect.String {
		return "", fmt.Errorf("%v not a string", input.Kind())
	}
	return input.String(), nil
}
