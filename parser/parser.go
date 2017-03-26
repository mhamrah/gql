//go:generate ragel -Z -G2 -o lexer_gen.go lexer.rl
//go:generate goyacc -o parser_gen.go -v parser.output parser.y

package parser

import (
	"github.com/mhamrah/gql"
)

type Document struct {
	Definitions []gql.Definition
	Schema      Schema
}

type Schema struct {
	OperationTypeDefinitions []gql.OperationTypeDefinition
	TypeDefinitions          map[string]gql.TypeDefinition
}

// ParseString
func ParseString(q string) (*Document, error) {
	return ParseBytes([]byte(q))
}

func ParseBytes(b []byte) (*Document, error) {
	lex := newLexer(b)
	p := yyNewParser()
	e := p.Parse(lex)
	if e != 0 || lex.parseFailed {
		return nil, lex.err
	}
	return &lex.doc, nil
}
