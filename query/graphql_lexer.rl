package query

import (
        "fmt"
        "reflect"
        "strconv"
)

%%{
    machine graphql;
    write data;
    access lex.;
    variable p lex.p;
    variable pe lex.pe;
}%%

type lexer struct {
    doc Document
    err error
    parseFailed bool

    data []byte
    p, pe, cs int
    ts, te, act int
}

func newLexer(data []byte) *lexer {
    lex := &lexer{
        data: data,
        pe: len(data),
        parseFailed: false,
    }
    %% write init;
    return lex
}

func (lex *lexer) Lex(out *yySymType) int {
    eof := lex.pe
    tok := 0
    %%{
        name = [_A-Za-z][_0-9A-Za-z]*;
        value = ["_A-Za-z]["_0-9A-Za-z]*;
        stringValue = ('"' ([^"\n\\] | '\\' any)* '"');
        intValue = ('-'? [1-9] digit*);
        var = '$'[_0-9A-Za-z]+;

        # Symbols are sent to the parser as-is.
        symbol = [\{\}\(\):=\[\]@!\|];

        main := |*
            'directive' => { tok = DIRECTIVE; fbreak; };
            'enum' => { tok = ENUM; fbreak; };
            'extend' => { tok = EXTEND; fbreak; };
            'false' => { tok = FALSE; fbreak; };
            'fragment' => { tok = FRAGMENT; fbreak; };
            'implements' => { tok = IMPLEMENTS; fbreak; };
            'input' => { tok = INPUT; fbreak; };
            'interface' => { tok = INTERFACE; fbreak; };
            'mutation' => { tok = MUTATION; fbreak;};
            'null' => { tok = NULL; fbreak; };
            'query' => { tok = QUERY; fbreak;};
            'on' => { tok = ON; fbreak; };
            'scalar' => { tok = SCALAR; fbreak; };
            'schema' => { tok = SCHEMA; fbreak; };
            'subscription' => { tok = SUBSCRIPTION; fbreak; };
            'true' => { tok = TRUE; fbreak; };
            'type' => { tok = TYPE; fbreak; };
            'union' => { tok = UNION; fbreak; };
            '...' => {
                tok = SPREAD;
                fbreak;
            };
            'on' => {
                tok = ON;
                fbreak;
            };
            name => {
                out.str = string(lex.data[lex.ts:lex.te]);
                tok = IDENTIFIER;
                fbreak;
            };
            intValue => {
                i, err := strconv.Atoi(string(lex.data[lex.ts:lex.te]));
                if err != nil {
                    lex.Error("can't convert int")
                }
                out.val = reflect.ValueOf(i);
                tok = VALUE;
                fbreak;
            };
            stringValue => {
                out.val = reflect.ValueOf(string(lex.data[lex.ts+1:lex.te-1]));
                tok = VALUE
                fbreak;
            };
            symbol => {
                tok = int(lex.data[lex.ts])
                fbreak;
            };
            var => {
                out.str = string(lex.data[lex.ts+1:lex.te]);
                tok = VARIABLE
                fbreak;
            };
            space;
            any => { };
        *|;

         write exec;
    }%%

    return tok;
}

func (lex *lexer) Error(e string) {
    fmt.Println("errorx:", e)
    fmt.Println("t", string(lex.data[lex.ts:lex.te]))
    fmt.Println("p", string(lex.data[lex.p:lex.pe]))
}
