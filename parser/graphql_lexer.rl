package parser

import (
        "fmt"
        "reflect"
        "strconv"
        "github.com/mhamrah/gql/ast"
)

%%{
    machine graphql;
    write data;
    access lex.;
    variable p lex.p;
    variable pe lex.pe;
}%%

type lexer struct {
    doc *ast.Document
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
        # Symbols are sent to the parser as-is.
        symbol = [\{\}\(\):];

        main := |*
           'query' => { tok = QUERY; fbreak;};
           'mutation' => { tok = MUTATION; fbreak;};
           'true' | 'false' => {
               out.val = reflect.ValueOf(string(lex.data[lex.ts:lex.te]) == "true");
               tok = VALUE
               fbreak;
           };
            name => {
                out.str = string(lex.data[lex.ts:lex.te]);
                tok = NAME;
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

            space;
            any => { fmt.Println("any!", string(lex.data[lex.ts:lex.te])) };
        *|;

         write exec;
    }%%

    return tok;
}

func (lex *lexer) Error(e string) {
    fmt.Println("errorx:", e)
    fmt.Println("p", string(lex.data[lex.p:lex.pe]))
    fmt.Println("t", string(lex.data[lex.ts:lex.te]))
}
