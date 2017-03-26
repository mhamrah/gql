package generator

import (
	"fmt"
	"strings"

	"github.com/mhamrah/gql"

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

func typeInitializer(input gql.InputValueDefinition) callInit {
	rhs := "\"\""
	f := ""

	switch input.Type.Name {
	case "String":
		if input.Default.IsValid() {
			rhs = input.Default.String()
		}
		f = "ValueAsString"
	case "Int":
		d := int64(0)
		if input.Default.IsValid() {
			d = input.Default.Int()
		}
		rhs = fmt.Sprintf("%v", d)
		f = "ValueAsInt"
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
