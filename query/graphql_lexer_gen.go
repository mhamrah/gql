
//line graphql_lexer.rl:1
package query

import (
        "fmt"
        "reflect"
        "strconv"
)


//line graphql_lexer_gen.go:13
const graphql_start int = 2
const graphql_first_final int = 2
const graphql_error int = -1

const graphql_en_main int = 2


//line graphql_lexer.rl:15


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
    
//line graphql_lexer_gen.go:41
	{
	 lex.cs = graphql_start
	 lex.ts = 0
	 lex.te = 0
	 lex.act = 0
	}

//line graphql_lexer.rl:34
    return lex
}

func (lex *lexer) Lex(out *yySymType) int {
    eof := lex.pe
    tok := 0
    
//line graphql_lexer_gen.go:57
	{
	if ( lex.p) == ( lex.pe) {
		goto _test_eof
	}
	switch  lex.cs {
	case 2:
		goto st_case_2
	case 3:
		goto st_case_3
	case 0:
		goto st_case_0
	case 1:
		goto st_case_1
	case 4:
		goto st_case_4
	case 5:
		goto st_case_5
	case 6:
		goto st_case_6
	case 7:
		goto st_case_7
	case 8:
		goto st_case_8
	case 9:
		goto st_case_9
	case 10:
		goto st_case_10
	case 11:
		goto st_case_11
	case 12:
		goto st_case_12
	case 13:
		goto st_case_13
	case 14:
		goto st_case_14
	case 15:
		goto st_case_15
	case 16:
		goto st_case_16
	case 17:
		goto st_case_17
	case 18:
		goto st_case_18
	case 19:
		goto st_case_19
	case 20:
		goto st_case_20
	case 21:
		goto st_case_21
	case 22:
		goto st_case_22
	case 23:
		goto st_case_23
	}
	goto st_out
tr0:
//line graphql_lexer.rl:81
( lex.p) = ( lex.te) - 1
{ }
	goto st2
tr2:
//line graphql_lexer.rl:70
 lex.te = ( lex.p)+1
{
                out.val = reflect.ValueOf(string(lex.data[lex.ts+1:lex.te-1]));
                tok = VALUE
                {( lex.p)++;  lex.cs = 2; goto _out }
            }
	goto st2
tr4:
//line graphql_lexer.rl:81
 lex.te = ( lex.p)+1
{ }
	goto st2
tr5:
//line graphql_lexer.rl:80
 lex.te = ( lex.p)+1

	goto st2
tr7:
//line graphql_lexer.rl:75
 lex.te = ( lex.p)+1
{
                tok = int(lex.data[lex.ts])
                {( lex.p)++;  lex.cs = 2; goto _out }
            }
	goto st2
tr15:
//line graphql_lexer.rl:81
 lex.te = ( lex.p)
( lex.p)--
{ }
	goto st2
tr16:
//line graphql_lexer.rl:61
 lex.te = ( lex.p)
( lex.p)--
{
                i, err := strconv.Atoi(string(lex.data[lex.ts:lex.te]));
                if err != nil {
                    lex.Error("can't convert int")
                }
                out.val = reflect.ValueOf(i);
                tok = VALUE;
                {( lex.p)++;  lex.cs = 2; goto _out }
            }
	goto st2
tr17:
//line NONE:1
	switch  lex.act {
	case 1:
	{( lex.p) = ( lex.te) - 1
 tok = QUERY; {( lex.p)++;  lex.cs = 2; goto _out }}
	case 2:
	{( lex.p) = ( lex.te) - 1
 tok = MUTATION; {( lex.p)++;  lex.cs = 2; goto _out }}
	case 3:
	{( lex.p) = ( lex.te) - 1

               out.val = reflect.ValueOf(string(lex.data[lex.ts:lex.te]) == "true");
               tok = VALUE
               {( lex.p)++;  lex.cs = 2; goto _out }
           }
	case 4:
	{( lex.p) = ( lex.te) - 1

                out.str = string(lex.data[lex.ts:lex.te]);
                tok = NAME;
                {( lex.p)++;  lex.cs = 2; goto _out }
            }
	}
	
	goto st2
tr18:
//line graphql_lexer.rl:56
 lex.te = ( lex.p)
( lex.p)--
{
                out.str = string(lex.data[lex.ts:lex.te]);
                tok = NAME;
                {( lex.p)++;  lex.cs = 2; goto _out }
            }
	goto st2
	st2:
//line NONE:1
 lex.ts = 0

		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof2
		}
	st_case_2:
//line NONE:1
 lex.ts = ( lex.p)

//line graphql_lexer_gen.go:212
		switch  lex.data[( lex.p)] {
		case 32:
			goto tr5
		case 34:
			goto tr6
		case 45:
			goto st4
		case 58:
			goto tr7
		case 95:
			goto tr10
		case 102:
			goto st7
		case 109:
			goto st11
		case 113:
			goto st18
		case 116:
			goto st22
		case 123:
			goto tr7
		case 125:
			goto tr7
		}
		switch {
		case  lex.data[( lex.p)] < 49:
			switch {
			case  lex.data[( lex.p)] > 13:
				if 40 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 41 {
					goto tr7
				}
			case  lex.data[( lex.p)] >= 9:
				goto tr5
			}
		case  lex.data[( lex.p)] > 57:
			switch {
			case  lex.data[( lex.p)] > 90:
				if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
					goto tr10
				}
			case  lex.data[( lex.p)] >= 65:
				goto tr10
			}
		default:
			goto st5
		}
		goto tr4
tr6:
//line NONE:1
 lex.te = ( lex.p)+1

	goto st3
	st3:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof3
		}
	st_case_3:
//line graphql_lexer_gen.go:270
		switch  lex.data[( lex.p)] {
		case 10:
			goto tr15
		case 34:
			goto tr2
		case 92:
			goto st1
		}
		goto st0
	st0:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof0
		}
	st_case_0:
		switch  lex.data[( lex.p)] {
		case 10:
			goto tr0
		case 34:
			goto tr2
		case 92:
			goto st1
		}
		goto st0
	st1:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof1
		}
	st_case_1:
		goto st0
	st4:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof4
		}
	st_case_4:
		if 49 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
			goto st5
		}
		goto tr15
	st5:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof5
		}
	st_case_5:
		if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
			goto st5
		}
		goto tr16
tr10:
//line NONE:1
 lex.te = ( lex.p)+1

//line graphql_lexer.rl:56
 lex.act = 4;
	goto st6
tr22:
//line NONE:1
 lex.te = ( lex.p)+1

//line graphql_lexer.rl:51
 lex.act = 3;
	goto st6
tr29:
//line NONE:1
 lex.te = ( lex.p)+1

//line graphql_lexer.rl:50
 lex.act = 2;
	goto st6
tr33:
//line NONE:1
 lex.te = ( lex.p)+1

//line graphql_lexer.rl:49
 lex.act = 1;
	goto st6
	st6:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof6
		}
	st_case_6:
//line graphql_lexer_gen.go:351
		if  lex.data[( lex.p)] == 95 {
			goto tr10
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr10
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr10
			}
		default:
			goto tr10
		}
		goto tr17
	st7:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof7
		}
	st_case_7:
		switch  lex.data[( lex.p)] {
		case 95:
			goto tr10
		case 97:
			goto st8
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr10
			}
		case  lex.data[( lex.p)] > 90:
			if 98 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr10
			}
		default:
			goto tr10
		}
		goto tr18
	st8:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof8
		}
	st_case_8:
		switch  lex.data[( lex.p)] {
		case 95:
			goto tr10
		case 108:
			goto st9
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr10
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr10
			}
		default:
			goto tr10
		}
		goto tr18
	st9:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof9
		}
	st_case_9:
		switch  lex.data[( lex.p)] {
		case 95:
			goto tr10
		case 115:
			goto st10
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr10
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr10
			}
		default:
			goto tr10
		}
		goto tr18
	st10:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof10
		}
	st_case_10:
		switch  lex.data[( lex.p)] {
		case 95:
			goto tr10
		case 101:
			goto tr22
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr10
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr10
			}
		default:
			goto tr10
		}
		goto tr18
	st11:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof11
		}
	st_case_11:
		switch  lex.data[( lex.p)] {
		case 95:
			goto tr10
		case 117:
			goto st12
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr10
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr10
			}
		default:
			goto tr10
		}
		goto tr18
	st12:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof12
		}
	st_case_12:
		switch  lex.data[( lex.p)] {
		case 95:
			goto tr10
		case 116:
			goto st13
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr10
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr10
			}
		default:
			goto tr10
		}
		goto tr18
	st13:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof13
		}
	st_case_13:
		switch  lex.data[( lex.p)] {
		case 95:
			goto tr10
		case 97:
			goto st14
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr10
			}
		case  lex.data[( lex.p)] > 90:
			if 98 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr10
			}
		default:
			goto tr10
		}
		goto tr18
	st14:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof14
		}
	st_case_14:
		switch  lex.data[( lex.p)] {
		case 95:
			goto tr10
		case 116:
			goto st15
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr10
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr10
			}
		default:
			goto tr10
		}
		goto tr18
	st15:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof15
		}
	st_case_15:
		switch  lex.data[( lex.p)] {
		case 95:
			goto tr10
		case 105:
			goto st16
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr10
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr10
			}
		default:
			goto tr10
		}
		goto tr18
	st16:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof16
		}
	st_case_16:
		switch  lex.data[( lex.p)] {
		case 95:
			goto tr10
		case 111:
			goto st17
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr10
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr10
			}
		default:
			goto tr10
		}
		goto tr18
	st17:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof17
		}
	st_case_17:
		switch  lex.data[( lex.p)] {
		case 95:
			goto tr10
		case 110:
			goto tr29
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr10
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr10
			}
		default:
			goto tr10
		}
		goto tr18
	st18:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof18
		}
	st_case_18:
		switch  lex.data[( lex.p)] {
		case 95:
			goto tr10
		case 117:
			goto st19
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr10
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr10
			}
		default:
			goto tr10
		}
		goto tr18
	st19:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof19
		}
	st_case_19:
		switch  lex.data[( lex.p)] {
		case 95:
			goto tr10
		case 101:
			goto st20
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr10
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr10
			}
		default:
			goto tr10
		}
		goto tr18
	st20:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof20
		}
	st_case_20:
		switch  lex.data[( lex.p)] {
		case 95:
			goto tr10
		case 114:
			goto st21
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr10
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr10
			}
		default:
			goto tr10
		}
		goto tr18
	st21:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof21
		}
	st_case_21:
		switch  lex.data[( lex.p)] {
		case 95:
			goto tr10
		case 121:
			goto tr33
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr10
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr10
			}
		default:
			goto tr10
		}
		goto tr18
	st22:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof22
		}
	st_case_22:
		switch  lex.data[( lex.p)] {
		case 95:
			goto tr10
		case 114:
			goto st23
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr10
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr10
			}
		default:
			goto tr10
		}
		goto tr18
	st23:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof23
		}
	st_case_23:
		switch  lex.data[( lex.p)] {
		case 95:
			goto tr10
		case 117:
			goto st10
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr10
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr10
			}
		default:
			goto tr10
		}
		goto tr18
	st_out:
	_test_eof2:  lex.cs = 2; goto _test_eof
	_test_eof3:  lex.cs = 3; goto _test_eof
	_test_eof0:  lex.cs = 0; goto _test_eof
	_test_eof1:  lex.cs = 1; goto _test_eof
	_test_eof4:  lex.cs = 4; goto _test_eof
	_test_eof5:  lex.cs = 5; goto _test_eof
	_test_eof6:  lex.cs = 6; goto _test_eof
	_test_eof7:  lex.cs = 7; goto _test_eof
	_test_eof8:  lex.cs = 8; goto _test_eof
	_test_eof9:  lex.cs = 9; goto _test_eof
	_test_eof10:  lex.cs = 10; goto _test_eof
	_test_eof11:  lex.cs = 11; goto _test_eof
	_test_eof12:  lex.cs = 12; goto _test_eof
	_test_eof13:  lex.cs = 13; goto _test_eof
	_test_eof14:  lex.cs = 14; goto _test_eof
	_test_eof15:  lex.cs = 15; goto _test_eof
	_test_eof16:  lex.cs = 16; goto _test_eof
	_test_eof17:  lex.cs = 17; goto _test_eof
	_test_eof18:  lex.cs = 18; goto _test_eof
	_test_eof19:  lex.cs = 19; goto _test_eof
	_test_eof20:  lex.cs = 20; goto _test_eof
	_test_eof21:  lex.cs = 21; goto _test_eof
	_test_eof22:  lex.cs = 22; goto _test_eof
	_test_eof23:  lex.cs = 23; goto _test_eof

	_test_eof: {}
	if ( lex.p) == eof {
		switch  lex.cs {
		case 3:
			goto tr15
		case 0:
			goto tr0
		case 1:
			goto tr0
		case 4:
			goto tr15
		case 5:
			goto tr16
		case 6:
			goto tr17
		case 7:
			goto tr18
		case 8:
			goto tr18
		case 9:
			goto tr18
		case 10:
			goto tr18
		case 11:
			goto tr18
		case 12:
			goto tr18
		case 13:
			goto tr18
		case 14:
			goto tr18
		case 15:
			goto tr18
		case 16:
			goto tr18
		case 17:
			goto tr18
		case 18:
			goto tr18
		case 19:
			goto tr18
		case 20:
			goto tr18
		case 21:
			goto tr18
		case 22:
			goto tr18
		case 23:
			goto tr18
		}
	}

	_out: {}
	}

//line graphql_lexer.rl:85


    return tok;
}

func (lex *lexer) Error(e string) {
    fmt.Println("errorx:", e)
    fmt.Println("p", string(lex.data[lex.p:lex.pe]))
    fmt.Println("t", string(lex.data[lex.ts:lex.te]))
}
