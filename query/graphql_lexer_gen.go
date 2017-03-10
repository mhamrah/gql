
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
	case 28:
		goto st_case_28
	case 29:
		goto st_case_29
	case 30:
		goto st_case_30
	case 31:
		goto st_case_31
	case 32:
		goto st_case_32
	case 33:
		goto st_case_33
	case 34:
		goto st_case_34
	case 35:
		goto st_case_35
	case 36:
		goto st_case_36
	case 37:
		goto st_case_37
	case 38:
		goto st_case_38
	}
	goto st_out
tr0:
//line graphql_lexer.rl:97
( lex.p) = ( lex.te) - 1
{ }
	goto st3
tr2:
//line graphql_lexer.rl:81
 lex.te = ( lex.p)+1
{
                out.val = reflect.ValueOf(string(lex.data[lex.ts+1:lex.te-1]));
                tok = VALUE
                {( lex.p)++;  lex.cs = 3; goto _out }
            }
	goto st3
tr4:
//line graphql_lexer.rl:59
 lex.te = ( lex.p)+1
{
                tok = SPREAD;
                {( lex.p)++;  lex.cs = 3; goto _out }
            }
	goto st3
tr5:
//line graphql_lexer.rl:97
 lex.te = ( lex.p)+1
{ }
	goto st3
tr6:
//line graphql_lexer.rl:96
 lex.te = ( lex.p)+1

	goto st3
tr9:
//line graphql_lexer.rl:86
 lex.te = ( lex.p)+1
{

                tok = int(lex.data[lex.ts])
                {( lex.p)++;  lex.cs = 3; goto _out }
            }
	goto st3
tr20:
//line graphql_lexer.rl:97
 lex.te = ( lex.p)
( lex.p)--
{ }
	goto st3
tr21:
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
 tok = SUBSCRIPTION; {( lex.p)++;  lex.cs = 3; goto _out } }
	case 4:
	{( lex.p) = ( lex.te) - 1

               out.val = reflect.ValueOf(string(lex.data[lex.ts:lex.te]) == "true");
               tok = VALUE
               {( lex.p)++;  lex.cs = 3; goto _out }
            }
	case 6:
	{( lex.p) = ( lex.te) - 1

                tok = ON;
                {( lex.p)++;  lex.cs = 3; goto _out }
            }
	case 7:
	{( lex.p) = ( lex.te) - 1

                out.str = string(lex.data[lex.ts:lex.te]);
                tok = NAME;
                {( lex.p)++;  lex.cs = 3; goto _out }
            }
	case 11:
	{( lex.p) = ( lex.te) - 1

                out.str = string(lex.data[lex.ts+1:lex.te]);
                tok = VARIABLE
                {( lex.p)++;  lex.cs = 3; goto _out }
            }
	case 13:
	{( lex.p) = ( lex.te) - 1
 }
	}
	
	goto st3
tr23:
//line graphql_lexer.rl:72
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
tr25:
//line graphql_lexer.rl:67
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

//line graphql_lexer_gen.go:270
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
		case 115:
			goto st26
		case 116:
			goto st37
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
//line graphql_lexer_gen.go:344
		switch  lex.data[( lex.p)] {
		case 10:
			goto tr20
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

//line graphql_lexer.rl:97
 lex.act = 13;
	goto st5
tr22:
//line NONE:1
 lex.te = ( lex.p)+1

//line graphql_lexer.rl:91
 lex.act = 11;
	goto st5
	st5:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof5
		}
	st_case_5:
//line graphql_lexer_gen.go:393
		if  lex.data[( lex.p)] == 95 {
			goto tr22
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr22
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr22
			}
		default:
			goto tr22
		}
		goto tr21
	st6:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof6
		}
	st_case_6:
		if 49 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
			goto st7
		}
		goto tr20
	st7:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof7
		}
	st_case_7:
		if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
			goto st7
		}
		goto tr23
tr11:
//line NONE:1
 lex.te = ( lex.p)+1

	goto st8
	st8:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof8
		}
	st_case_8:
//line graphql_lexer_gen.go:438
		if  lex.data[( lex.p)] == 46 {
			goto st2
		}
		goto tr20
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

//line graphql_lexer.rl:67
 lex.act = 7;
	goto st9
tr29:
//line NONE:1
 lex.te = ( lex.p)+1

//line graphql_lexer.rl:54
 lex.act = 4;
	goto st9
tr36:
//line NONE:1
 lex.te = ( lex.p)+1

//line graphql_lexer.rl:52
 lex.act = 2;
	goto st9
tr37:
//line NONE:1
 lex.te = ( lex.p)+1

//line graphql_lexer.rl:63
 lex.act = 6;
	goto st9
tr41:
//line NONE:1
 lex.te = ( lex.p)+1

//line graphql_lexer.rl:51
 lex.act = 1;
	goto st9
tr52:
//line NONE:1
 lex.te = ( lex.p)+1

//line graphql_lexer.rl:53
 lex.act = 3;
	goto st9
	st9:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof9
		}
	st_case_9:
//line graphql_lexer_gen.go:499
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
		goto tr21
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
		goto tr25
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
		goto tr25
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
		goto tr25
	st13:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof13
		}
	st_case_13:
		switch  lex.data[( lex.p)] {
		case 95:
			goto tr13
		case 101:
			goto tr29
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
		goto tr25
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
		goto tr25
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
		goto tr25
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
		goto tr25
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
		goto tr25
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
		goto tr25
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
		goto tr25
	st20:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof20
		}
	st_case_20:
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
		goto tr25
	st21:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof21
		}
	st_case_21:
		switch  lex.data[( lex.p)] {
		case 95:
			goto tr13
		case 110:
			goto tr37
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
		goto tr25
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
		goto tr25
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
		goto tr25
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
		goto tr25
	st25:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof25
		}
	st_case_25:
		switch  lex.data[( lex.p)] {
		case 95:
			goto tr13
		case 121:
			goto tr41
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
		goto tr25
	st26:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof26
		}
	st_case_26:
		switch  lex.data[( lex.p)] {
		case 95:
			goto tr13
		case 117:
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
		goto tr25
	st27:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof27
		}
	st_case_27:
		switch  lex.data[( lex.p)] {
		case 95:
			goto tr13
		case 98:
			goto st28
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
		goto tr25
	st28:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof28
		}
	st_case_28:
		switch  lex.data[( lex.p)] {
		case 95:
			goto tr13
		case 115:
			goto st29
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
		goto tr25
	st29:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof29
		}
	st_case_29:
		switch  lex.data[( lex.p)] {
		case 95:
			goto tr13
		case 99:
			goto st30
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
		goto tr25
	st30:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof30
		}
	st_case_30:
		switch  lex.data[( lex.p)] {
		case 95:
			goto tr13
		case 114:
			goto st31
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
		goto tr25
	st31:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof31
		}
	st_case_31:
		switch  lex.data[( lex.p)] {
		case 95:
			goto tr13
		case 105:
			goto st32
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
		goto tr25
	st32:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof32
		}
	st_case_32:
		switch  lex.data[( lex.p)] {
		case 95:
			goto tr13
		case 112:
			goto st33
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
		goto tr25
	st33:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof33
		}
	st_case_33:
		switch  lex.data[( lex.p)] {
		case 95:
			goto tr13
		case 116:
			goto st34
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
		goto tr25
	st34:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof34
		}
	st_case_34:
		switch  lex.data[( lex.p)] {
		case 95:
			goto tr13
		case 105:
			goto st35
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
		goto tr25
	st35:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof35
		}
	st_case_35:
		switch  lex.data[( lex.p)] {
		case 95:
			goto tr13
		case 111:
			goto st36
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
		goto tr25
	st36:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof36
		}
	st_case_36:
		switch  lex.data[( lex.p)] {
		case 95:
			goto tr13
		case 110:
			goto tr52
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
		goto tr25
	st37:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof37
		}
	st_case_37:
		switch  lex.data[( lex.p)] {
		case 95:
			goto tr13
		case 114:
			goto st38
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
		goto tr25
	st38:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof38
		}
	st_case_38:
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
		goto tr25
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
	_test_eof28:  lex.cs = 28; goto _test_eof
	_test_eof29:  lex.cs = 29; goto _test_eof
	_test_eof30:  lex.cs = 30; goto _test_eof
	_test_eof31:  lex.cs = 31; goto _test_eof
	_test_eof32:  lex.cs = 32; goto _test_eof
	_test_eof33:  lex.cs = 33; goto _test_eof
	_test_eof34:  lex.cs = 34; goto _test_eof
	_test_eof35:  lex.cs = 35; goto _test_eof
	_test_eof36:  lex.cs = 36; goto _test_eof
	_test_eof37:  lex.cs = 37; goto _test_eof
	_test_eof38:  lex.cs = 38; goto _test_eof

	_test_eof: {}
	if ( lex.p) == eof {
		switch  lex.cs {
		case 4:
			goto tr20
		case 0:
			goto tr0
		case 1:
			goto tr0
		case 5:
			goto tr21
		case 6:
			goto tr20
		case 7:
			goto tr23
		case 8:
			goto tr20
		case 2:
			goto tr0
		case 9:
			goto tr21
		case 10:
			goto tr25
		case 11:
			goto tr25
		case 12:
			goto tr25
		case 13:
			goto tr25
		case 14:
			goto tr25
		case 15:
			goto tr25
		case 16:
			goto tr25
		case 17:
			goto tr25
		case 18:
			goto tr25
		case 19:
			goto tr25
		case 20:
			goto tr25
		case 21:
			goto tr25
		case 22:
			goto tr25
		case 23:
			goto tr25
		case 24:
			goto tr25
		case 25:
			goto tr25
		case 26:
			goto tr25
		case 27:
			goto tr25
		case 28:
			goto tr25
		case 29:
			goto tr25
		case 30:
			goto tr25
		case 31:
			goto tr25
		case 32:
			goto tr25
		case 33:
			goto tr25
		case 34:
			goto tr25
		case 35:
			goto tr25
		case 36:
			goto tr25
		case 37:
			goto tr25
		case 38:
			goto tr25
		}
	}

	_out: {}
	}

//line graphql_lexer.rl:101


    return tok;
}

func (lex *lexer) Error(e string) {
    fmt.Println("errorx:", e)
    fmt.Println("t", string(lex.data[lex.ts:lex.te]))
    fmt.Println("p", string(lex.data[lex.p:lex.pe]))
}
