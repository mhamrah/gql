
//line lexer.rl:1
package parser

import (
        "fmt"
        "reflect"
        "strconv"
        "github.com/mhamrah/gql"
)


//line lexer_gen.go:14
const graphql_start int = 3
const graphql_first_final int = 3
const graphql_error int = -1

const graphql_en_main int = 3


//line lexer.rl:16


type lexer struct {
    doc gql.Document
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
    
//line lexer_gen.go:42
	{
	 lex.cs = graphql_start
	 lex.ts = 0
	 lex.te = 0
	 lex.act = 0
	}

//line lexer.rl:35
    return lex
}

func (lex *lexer) Lex(out *yySymType) int {
    eof := lex.pe
    tok := 0
    
//line lexer_gen.go:58
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
	case 39:
		goto st_case_39
	case 40:
		goto st_case_40
	case 41:
		goto st_case_41
	case 42:
		goto st_case_42
	case 43:
		goto st_case_43
	case 44:
		goto st_case_44
	case 45:
		goto st_case_45
	case 46:
		goto st_case_46
	case 47:
		goto st_case_47
	case 48:
		goto st_case_48
	case 49:
		goto st_case_49
	case 50:
		goto st_case_50
	case 51:
		goto st_case_51
	case 52:
		goto st_case_52
	case 53:
		goto st_case_53
	case 54:
		goto st_case_54
	case 55:
		goto st_case_55
	case 56:
		goto st_case_56
	case 57:
		goto st_case_57
	case 58:
		goto st_case_58
	case 59:
		goto st_case_59
	case 60:
		goto st_case_60
	case 61:
		goto st_case_61
	case 62:
		goto st_case_62
	case 63:
		goto st_case_63
	case 64:
		goto st_case_64
	case 65:
		goto st_case_65
	case 66:
		goto st_case_66
	case 67:
		goto st_case_67
	case 68:
		goto st_case_68
	case 69:
		goto st_case_69
	case 70:
		goto st_case_70
	case 71:
		goto st_case_71
	case 72:
		goto st_case_72
	case 73:
		goto st_case_73
	case 74:
		goto st_case_74
	case 75:
		goto st_case_75
	case 76:
		goto st_case_76
	case 77:
		goto st_case_77
	case 78:
		goto st_case_78
	case 79:
		goto st_case_79
	case 80:
		goto st_case_80
	case 81:
		goto st_case_81
	case 82:
		goto st_case_82
	case 83:
		goto st_case_83
	case 84:
		goto st_case_84
	case 85:
		goto st_case_85
	case 86:
		goto st_case_86
	case 87:
		goto st_case_87
	case 88:
		goto st_case_88
	case 89:
		goto st_case_89
	case 90:
		goto st_case_90
	case 91:
		goto st_case_91
	case 92:
		goto st_case_92
	case 93:
		goto st_case_93
	case 94:
		goto st_case_94
	}
	goto st_out
tr0:
//line lexer.rl:107
( lex.p) = ( lex.te) - 1
{ }
	goto st3
tr2:
//line lexer.rl:92
 lex.te = ( lex.p)+1
{
                out.val = reflect.ValueOf(string(lex.data[lex.ts+1:lex.te-1]));
                tok = VALUE
                {( lex.p)++;  lex.cs = 3; goto _out }
            }
	goto st3
tr4:
//line lexer.rl:70
 lex.te = ( lex.p)+1
{
                tok = SPREAD;
                {( lex.p)++;  lex.cs = 3; goto _out }
            }
	goto st3
tr5:
//line lexer.rl:107
 lex.te = ( lex.p)+1
{ }
	goto st3
tr6:
//line lexer.rl:106
 lex.te = ( lex.p)+1

	goto st3
tr7:
//line lexer.rl:97
 lex.te = ( lex.p)+1
{
                tok = int(lex.data[lex.ts])
                {( lex.p)++;  lex.cs = 3; goto _out }
            }
	goto st3
tr25:
//line lexer.rl:107
 lex.te = ( lex.p)
( lex.p)--
{ }
	goto st3
tr26:
//line NONE:1
	switch  lex.act {
	case 1:
	{( lex.p) = ( lex.te) - 1
 tok = DIRECTIVE; {( lex.p)++;  lex.cs = 3; goto _out } }
	case 2:
	{( lex.p) = ( lex.te) - 1
 tok = ENUM; {( lex.p)++;  lex.cs = 3; goto _out } }
	case 3:
	{( lex.p) = ( lex.te) - 1
 tok = EXTEND; {( lex.p)++;  lex.cs = 3; goto _out } }
	case 4:
	{( lex.p) = ( lex.te) - 1
 tok = FALSE; {( lex.p)++;  lex.cs = 3; goto _out } }
	case 5:
	{( lex.p) = ( lex.te) - 1
 tok = FRAGMENT; {( lex.p)++;  lex.cs = 3; goto _out } }
	case 6:
	{( lex.p) = ( lex.te) - 1
 tok = IMPLEMENTS; {( lex.p)++;  lex.cs = 3; goto _out } }
	case 7:
	{( lex.p) = ( lex.te) - 1
 tok = INPUT; {( lex.p)++;  lex.cs = 3; goto _out } }
	case 8:
	{( lex.p) = ( lex.te) - 1
 tok = INTERFACE; {( lex.p)++;  lex.cs = 3; goto _out } }
	case 9:
	{( lex.p) = ( lex.te) - 1
 tok = MUTATION; {( lex.p)++;  lex.cs = 3; goto _out }}
	case 10:
	{( lex.p) = ( lex.te) - 1
 tok = NULL; {( lex.p)++;  lex.cs = 3; goto _out } }
	case 11:
	{( lex.p) = ( lex.te) - 1
 tok = QUERY; {( lex.p)++;  lex.cs = 3; goto _out }}
	case 12:
	{( lex.p) = ( lex.te) - 1
 tok = ON; {( lex.p)++;  lex.cs = 3; goto _out } }
	case 13:
	{( lex.p) = ( lex.te) - 1
 tok = SCALAR; {( lex.p)++;  lex.cs = 3; goto _out } }
	case 14:
	{( lex.p) = ( lex.te) - 1
 tok = SCHEMA; {( lex.p)++;  lex.cs = 3; goto _out } }
	case 15:
	{( lex.p) = ( lex.te) - 1
 tok = SUBSCRIPTION; {( lex.p)++;  lex.cs = 3; goto _out } }
	case 16:
	{( lex.p) = ( lex.te) - 1
 tok = TRUE; {( lex.p)++;  lex.cs = 3; goto _out } }
	case 17:
	{( lex.p) = ( lex.te) - 1
 tok = TYPE; {( lex.p)++;  lex.cs = 3; goto _out } }
	case 18:
	{( lex.p) = ( lex.te) - 1
 tok = UNION; {( lex.p)++;  lex.cs = 3; goto _out } }
	case 21:
	{( lex.p) = ( lex.te) - 1

                out.str = string(lex.data[lex.ts:lex.te]);
                tok = IDENTIFIER;
                {( lex.p)++;  lex.cs = 3; goto _out }
            }
	case 25:
	{( lex.p) = ( lex.te) - 1

                out.str = string(lex.data[lex.ts+1:lex.te]);
                tok = VARIABLE
                {( lex.p)++;  lex.cs = 3; goto _out }
            }
	case 27:
	{( lex.p) = ( lex.te) - 1
 }
	}
	
	goto st3
tr28:
//line lexer.rl:83
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
tr30:
//line lexer.rl:78
 lex.te = ( lex.p)
( lex.p)--
{
                out.str = string(lex.data[lex.ts:lex.te]);
                tok = IDENTIFIER;
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

//line lexer_gen.go:414
		switch  lex.data[( lex.p)] {
		case 32:
			goto tr6
		case 33:
			goto tr7
		case 34:
			goto tr8
		case 36:
			goto tr9
		case 45:
			goto st6
		case 46:
			goto tr11
		case 58:
			goto tr7
		case 61:
			goto tr7
		case 64:
			goto tr7
		case 91:
			goto tr7
		case 93:
			goto tr7
		case 95:
			goto tr13
		case 100:
			goto st10
		case 101:
			goto st18
		case 102:
			goto st25
		case 105:
			goto st35
		case 109:
			goto st53
		case 110:
			goto st60
		case 111:
			goto st63
		case 113:
			goto st64
		case 115:
			goto st68
		case 116:
			goto st86
		case 117:
			goto st91
		}
		switch {
		case  lex.data[( lex.p)] < 49:
			switch {
			case  lex.data[( lex.p)] > 13:
				if 40 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 41 {
					goto tr7
				}
			case  lex.data[( lex.p)] >= 9:
				goto tr6
			}
		case  lex.data[( lex.p)] > 57:
			switch {
			case  lex.data[( lex.p)] < 97:
				if 65 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 90 {
					goto tr13
				}
			case  lex.data[( lex.p)] > 122:
				if 123 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 125 {
					goto tr7
				}
			default:
				goto tr13
			}
		default:
			goto st7
		}
		goto tr5
tr8:
//line NONE:1
 lex.te = ( lex.p)+1

	goto st4
	st4:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof4
		}
	st_case_4:
//line lexer_gen.go:500
		switch  lex.data[( lex.p)] {
		case 10:
			goto tr25
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
tr9:
//line NONE:1
 lex.te = ( lex.p)+1

//line lexer.rl:107
 lex.act = 27;
	goto st5
tr27:
//line NONE:1
 lex.te = ( lex.p)+1

//line lexer.rl:101
 lex.act = 25;
	goto st5
	st5:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof5
		}
	st_case_5:
//line lexer_gen.go:549
		if  lex.data[( lex.p)] == 95 {
			goto tr27
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr27
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr27
			}
		default:
			goto tr27
		}
		goto tr26
	st6:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof6
		}
	st_case_6:
		if 49 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
			goto st7
		}
		goto tr25
	st7:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof7
		}
	st_case_7:
		if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
			goto st7
		}
		goto tr28
tr11:
//line NONE:1
 lex.te = ( lex.p)+1

	goto st8
	st8:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof8
		}
	st_case_8:
//line lexer_gen.go:594
		if  lex.data[( lex.p)] == 46 {
			goto st2
		}
		goto tr25
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

//line lexer.rl:78
 lex.act = 21;
	goto st9
tr38:
//line NONE:1
 lex.te = ( lex.p)+1

//line lexer.rl:52
 lex.act = 1;
	goto st9
tr42:
//line NONE:1
 lex.te = ( lex.p)+1

//line lexer.rl:53
 lex.act = 2;
	goto st9
tr46:
//line NONE:1
 lex.te = ( lex.p)+1

//line lexer.rl:54
 lex.act = 3;
	goto st9
tr51:
//line NONE:1
 lex.te = ( lex.p)+1

//line lexer.rl:55
 lex.act = 4;
	goto st9
tr57:
//line NONE:1
 lex.te = ( lex.p)+1

//line lexer.rl:56
 lex.act = 5;
	goto st9
tr67:
//line NONE:1
 lex.te = ( lex.p)+1

//line lexer.rl:57
 lex.act = 6;
	goto st9
tr71:
//line NONE:1
 lex.te = ( lex.p)+1

//line lexer.rl:58
 lex.act = 7;
	goto st9
tr77:
//line NONE:1
 lex.te = ( lex.p)+1

//line lexer.rl:59
 lex.act = 8;
	goto st9
tr84:
//line NONE:1
 lex.te = ( lex.p)+1

//line lexer.rl:60
 lex.act = 9;
	goto st9
tr87:
//line NONE:1
 lex.te = ( lex.p)+1

//line lexer.rl:61
 lex.act = 10;
	goto st9
tr88:
//line NONE:1
 lex.te = ( lex.p)+1

//line lexer.rl:63
 lex.act = 12;
	goto st9
tr92:
//line NONE:1
 lex.te = ( lex.p)+1

//line lexer.rl:62
 lex.act = 11;
	goto st9
tr99:
//line NONE:1
 lex.te = ( lex.p)+1

//line lexer.rl:64
 lex.act = 13;
	goto st9
tr102:
//line NONE:1
 lex.te = ( lex.p)+1

//line lexer.rl:65
 lex.act = 14;
	goto st9
tr112:
//line NONE:1
 lex.te = ( lex.p)+1

//line lexer.rl:66
 lex.act = 15;
	goto st9
tr116:
//line NONE:1
 lex.te = ( lex.p)+1

//line lexer.rl:67
 lex.act = 16;
	goto st9
tr118:
//line NONE:1
 lex.te = ( lex.p)+1

//line lexer.rl:68
 lex.act = 17;
	goto st9
tr122:
//line NONE:1
 lex.te = ( lex.p)+1

//line lexer.rl:69
 lex.act = 18;
	goto st9
	st9:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof9
		}
	st_case_9:
//line lexer_gen.go:746
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
		goto tr26
	st10:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof10
		}
	st_case_10:
		switch  lex.data[( lex.p)] {
		case 95:
			goto tr13
		case 105:
			goto st11
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
		goto tr30
	st11:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof11
		}
	st_case_11:
		switch  lex.data[( lex.p)] {
		case 95:
			goto tr13
		case 114:
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
		goto tr30
	st12:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof12
		}
	st_case_12:
		switch  lex.data[( lex.p)] {
		case 95:
			goto tr13
		case 101:
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
		goto tr30
	st13:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof13
		}
	st_case_13:
		switch  lex.data[( lex.p)] {
		case 95:
			goto tr13
		case 99:
			goto st14
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
		goto tr30
	st14:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof14
		}
	st_case_14:
		switch  lex.data[( lex.p)] {
		case 95:
			goto tr13
		case 116:
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
		goto tr30
	st15:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof15
		}
	st_case_15:
		switch  lex.data[( lex.p)] {
		case 95:
			goto tr13
		case 105:
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
		goto tr30
	st16:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof16
		}
	st_case_16:
		switch  lex.data[( lex.p)] {
		case 95:
			goto tr13
		case 118:
			goto st17
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
		goto tr30
	st17:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof17
		}
	st_case_17:
		switch  lex.data[( lex.p)] {
		case 95:
			goto tr13
		case 101:
			goto tr38
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
		goto tr30
	st18:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof18
		}
	st_case_18:
		switch  lex.data[( lex.p)] {
		case 95:
			goto tr13
		case 110:
			goto st19
		case 120:
			goto st21
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
		goto tr30
	st19:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof19
		}
	st_case_19:
		switch  lex.data[( lex.p)] {
		case 95:
			goto tr13
		case 117:
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
		goto tr30
	st20:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof20
		}
	st_case_20:
		switch  lex.data[( lex.p)] {
		case 95:
			goto tr13
		case 109:
			goto tr42
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
		goto tr30
	st21:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof21
		}
	st_case_21:
		switch  lex.data[( lex.p)] {
		case 95:
			goto tr13
		case 116:
			goto st22
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
		goto tr30
	st22:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof22
		}
	st_case_22:
		switch  lex.data[( lex.p)] {
		case 95:
			goto tr13
		case 101:
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
		goto tr30
	st23:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof23
		}
	st_case_23:
		switch  lex.data[( lex.p)] {
		case 95:
			goto tr13
		case 110:
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
		goto tr30
	st24:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof24
		}
	st_case_24:
		switch  lex.data[( lex.p)] {
		case 95:
			goto tr13
		case 100:
			goto tr46
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
		goto tr30
	st25:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof25
		}
	st_case_25:
		switch  lex.data[( lex.p)] {
		case 95:
			goto tr13
		case 97:
			goto st26
		case 114:
			goto st29
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
		goto tr30
	st26:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof26
		}
	st_case_26:
		switch  lex.data[( lex.p)] {
		case 95:
			goto tr13
		case 108:
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
		goto tr30
	st27:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof27
		}
	st_case_27:
		switch  lex.data[( lex.p)] {
		case 95:
			goto tr13
		case 115:
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
		goto tr30
	st28:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof28
		}
	st_case_28:
		switch  lex.data[( lex.p)] {
		case 95:
			goto tr13
		case 101:
			goto tr51
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
		goto tr30
	st29:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof29
		}
	st_case_29:
		switch  lex.data[( lex.p)] {
		case 95:
			goto tr13
		case 97:
			goto st30
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
		goto tr30
	st30:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof30
		}
	st_case_30:
		switch  lex.data[( lex.p)] {
		case 95:
			goto tr13
		case 103:
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
		goto tr30
	st31:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof31
		}
	st_case_31:
		switch  lex.data[( lex.p)] {
		case 95:
			goto tr13
		case 109:
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
		goto tr30
	st32:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof32
		}
	st_case_32:
		switch  lex.data[( lex.p)] {
		case 95:
			goto tr13
		case 101:
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
		goto tr30
	st33:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof33
		}
	st_case_33:
		switch  lex.data[( lex.p)] {
		case 95:
			goto tr13
		case 110:
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
		goto tr30
	st34:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof34
		}
	st_case_34:
		switch  lex.data[( lex.p)] {
		case 95:
			goto tr13
		case 116:
			goto tr57
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
		goto tr30
	st35:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof35
		}
	st_case_35:
		switch  lex.data[( lex.p)] {
		case 95:
			goto tr13
		case 109:
			goto st36
		case 110:
			goto st44
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
		goto tr30
	st36:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof36
		}
	st_case_36:
		switch  lex.data[( lex.p)] {
		case 95:
			goto tr13
		case 112:
			goto st37
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
		goto tr30
	st37:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof37
		}
	st_case_37:
		switch  lex.data[( lex.p)] {
		case 95:
			goto tr13
		case 108:
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
		goto tr30
	st38:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof38
		}
	st_case_38:
		switch  lex.data[( lex.p)] {
		case 95:
			goto tr13
		case 101:
			goto st39
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
		goto tr30
	st39:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof39
		}
	st_case_39:
		switch  lex.data[( lex.p)] {
		case 95:
			goto tr13
		case 109:
			goto st40
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
		goto tr30
	st40:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof40
		}
	st_case_40:
		switch  lex.data[( lex.p)] {
		case 95:
			goto tr13
		case 101:
			goto st41
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
		goto tr30
	st41:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof41
		}
	st_case_41:
		switch  lex.data[( lex.p)] {
		case 95:
			goto tr13
		case 110:
			goto st42
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
		goto tr30
	st42:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof42
		}
	st_case_42:
		switch  lex.data[( lex.p)] {
		case 95:
			goto tr13
		case 116:
			goto st43
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
		goto tr30
	st43:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof43
		}
	st_case_43:
		switch  lex.data[( lex.p)] {
		case 95:
			goto tr13
		case 115:
			goto tr67
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
		goto tr30
	st44:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof44
		}
	st_case_44:
		switch  lex.data[( lex.p)] {
		case 95:
			goto tr13
		case 112:
			goto st45
		case 116:
			goto st47
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
		goto tr30
	st45:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof45
		}
	st_case_45:
		switch  lex.data[( lex.p)] {
		case 95:
			goto tr13
		case 117:
			goto st46
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
		goto tr30
	st46:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof46
		}
	st_case_46:
		switch  lex.data[( lex.p)] {
		case 95:
			goto tr13
		case 116:
			goto tr71
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
		goto tr30
	st47:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof47
		}
	st_case_47:
		switch  lex.data[( lex.p)] {
		case 95:
			goto tr13
		case 101:
			goto st48
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
		goto tr30
	st48:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof48
		}
	st_case_48:
		switch  lex.data[( lex.p)] {
		case 95:
			goto tr13
		case 114:
			goto st49
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
		goto tr30
	st49:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof49
		}
	st_case_49:
		switch  lex.data[( lex.p)] {
		case 95:
			goto tr13
		case 102:
			goto st50
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
		goto tr30
	st50:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof50
		}
	st_case_50:
		switch  lex.data[( lex.p)] {
		case 95:
			goto tr13
		case 97:
			goto st51
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
		goto tr30
	st51:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof51
		}
	st_case_51:
		switch  lex.data[( lex.p)] {
		case 95:
			goto tr13
		case 99:
			goto st52
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
		goto tr30
	st52:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof52
		}
	st_case_52:
		switch  lex.data[( lex.p)] {
		case 95:
			goto tr13
		case 101:
			goto tr77
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
		goto tr30
	st53:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof53
		}
	st_case_53:
		switch  lex.data[( lex.p)] {
		case 95:
			goto tr13
		case 117:
			goto st54
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
		goto tr30
	st54:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof54
		}
	st_case_54:
		switch  lex.data[( lex.p)] {
		case 95:
			goto tr13
		case 116:
			goto st55
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
		goto tr30
	st55:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof55
		}
	st_case_55:
		switch  lex.data[( lex.p)] {
		case 95:
			goto tr13
		case 97:
			goto st56
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
		goto tr30
	st56:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof56
		}
	st_case_56:
		switch  lex.data[( lex.p)] {
		case 95:
			goto tr13
		case 116:
			goto st57
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
		goto tr30
	st57:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof57
		}
	st_case_57:
		switch  lex.data[( lex.p)] {
		case 95:
			goto tr13
		case 105:
			goto st58
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
		goto tr30
	st58:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof58
		}
	st_case_58:
		switch  lex.data[( lex.p)] {
		case 95:
			goto tr13
		case 111:
			goto st59
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
		goto tr30
	st59:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof59
		}
	st_case_59:
		switch  lex.data[( lex.p)] {
		case 95:
			goto tr13
		case 110:
			goto tr84
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
		goto tr30
	st60:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof60
		}
	st_case_60:
		switch  lex.data[( lex.p)] {
		case 95:
			goto tr13
		case 117:
			goto st61
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
		goto tr30
	st61:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof61
		}
	st_case_61:
		switch  lex.data[( lex.p)] {
		case 95:
			goto tr13
		case 108:
			goto st62
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
		goto tr30
	st62:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof62
		}
	st_case_62:
		switch  lex.data[( lex.p)] {
		case 95:
			goto tr13
		case 108:
			goto tr87
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
		goto tr30
	st63:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof63
		}
	st_case_63:
		switch  lex.data[( lex.p)] {
		case 95:
			goto tr13
		case 110:
			goto tr88
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
		goto tr30
	st64:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof64
		}
	st_case_64:
		switch  lex.data[( lex.p)] {
		case 95:
			goto tr13
		case 117:
			goto st65
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
		goto tr30
	st65:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof65
		}
	st_case_65:
		switch  lex.data[( lex.p)] {
		case 95:
			goto tr13
		case 101:
			goto st66
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
		goto tr30
	st66:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof66
		}
	st_case_66:
		switch  lex.data[( lex.p)] {
		case 95:
			goto tr13
		case 114:
			goto st67
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
		goto tr30
	st67:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof67
		}
	st_case_67:
		switch  lex.data[( lex.p)] {
		case 95:
			goto tr13
		case 121:
			goto tr92
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
		goto tr30
	st68:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof68
		}
	st_case_68:
		switch  lex.data[( lex.p)] {
		case 95:
			goto tr13
		case 99:
			goto st69
		case 117:
			goto st76
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
		goto tr30
	st69:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof69
		}
	st_case_69:
		switch  lex.data[( lex.p)] {
		case 95:
			goto tr13
		case 97:
			goto st70
		case 104:
			goto st73
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
		goto tr30
	st70:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof70
		}
	st_case_70:
		switch  lex.data[( lex.p)] {
		case 95:
			goto tr13
		case 108:
			goto st71
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
		goto tr30
	st71:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof71
		}
	st_case_71:
		switch  lex.data[( lex.p)] {
		case 95:
			goto tr13
		case 97:
			goto st72
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
		goto tr30
	st72:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof72
		}
	st_case_72:
		switch  lex.data[( lex.p)] {
		case 95:
			goto tr13
		case 114:
			goto tr99
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
		goto tr30
	st73:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof73
		}
	st_case_73:
		switch  lex.data[( lex.p)] {
		case 95:
			goto tr13
		case 101:
			goto st74
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
		goto tr30
	st74:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof74
		}
	st_case_74:
		switch  lex.data[( lex.p)] {
		case 95:
			goto tr13
		case 109:
			goto st75
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
		goto tr30
	st75:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof75
		}
	st_case_75:
		switch  lex.data[( lex.p)] {
		case 95:
			goto tr13
		case 97:
			goto tr102
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
		goto tr30
	st76:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof76
		}
	st_case_76:
		switch  lex.data[( lex.p)] {
		case 95:
			goto tr13
		case 98:
			goto st77
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
		goto tr30
	st77:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof77
		}
	st_case_77:
		switch  lex.data[( lex.p)] {
		case 95:
			goto tr13
		case 115:
			goto st78
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
		goto tr30
	st78:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof78
		}
	st_case_78:
		switch  lex.data[( lex.p)] {
		case 95:
			goto tr13
		case 99:
			goto st79
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
		goto tr30
	st79:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof79
		}
	st_case_79:
		switch  lex.data[( lex.p)] {
		case 95:
			goto tr13
		case 114:
			goto st80
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
		goto tr30
	st80:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof80
		}
	st_case_80:
		switch  lex.data[( lex.p)] {
		case 95:
			goto tr13
		case 105:
			goto st81
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
		goto tr30
	st81:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof81
		}
	st_case_81:
		switch  lex.data[( lex.p)] {
		case 95:
			goto tr13
		case 112:
			goto st82
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
		goto tr30
	st82:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof82
		}
	st_case_82:
		switch  lex.data[( lex.p)] {
		case 95:
			goto tr13
		case 116:
			goto st83
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
		goto tr30
	st83:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof83
		}
	st_case_83:
		switch  lex.data[( lex.p)] {
		case 95:
			goto tr13
		case 105:
			goto st84
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
		goto tr30
	st84:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof84
		}
	st_case_84:
		switch  lex.data[( lex.p)] {
		case 95:
			goto tr13
		case 111:
			goto st85
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
		goto tr30
	st85:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof85
		}
	st_case_85:
		switch  lex.data[( lex.p)] {
		case 95:
			goto tr13
		case 110:
			goto tr112
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
		goto tr30
	st86:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof86
		}
	st_case_86:
		switch  lex.data[( lex.p)] {
		case 95:
			goto tr13
		case 114:
			goto st87
		case 121:
			goto st89
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
		goto tr30
	st87:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof87
		}
	st_case_87:
		switch  lex.data[( lex.p)] {
		case 95:
			goto tr13
		case 117:
			goto st88
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
		goto tr30
	st88:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof88
		}
	st_case_88:
		switch  lex.data[( lex.p)] {
		case 95:
			goto tr13
		case 101:
			goto tr116
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
		goto tr30
	st89:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof89
		}
	st_case_89:
		switch  lex.data[( lex.p)] {
		case 95:
			goto tr13
		case 112:
			goto st90
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
		goto tr30
	st90:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof90
		}
	st_case_90:
		switch  lex.data[( lex.p)] {
		case 95:
			goto tr13
		case 101:
			goto tr118
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
		goto tr30
	st91:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof91
		}
	st_case_91:
		switch  lex.data[( lex.p)] {
		case 95:
			goto tr13
		case 110:
			goto st92
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
		goto tr30
	st92:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof92
		}
	st_case_92:
		switch  lex.data[( lex.p)] {
		case 95:
			goto tr13
		case 105:
			goto st93
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
		goto tr30
	st93:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof93
		}
	st_case_93:
		switch  lex.data[( lex.p)] {
		case 95:
			goto tr13
		case 111:
			goto st94
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
		goto tr30
	st94:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof94
		}
	st_case_94:
		switch  lex.data[( lex.p)] {
		case 95:
			goto tr13
		case 110:
			goto tr122
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
		goto tr30
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
	_test_eof39:  lex.cs = 39; goto _test_eof
	_test_eof40:  lex.cs = 40; goto _test_eof
	_test_eof41:  lex.cs = 41; goto _test_eof
	_test_eof42:  lex.cs = 42; goto _test_eof
	_test_eof43:  lex.cs = 43; goto _test_eof
	_test_eof44:  lex.cs = 44; goto _test_eof
	_test_eof45:  lex.cs = 45; goto _test_eof
	_test_eof46:  lex.cs = 46; goto _test_eof
	_test_eof47:  lex.cs = 47; goto _test_eof
	_test_eof48:  lex.cs = 48; goto _test_eof
	_test_eof49:  lex.cs = 49; goto _test_eof
	_test_eof50:  lex.cs = 50; goto _test_eof
	_test_eof51:  lex.cs = 51; goto _test_eof
	_test_eof52:  lex.cs = 52; goto _test_eof
	_test_eof53:  lex.cs = 53; goto _test_eof
	_test_eof54:  lex.cs = 54; goto _test_eof
	_test_eof55:  lex.cs = 55; goto _test_eof
	_test_eof56:  lex.cs = 56; goto _test_eof
	_test_eof57:  lex.cs = 57; goto _test_eof
	_test_eof58:  lex.cs = 58; goto _test_eof
	_test_eof59:  lex.cs = 59; goto _test_eof
	_test_eof60:  lex.cs = 60; goto _test_eof
	_test_eof61:  lex.cs = 61; goto _test_eof
	_test_eof62:  lex.cs = 62; goto _test_eof
	_test_eof63:  lex.cs = 63; goto _test_eof
	_test_eof64:  lex.cs = 64; goto _test_eof
	_test_eof65:  lex.cs = 65; goto _test_eof
	_test_eof66:  lex.cs = 66; goto _test_eof
	_test_eof67:  lex.cs = 67; goto _test_eof
	_test_eof68:  lex.cs = 68; goto _test_eof
	_test_eof69:  lex.cs = 69; goto _test_eof
	_test_eof70:  lex.cs = 70; goto _test_eof
	_test_eof71:  lex.cs = 71; goto _test_eof
	_test_eof72:  lex.cs = 72; goto _test_eof
	_test_eof73:  lex.cs = 73; goto _test_eof
	_test_eof74:  lex.cs = 74; goto _test_eof
	_test_eof75:  lex.cs = 75; goto _test_eof
	_test_eof76:  lex.cs = 76; goto _test_eof
	_test_eof77:  lex.cs = 77; goto _test_eof
	_test_eof78:  lex.cs = 78; goto _test_eof
	_test_eof79:  lex.cs = 79; goto _test_eof
	_test_eof80:  lex.cs = 80; goto _test_eof
	_test_eof81:  lex.cs = 81; goto _test_eof
	_test_eof82:  lex.cs = 82; goto _test_eof
	_test_eof83:  lex.cs = 83; goto _test_eof
	_test_eof84:  lex.cs = 84; goto _test_eof
	_test_eof85:  lex.cs = 85; goto _test_eof
	_test_eof86:  lex.cs = 86; goto _test_eof
	_test_eof87:  lex.cs = 87; goto _test_eof
	_test_eof88:  lex.cs = 88; goto _test_eof
	_test_eof89:  lex.cs = 89; goto _test_eof
	_test_eof90:  lex.cs = 90; goto _test_eof
	_test_eof91:  lex.cs = 91; goto _test_eof
	_test_eof92:  lex.cs = 92; goto _test_eof
	_test_eof93:  lex.cs = 93; goto _test_eof
	_test_eof94:  lex.cs = 94; goto _test_eof

	_test_eof: {}
	if ( lex.p) == eof {
		switch  lex.cs {
		case 4:
			goto tr25
		case 0:
			goto tr0
		case 1:
			goto tr0
		case 5:
			goto tr26
		case 6:
			goto tr25
		case 7:
			goto tr28
		case 8:
			goto tr25
		case 2:
			goto tr0
		case 9:
			goto tr26
		case 10:
			goto tr30
		case 11:
			goto tr30
		case 12:
			goto tr30
		case 13:
			goto tr30
		case 14:
			goto tr30
		case 15:
			goto tr30
		case 16:
			goto tr30
		case 17:
			goto tr30
		case 18:
			goto tr30
		case 19:
			goto tr30
		case 20:
			goto tr30
		case 21:
			goto tr30
		case 22:
			goto tr30
		case 23:
			goto tr30
		case 24:
			goto tr30
		case 25:
			goto tr30
		case 26:
			goto tr30
		case 27:
			goto tr30
		case 28:
			goto tr30
		case 29:
			goto tr30
		case 30:
			goto tr30
		case 31:
			goto tr30
		case 32:
			goto tr30
		case 33:
			goto tr30
		case 34:
			goto tr30
		case 35:
			goto tr30
		case 36:
			goto tr30
		case 37:
			goto tr30
		case 38:
			goto tr30
		case 39:
			goto tr30
		case 40:
			goto tr30
		case 41:
			goto tr30
		case 42:
			goto tr30
		case 43:
			goto tr30
		case 44:
			goto tr30
		case 45:
			goto tr30
		case 46:
			goto tr30
		case 47:
			goto tr30
		case 48:
			goto tr30
		case 49:
			goto tr30
		case 50:
			goto tr30
		case 51:
			goto tr30
		case 52:
			goto tr30
		case 53:
			goto tr30
		case 54:
			goto tr30
		case 55:
			goto tr30
		case 56:
			goto tr30
		case 57:
			goto tr30
		case 58:
			goto tr30
		case 59:
			goto tr30
		case 60:
			goto tr30
		case 61:
			goto tr30
		case 62:
			goto tr30
		case 63:
			goto tr30
		case 64:
			goto tr30
		case 65:
			goto tr30
		case 66:
			goto tr30
		case 67:
			goto tr30
		case 68:
			goto tr30
		case 69:
			goto tr30
		case 70:
			goto tr30
		case 71:
			goto tr30
		case 72:
			goto tr30
		case 73:
			goto tr30
		case 74:
			goto tr30
		case 75:
			goto tr30
		case 76:
			goto tr30
		case 77:
			goto tr30
		case 78:
			goto tr30
		case 79:
			goto tr30
		case 80:
			goto tr30
		case 81:
			goto tr30
		case 82:
			goto tr30
		case 83:
			goto tr30
		case 84:
			goto tr30
		case 85:
			goto tr30
		case 86:
			goto tr30
		case 87:
			goto tr30
		case 88:
			goto tr30
		case 89:
			goto tr30
		case 90:
			goto tr30
		case 91:
			goto tr30
		case 92:
			goto tr30
		case 93:
			goto tr30
		case 94:
			goto tr30
		}
	}

	_out: {}
	}

//line lexer.rl:111


    return tok;
}

func (lex *lexer) Error(e string) {
    fmt.Println("errorx:", e)
    fmt.Println("t", string(lex.data[lex.ts:lex.te]))
    fmt.Println("p", string(lex.data[lex.p:lex.pe]))
}
