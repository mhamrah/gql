
//line graphql_lexer.rl:1
package query

import (
        "fmt"
        "reflect"
        "strconv"
)


//line graphql_lexer_gen.go:13
const graphql_start int = 3
const graphql_first_final int = 3
const graphql_error int = -1

const graphql_en_main int = 3


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
	case 3:
		goto st_case_3
	case 4:
		goto st_case_4
	case 0:
		goto st_case_0
	case 1:
		goto st_case_1
	case 5:
		goto st_case_5
	case 6:
		goto st_case_6
	case 7:
		goto st_case_7
	case 8:
		goto st_case_8
	case 2:
		goto st_case_2
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
	case 24:
		goto st_case_24
	case 25:
		goto st_case_25
	case 26:
		goto st_case_26
	case 27:
		goto st_case_27
	}
	goto st_out
tr0:
//line graphql_lexer.rl:99
( lex.p) = ( lex.te) - 1
{ }
	goto st3
tr2:
//line graphql_lexer.rl:83
 lex.te = ( lex.p)+1
{
                out.val = reflect.ValueOf(string(lex.data[lex.ts+1:lex.te-1]));
                tok = VALUE
                {( lex.p)++;  lex.cs = 3; goto _out }
            }
	goto st3
tr4:
//line graphql_lexer.rl:61
 lex.te = ( lex.p)+1
{
                tok = SPREAD;
                {( lex.p)++;  lex.cs = 3; goto _out }
            }
	goto st3
tr5:
//line graphql_lexer.rl:99
 lex.te = ( lex.p)+1
{ }
	goto st3
tr6:
//line graphql_lexer.rl:98
 lex.te = ( lex.p)+1

	goto st3
tr9:
//line graphql_lexer.rl:88
 lex.te = ( lex.p)+1
{

                tok = int(lex.data[lex.ts])
                {( lex.p)++;  lex.cs = 3; goto _out }
            }
	goto st3
tr19:
//line graphql_lexer.rl:99
 lex.te = ( lex.p)
( lex.p)--
{ }
	goto st3
tr20:
//line NONE:1
	switch  lex.act {
	case 1:
	{( lex.p) = ( lex.te) - 1
 tok = QUERY; {( lex.p)++;  lex.cs = 3; goto _out }}
	case 2:
	{( lex.p) = ( lex.te) - 1
 tok = MUTATION; {( lex.p)++;  lex.cs = 3; goto _out }}
	case 3:
	{( lex.p) = ( lex.te) - 1

               out.val = reflect.ValueOf(string(lex.data[lex.ts:lex.te]) == "true");
               tok = VALUE
               {( lex.p)++;  lex.cs = 3; goto _out }
            }
	case 5:
	{( lex.p) = ( lex.te) - 1

                tok = ON;
                {( lex.p)++;  lex.cs = 3; goto _out }
            }
	case 6:
	{( lex.p) = ( lex.te) - 1

                out.str = string(lex.data[lex.ts:lex.te]);
                tok = NAME;
                {( lex.p)++;  lex.cs = 3; goto _out }
            }
	case 10:
	{( lex.p) = ( lex.te) - 1

                out.str = string(lex.data[lex.ts+1:lex.te]);
                tok = VARIABLE
                {( lex.p)++;  lex.cs = 3; goto _out }
            }
	case 12:
	{( lex.p) = ( lex.te) - 1
 }
	}
	
	goto st3
tr22:
//line graphql_lexer.rl:74
 lex.te = ( lex.p)
( lex.p)--
{
                i, err := strconv.Atoi(string(lex.data[lex.ts:lex.te]));
                if err != nil {
                    lex.Error("can't convert int")
                }
                out.val = reflect.ValueOf(i);
                tok = VALUE;
                {( lex.p)++;  lex.cs = 3; goto _out }
            }
	goto st3
tr24:
//line graphql_lexer.rl:69
 lex.te = ( lex.p)
( lex.p)--
{
                out.str = string(lex.data[lex.ts:lex.te]);
                tok = NAME;
                {( lex.p)++;  lex.cs = 3; goto _out }
            }
	goto st3
	st3:
//line NONE:1
 lex.ts = 0

		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof3
		}
	st_case_3:
//line NONE:1
 lex.ts = ( lex.p)

//line graphql_lexer_gen.go:245
		switch  lex.data[( lex.p)] {
		case 32:
			goto tr6
		case 34:
			goto tr7
		case 36:
			goto tr8
		case 45:
			goto st6
		case 46:
			goto tr11
		case 58:
			goto tr9
		case 61:
			goto tr9
		case 64:
			goto tr9
		case 91:
			goto tr9
		case 93:
			goto tr9
		case 95:
			goto tr13
		case 102:
			goto st10
		case 109:
			goto st14
		case 111:
			goto st21
		case 113:
			goto st22
		case 116:
			goto st26
		case 123:
			goto tr9
		case 125:
			goto tr9
		}
		switch {
		case  lex.data[( lex.p)] < 49:
			switch {
			case  lex.data[( lex.p)] > 13:
				if 40 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 41 {
					goto tr9
				}
			case  lex.data[( lex.p)] >= 9:
				goto tr6
			}
		case  lex.data[( lex.p)] > 57:
			switch {
			case  lex.data[( lex.p)] > 90:
				if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
					goto tr13
				}
			case  lex.data[( lex.p)] >= 65:
				goto tr13
			}
		default:
			goto st7
		}
		goto tr5
tr7:
//line NONE:1
 lex.te = ( lex.p)+1

	goto st4
	st4:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof4
		}
	st_case_4:
//line graphql_lexer_gen.go:317
		switch  lex.data[( lex.p)] {
		case 10:
			goto tr19
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
tr8:
//line NONE:1
 lex.te = ( lex.p)+1

//line graphql_lexer.rl:99
 lex.act = 12;
	goto st5
tr21:
//line NONE:1
 lex.te = ( lex.p)+1

//line graphql_lexer.rl:93
 lex.act = 10;
	goto st5
	st5:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof5
		}
	st_case_5:
//line graphql_lexer_gen.go:366
		if  lex.data[( lex.p)] == 95 {
			goto tr21
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr21
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr21
			}
		default:
			goto tr21
		}
		goto tr20
	st6:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof6
		}
	st_case_6:
		if 49 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
			goto st7
		}
		goto tr19
	st7:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof7
		}
	st_case_7:
		if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
			goto st7
		}
		goto tr22
tr11:
//line NONE:1
 lex.te = ( lex.p)+1

	goto st8
	st8:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof8
		}
	st_case_8:
//line graphql_lexer_gen.go:411
		if  lex.data[( lex.p)] == 46 {
			goto st2
		}
		goto tr19
	st2:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof2
		}
	st_case_2:
		if  lex.data[( lex.p)] == 46 {
			goto tr4
		}
		goto tr0
tr13:
//line NONE:1
 lex.te = ( lex.p)+1

//line graphql_lexer.rl:69
 lex.act = 6;
	goto st9
tr28:
//line NONE:1
 lex.te = ( lex.p)+1

//line graphql_lexer.rl:56
 lex.act = 3;
	goto st9
tr35:
//line NONE:1
 lex.te = ( lex.p)+1

//line graphql_lexer.rl:55
 lex.act = 2;
	goto st9
tr36:
//line NONE:1
 lex.te = ( lex.p)+1

//line graphql_lexer.rl:65
 lex.act = 5;
	goto st9
tr40:
//line NONE:1
 lex.te = ( lex.p)+1

//line graphql_lexer.rl:54
 lex.act = 1;
	goto st9
	st9:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof9
		}
	st_case_9:
//line graphql_lexer_gen.go:465
		if  lex.data[( lex.p)] == 95 {
			goto tr13
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr13
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr13
			}
		default:
			goto tr13
		}
		goto tr20
	st10:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof10
		}
	st_case_10:
		switch  lex.data[( lex.p)] {
		case 95:
			goto tr13
		case 97:
			goto st11
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr13
			}
		case  lex.data[( lex.p)] > 90:
			if 98 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr13
			}
		default:
			goto tr13
		}
		goto tr24
	st11:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof11
		}
	st_case_11:
		switch  lex.data[( lex.p)] {
		case 95:
			goto tr13
		case 108:
			goto st12
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr13
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr13
			}
		default:
			goto tr13
		}
		goto tr24
	st12:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof12
		}
	st_case_12:
		switch  lex.data[( lex.p)] {
		case 95:
			goto tr13
		case 115:
			goto st13
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr13
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr13
			}
		default:
			goto tr13
		}
		goto tr24
	st13:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof13
		}
	st_case_13:
		switch  lex.data[( lex.p)] {
		case 95:
			goto tr13
		case 101:
			goto tr28
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr13
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr13
			}
		default:
			goto tr13
		}
		goto tr24
	st14:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof14
		}
	st_case_14:
		switch  lex.data[( lex.p)] {
		case 95:
			goto tr13
		case 117:
			goto st15
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr13
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr13
			}
		default:
			goto tr13
		}
		goto tr24
	st15:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof15
		}
	st_case_15:
		switch  lex.data[( lex.p)] {
		case 95:
			goto tr13
		case 116:
			goto st16
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr13
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr13
			}
		default:
			goto tr13
		}
		goto tr24
	st16:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof16
		}
	st_case_16:
		switch  lex.data[( lex.p)] {
		case 95:
			goto tr13
		case 97:
			goto st17
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr13
			}
		case  lex.data[( lex.p)] > 90:
			if 98 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr13
			}
		default:
			goto tr13
		}
		goto tr24
	st17:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof17
		}
	st_case_17:
		switch  lex.data[( lex.p)] {
		case 95:
			goto tr13
		case 116:
			goto st18
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr13
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr13
			}
		default:
			goto tr13
		}
		goto tr24
	st18:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof18
		}
	st_case_18:
		switch  lex.data[( lex.p)] {
		case 95:
			goto tr13
		case 105:
			goto st19
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr13
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr13
			}
		default:
			goto tr13
		}
		goto tr24
	st19:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof19
		}
	st_case_19:
		switch  lex.data[( lex.p)] {
		case 95:
			goto tr13
		case 111:
			goto st20
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr13
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr13
			}
		default:
			goto tr13
		}
		goto tr24
	st20:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof20
		}
	st_case_20:
		switch  lex.data[( lex.p)] {
		case 95:
			goto tr13
		case 110:
			goto tr35
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr13
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr13
			}
		default:
			goto tr13
		}
		goto tr24
	st21:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof21
		}
	st_case_21:
		switch  lex.data[( lex.p)] {
		case 95:
			goto tr13
		case 110:
			goto tr36
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr13
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr13
			}
		default:
			goto tr13
		}
		goto tr24
	st22:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof22
		}
	st_case_22:
		switch  lex.data[( lex.p)] {
		case 95:
			goto tr13
		case 117:
			goto st23
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr13
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr13
			}
		default:
			goto tr13
		}
		goto tr24
	st23:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof23
		}
	st_case_23:
		switch  lex.data[( lex.p)] {
		case 95:
			goto tr13
		case 101:
			goto st24
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr13
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr13
			}
		default:
			goto tr13
		}
		goto tr24
	st24:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof24
		}
	st_case_24:
		switch  lex.data[( lex.p)] {
		case 95:
			goto tr13
		case 114:
			goto st25
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr13
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr13
			}
		default:
			goto tr13
		}
		goto tr24
	st25:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof25
		}
	st_case_25:
		switch  lex.data[( lex.p)] {
		case 95:
			goto tr13
		case 121:
			goto tr40
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr13
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr13
			}
		default:
			goto tr13
		}
		goto tr24
	st26:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof26
		}
	st_case_26:
		switch  lex.data[( lex.p)] {
		case 95:
			goto tr13
		case 114:
			goto st27
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr13
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr13
			}
		default:
			goto tr13
		}
		goto tr24
	st27:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof27
		}
	st_case_27:
		switch  lex.data[( lex.p)] {
		case 95:
			goto tr13
		case 117:
			goto st13
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr13
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr13
			}
		default:
			goto tr13
		}
		goto tr24
	st_out:
	_test_eof3:  lex.cs = 3; goto _test_eof
	_test_eof4:  lex.cs = 4; goto _test_eof
	_test_eof0:  lex.cs = 0; goto _test_eof
	_test_eof1:  lex.cs = 1; goto _test_eof
	_test_eof5:  lex.cs = 5; goto _test_eof
	_test_eof6:  lex.cs = 6; goto _test_eof
	_test_eof7:  lex.cs = 7; goto _test_eof
	_test_eof8:  lex.cs = 8; goto _test_eof
	_test_eof2:  lex.cs = 2; goto _test_eof
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
	_test_eof24:  lex.cs = 24; goto _test_eof
	_test_eof25:  lex.cs = 25; goto _test_eof
	_test_eof26:  lex.cs = 26; goto _test_eof
	_test_eof27:  lex.cs = 27; goto _test_eof

	_test_eof: {}
	if ( lex.p) == eof {
		switch  lex.cs {
		case 4:
			goto tr19
		case 0:
			goto tr0
		case 1:
			goto tr0
		case 5:
			goto tr20
		case 6:
			goto tr19
		case 7:
			goto tr22
		case 8:
			goto tr19
		case 2:
			goto tr0
		case 9:
			goto tr20
		case 10:
			goto tr24
		case 11:
			goto tr24
		case 12:
			goto tr24
		case 13:
			goto tr24
		case 14:
			goto tr24
		case 15:
			goto tr24
		case 16:
			goto tr24
		case 17:
			goto tr24
		case 18:
			goto tr24
		case 19:
			goto tr24
		case 20:
			goto tr24
		case 21:
			goto tr24
		case 22:
			goto tr24
		case 23:
			goto tr24
		case 24:
			goto tr24
		case 25:
			goto tr24
		case 26:
			goto tr24
		case 27:
			goto tr24
		}
	}

	_out: {}
	}

//line graphql_lexer.rl:103


    return tok;
}

func (lex *lexer) Error(e string) {
    fmt.Println("errorx:", e)
    fmt.Println("t", string(lex.data[lex.ts:lex.te]))
    fmt.Println("p", string(lex.data[lex.p:lex.pe]))
}
