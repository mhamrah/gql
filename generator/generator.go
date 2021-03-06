package generator

import (
	"bytes"
	"fmt"

	"github.com/mhamrah/gql"
	"github.com/mhamrah/gql/validator"
)

const generatedHeader = "// @generated\n// generated by gql\n\n"

func Generate(schema *gql.Schema, packageName string) (map[string]bytes.Buffer, error) {
	files := make(map[string]bytes.Buffer)

	if schema == nil {
		return nil, validator.ErrNoQuery
	}

	op, err := writeOperation(schema.QueryType, packageName)
	if err != nil {
		return nil, err
	}

	files["service.go"] = op

	if len(schema.Types) > 0 {
		types, err := writeTypes(schema.Types, packageName)
		if err != nil {
			return nil, err
		}
		files["types.go"] = types
	}

	return files, nil
}

func writeOperation(od gql.ObjectDefinition, packageName string) (bytes.Buffer, error) {
	var op bytes.Buffer
	writeHeader(&op, packageName)
	writeImports(&op)
	if err := GenerateService(&op, od); err != nil {
		return op, err
	}
	return op, nil
}

func writeHeader(b *bytes.Buffer, packageName string) {
	b.WriteString(generatedHeader)
	b.WriteString(fmt.Sprintf("package %v\n\n", packageName))
}

func writeImports(b *bytes.Buffer) {
	b.WriteString("import (\n")
	b.WriteString("\"context\"\n")
	b.WriteString("\"github.com/mhamrah/gql\"\n")
	b.WriteString(")\n\n")
}

func writeTypes(types map[string]gql.TypeDefinition, packageName string) (bytes.Buffer, error) {
	var b bytes.Buffer

	writeHeader(&b, packageName)
	for _, t := range types {
		GenerateTypeDefinition(&b, t)
	}

	return b, nil
}
