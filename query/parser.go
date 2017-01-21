//go:generate ragel -Z -G2 -o graphql_lexer_gen.go graphql_lexer.rl
//go:generate go tool yacc -o graphql_parser_gen.go -v graphql_parser.output graphql_parser.y

package query

// ParseString
func ParseString(q string) (*Document, error) {
	return ParseBytes([]byte(q))
}

func ParseBytes(b []byte) (*Document, error) {
	lex := newLexer(b)
	e := yyParse(lex)
	if e != 0 || lex.parseFailed {
		return nil, lex.err
	}
	return &lex.doc, nil
}
