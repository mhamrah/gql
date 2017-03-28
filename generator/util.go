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

func typeInitializer(input gql.InputValueDefinition) string {
	rhs := "\"\""
	funcName := ""

	switch input.Type.Name {
	case "String":
		if input.Default.IsValid() {
			rhs = input.Default.String()
		}
		funcName = "StringValue"
	case "Int":
		d := int64(0)
		if input.Default.IsValid() {
			d = input.Default.Int()
		}
		rhs = fmt.Sprintf("%v", d)
		funcName = "IntValue"
	default:
		rhs = fmt.Sprintf("&%v{}", input.Type.Name)
		//f = "copyStruct"
	}

	return fmt.Sprintf("%v, err := operation.Field.%v(\"%v\", %v, %v)", input.Name, funcName, input.Name, rhs, gql.IsRequired(input.Type.Flags))

}
