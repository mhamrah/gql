//line graphql_parser.y:2
package query

import __yyfmt__ "fmt"

//line graphql_parser.y:2
import (
	"fmt"
	"reflect"
)

//line graphql_parser.y:10
type yySymType struct {
	yys  int
	str  string
	val  reflect.Value
	vals []reflect.Value
	doc  Document

	definitions []Definition
	definition  Definition

	directives []Directive
	directive  Directive

	arguments []Argument
	argument  Argument

	variables []Variable
	variable  Variable

	selectionSet []Selection
	selection    Selection

	operation      Operation
	opType         OpType
	field          Field
	fragment       Fragment
	fragmentSpread FragmentSpread
	objectField    ObjectField
	objectFields   []ObjectField
}

const DIRECTIVE = 57346
const ENUM = 57347
const EXTEND = 57348
const FALSE = 57349
const FRAGMENT = 57350
const IMPLEMENTS = 57351
const INPUT = 57352
const INTERFACE = 57353
const MUTATION = 57354
const NULL = 57355
const QUERY = 57356
const ON = 57357
const SCALAR = 57358
const SCHEMA = 57359
const SUBSCRIPTION = 57360
const TRUE = 57361
const TYPE = 57362
const UNION = 57363
const VARIABLE = 57364
const IDENTIFIER = 57365
const VALUE = 57366
const SPREAD = 57367
const FRAGMENT_NAME = 57368

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"DIRECTIVE",
	"ENUM",
	"EXTEND",
	"FALSE",
	"FRAGMENT",
	"IMPLEMENTS",
	"INPUT",
	"INTERFACE",
	"MUTATION",
	"NULL",
	"QUERY",
	"ON",
	"SCALAR",
	"SCHEMA",
	"SUBSCRIPTION",
	"TRUE",
	"TYPE",
	"UNION",
	"VARIABLE",
	"IDENTIFIER",
	"VALUE",
	"SPREAD",
	"FRAGMENT_NAME",
	"'('",
	"')'",
	"':'",
	"'='",
	"'{'",
	"'}'",
	"'['",
	"']'",
	"'@'",
}
var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 55,
	35, 100,
	-2, 101,
}

const yyNprod = 103
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 302

var yyAct = [...]int{

	99, 59, 84, 83, 58, 57, 7, 60, 79, 42,
	15, 10, 40, 56, 10, 56, 97, 38, 110, 111,
	112, 108, 113, 115, 116, 117, 118, 106, 120, 119,
	121, 122, 123, 107, 124, 125, 102, 114, 100, 49,
	16, 135, 51, 53, 63, 126, 61, 109, 130, 41,
	36, 55, 52, 10, 93, 67, 80, 63, 71, 70,
	66, 43, 105, 74, 129, 68, 133, 72, 48, 78,
	73, 77, 47, 103, 76, 82, 128, 88, 101, 81,
	96, 95, 85, 104, 90, 91, 69, 89, 54, 39,
	6, 64, 94, 87, 75, 92, 98, 62, 127, 110,
	111, 112, 108, 113, 115, 116, 117, 118, 106, 120,
	119, 121, 122, 123, 107, 124, 125, 102, 114, 100,
	46, 45, 37, 8, 4, 5, 126, 3, 109, 131,
	14, 134, 2, 1, 0, 0, 136, 18, 19, 20,
	21, 22, 24, 25, 26, 27, 28, 29, 17, 30,
	31, 32, 33, 34, 35, 0, 23, 0, 50, 0,
	0, 0, 0, 0, 0, 44, 18, 19, 20, 21,
	22, 24, 25, 26, 27, 28, 29, 17, 30, 31,
	32, 33, 34, 35, 0, 23, 0, 0, 0, 0,
	0, 0, 0, 0, 132, 18, 19, 20, 21, 22,
	24, 25, 26, 27, 28, 29, 17, 30, 31, 32,
	33, 34, 35, 0, 23, 0, 0, 0, 0, 86,
	18, 19, 20, 21, 22, 24, 25, 26, 27, 28,
	29, 17, 30, 31, 32, 33, 34, 35, 0, 23,
	18, 19, 20, 21, 22, 24, 25, 26, 27, 28,
	29, 65, 30, 31, 32, 33, 34, 35, 0, 23,
	18, 19, 20, 21, 22, 24, 25, 26, 27, 28,
	29, 0, 30, 31, 32, 33, 34, 35, 9, 23,
	0, 0, 12, 0, 11, 0, 0, 0, 13, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 10,
}
var yyPact = [...]int{

	-1000, -1000, -1000, 270, -1000, -1000, -1000, -1000, 216, 256,
	-1000, -1000, -1000, -1000, 22, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, 46, 133, -1000, -17,
	-17, -1000, -22, 216, -1000, -1000, -1000, -1000, -1000, 17,
	236, -1000, -17, -1000, 37, -1000, 216, -1000, -1000, -1000,
	-1000, 216, -1000, -1000, -1000, 216, -20, -1000, -1000, -1000,
	27, 30, -20, -20, 30, 191, -22, -1000, -1000, -1000,
	216, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 25, -20,
	-14, -1000, -20, 95, -1000, -1000, -1000, 95, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 14, 162,
	-1000, -1000, -1000, -1000, 12, 95, -1000,
}
var yyPgo = [...]int{

	0, 133, 132, 40, 1, 130, 127, 125, 124, 123,
	2, 3, 122, 121, 120, 7, 97, 94, 93, 90,
	8, 89, 88, 86, 84, 4, 83, 81, 80, 0,
	78, 76, 73, 72, 68, 9, 12, 5, 66, 64,
	62,
}
var yyR1 = [...]int{

	0, 1, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	4, 4, 5, 5, 2, 6, 6, 8, 8, 7,
	7, 7, 7, 7, 9, 9, 9, 21, 22, 22,
	23, 27, 27, 28, 10, 11, 11, 12, 12, 13,
	13, 13, 14, 14, 16, 15, 15, 17, 17, 18,
	33, 34, 34, 19, 37, 29, 29, 29, 29, 29,
	29, 29, 30, 30, 26, 26, 26, 26, 26, 26,
	26, 26, 26, 26, 26, 26, 26, 26, 26, 26,
	32, 31, 31, 40, 39, 39, 38, 24, 25, 35,
	35, 36, 20,
}
var yyR2 = [...]int{

	0, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 0, 1, 1, 0, 2, 1, 1, 1,
	3, 4, 4, 5, 1, 1, 1, 3, 0, 2,
	4, 0, 1, 2, 3, 0, 1, 0, 2, 1,
	1, 1, 4, 6, 3, 0, 1, 0, 2, 3,
	3, 5, 3, 6, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	3, 0, 2, 3, 0, 2, 3, 1, 1, 0,
	2, 2, 3,
}
var yyChk = [...]int{

	-1000, -1, -2, -6, -8, -7, -19, -10, -9, 8,
	31, 14, 12, 18, -5, -4, -3, 15, 4, 5,
	6, 7, 8, 23, 9, 10, 11, 12, 13, 14,
	16, 17, 18, 19, 20, 21, -3, -12, -10, -21,
	-36, 27, -35, 15, 32, -13, -14, -33, -34, -4,
	25, -10, -36, -10, -22, -20, 35, -37, -25, -4,
	-15, 29, -16, 27, -3, 15, -35, -10, 28, -23,
	22, -4, -35, -35, -4, -17, -35, -37, -10, -20,
	29, -15, -10, -11, -10, -15, 28, -18, -4, -35,
	-24, -25, -35, 29, -10, -27, -28, 30, -11, -29,
	24, -30, 22, -32, -26, -40, 13, 19, 7, 33,
	4, 5, 6, 8, 23, 9, 10, 11, 12, 15,
	14, 16, 17, 18, 20, 21, 31, -29, -31, -39,
	34, -29, 32, -38, -4, 29, -29,
}
var yyDef = [...]int{

	25, -2, 1, 24, 26, 27, 28, 29, 22, 0,
	47, 34, 35, 36, 99, 23, 20, 21, 2, 3,
	4, 5, 6, 7, 8, 9, 10, 11, 12, 13,
	14, 15, 16, 17, 18, 19, 0, 0, 30, 99,
	0, 38, 0, 0, 44, 48, 49, 50, 51, 55,
	99, 31, 0, 32, 0, -2, 0, 99, 64, 98,
	99, 0, 56, 57, 99, 0, 0, 33, 37, 39,
	0, 55, 0, 45, 55, 0, 60, 99, 62, 100,
	0, 102, 63, 52, 46, 99, 54, 58, 0, 0,
	41, 97, 45, 0, 61, 40, 42, 0, 53, 59,
	65, 66, 67, 68, 69, 70, 71, 72, 73, 91,
	74, 75, 76, 77, 78, 79, 80, 81, 82, 83,
	84, 85, 86, 87, 88, 89, 94, 43, 0, 0,
	90, 92, 93, 95, 0, 0, 96,
}
var yyTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	27, 28, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 29, 3,
	3, 30, 3, 3, 35, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 33, 3, 34, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 31, 3, 32,
}
var yyTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26,
}
var yyTok3 = [...]int{
	0,
}

var yyErrorMessages = [...]struct {
	state int
	token int
	msg   string
}{}

//line yaccpar:1

/*	parser for yacc output	*/

var (
	yyDebug        = 0
	yyErrorVerbose = false
)

type yyLexer interface {
	Lex(lval *yySymType) int
	Error(s string)
}

type yyParser interface {
	Parse(yyLexer) int
	Lookahead() int
}

type yyParserImpl struct {
	lval  yySymType
	stack [yyInitialStackSize]yySymType
	char  int
}

func (p *yyParserImpl) Lookahead() int {
	return p.char
}

func yyNewParser() yyParser {
	return &yyParserImpl{}
}

const yyFlag = -1000

func yyTokname(c int) string {
	if c >= 1 && c-1 < len(yyToknames) {
		if yyToknames[c-1] != "" {
			return yyToknames[c-1]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func yyStatname(s int) string {
	if s >= 0 && s < len(yyStatenames) {
		if yyStatenames[s] != "" {
			return yyStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func yyErrorMessage(state, lookAhead int) string {
	const TOKSTART = 4

	if !yyErrorVerbose {
		return "syntax error"
	}

	for _, e := range yyErrorMessages {
		if e.state == state && e.token == lookAhead {
			return "syntax error: " + e.msg
		}
	}

	res := "syntax error: unexpected " + yyTokname(lookAhead)

	// To match Bison, suggest at most four expected tokens.
	expected := make([]int, 0, 4)

	// Look for shiftable tokens.
	base := yyPact[state]
	for tok := TOKSTART; tok-1 < len(yyToknames); tok++ {
		if n := base + tok; n >= 0 && n < yyLast && yyChk[yyAct[n]] == tok {
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}
	}

	if yyDef[state] == -2 {
		i := 0
		for yyExca[i] != -1 || yyExca[i+1] != state {
			i += 2
		}

		// Look for tokens that we accept or reduce.
		for i += 2; yyExca[i] >= 0; i += 2 {
			tok := yyExca[i]
			if tok < TOKSTART || yyExca[i+1] == 0 {
				continue
			}
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}

		// If the default action is to accept or reduce, give up.
		if yyExca[i+1] != 0 {
			return res
		}
	}

	for i, tok := range expected {
		if i == 0 {
			res += ", expecting "
		} else {
			res += " or "
		}
		res += yyTokname(tok)
	}
	return res
}

func yylex1(lex yyLexer, lval *yySymType) (char, token int) {
	token = 0
	char = lex.Lex(lval)
	if char <= 0 {
		token = yyTok1[0]
		goto out
	}
	if char < len(yyTok1) {
		token = yyTok1[char]
		goto out
	}
	if char >= yyPrivate {
		if char < yyPrivate+len(yyTok2) {
			token = yyTok2[char-yyPrivate]
			goto out
		}
	}
	for i := 0; i < len(yyTok3); i += 2 {
		token = yyTok3[i+0]
		if token == char {
			token = yyTok3[i+1]
			goto out
		}
	}

out:
	if token == 0 {
		token = yyTok2[1] /* unknown char */
	}
	if yyDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", yyTokname(token), uint(char))
	}
	return char, token
}

func yyParse(yylex yyLexer) int {
	return yyNewParser().Parse(yylex)
}

func (yyrcvr *yyParserImpl) Parse(yylex yyLexer) int {
	var yyn int
	var yyVAL yySymType
	var yyDollar []yySymType
	_ = yyDollar // silence set and not used
	yyS := yyrcvr.stack[:]

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yystate := 0
	yyrcvr.char = -1
	yytoken := -1 // yyrcvr.char translated into internal numbering
	defer func() {
		// Make sure we report no lookahead when not parsing.
		yystate = -1
		yyrcvr.char = -1
		yytoken = -1
	}()
	yyp := -1
	goto yystack

ret0:
	return 0

ret1:
	return 1

yystack:
	/* put a state and value onto the stack */
	if yyDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", yyTokname(yytoken), yyStatname(yystate))
	}

	yyp++
	if yyp >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyS[yyp] = yyVAL
	yyS[yyp].yys = yystate

yynewstate:
	yyn = yyPact[yystate]
	if yyn <= yyFlag {
		goto yydefault /* simple state */
	}
	if yyrcvr.char < 0 {
		yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
	}
	yyn += yytoken
	if yyn < 0 || yyn >= yyLast {
		goto yydefault
	}
	yyn = yyAct[yyn]
	if yyChk[yyn] == yytoken { /* valid shift */
		yyrcvr.char = -1
		yytoken = -1
		yyVAL = yyrcvr.lval
		yystate = yyn
		if Errflag > 0 {
			Errflag--
		}
		goto yystack
	}

yydefault:
	/* default state action */
	yyn = yyDef[yystate]
	if yyn == -2 {
		if yyrcvr.char < 0 {
			yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
		}

		/* look through exception table */
		xi := 0
		for {
			if yyExca[xi+0] == -1 && yyExca[xi+1] == yystate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			yyn = yyExca[xi+0]
			if yyn < 0 || yyn == yytoken {
				break
			}
		}
		yyn = yyExca[xi+1]
		if yyn < 0 {
			goto ret0
		}
	}
	if yyn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			yylex.Error(yyErrorMessage(yystate, yytoken))
			Nerrs++
			if yyDebug >= 1 {
				__yyfmt__.Printf("%s", yyStatname(yystate))
				__yyfmt__.Printf(" saw %s\n", yyTokname(yytoken))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for yyp >= 0 {
				yyn = yyPact[yyS[yyp].yys] + yyErrCode
				if yyn >= 0 && yyn < yyLast {
					yystate = yyAct[yyn] /* simulate a shift of "error" */
					if yyChk[yystate] == yyErrCode {
						goto yystack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if yyDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", yyS[yyp].yys)
				}
				yyp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if yyDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", yyTokname(yytoken))
			}
			if yytoken == yyEofCode {
				goto ret1
			}
			yyrcvr.char = -1
			yytoken = -1
			goto yynewstate /* try again in the same state */
		}
	}

	/* reduction by production yyn */
	if yyDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", yyn, yyStatname(yystate))
	}

	yynt := yyn
	yypt := yyp
	_ = yypt // guard against "declared and not used"

	yyp -= yyR2[yyn]
	// yyp is now the index of $0. Perform the default action. Iff the
	// reduced production is ε, $1 is possibly out of range.
	if yyp+1 >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyVAL = yyS[yyp+1]

	/* consult goto table to find next state */
	yyn = yyR1[yyn]
	yyg := yyPgo[yyn]
	yyj := yyg + yyS[yyp].yys + 1

	if yyj >= yyLast {
		yystate = yyAct[yyg]
	} else {
		yystate = yyAct[yyj]
		if yyChk[yystate] != -yyn {
			yystate = yyAct[yyg]
		}
	}
	// dummy call; replaced with literal code
	switch yynt {

	case 1:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line graphql_parser.y:112
		{
			fmt.Println("doc:", yyDollar[1].doc)
			fmt.Println("err:", yylex.(*lexer).err)
			fmt.Println("failed:", yylex.(*lexer).parseFailed)
			// TODO check error?
			yylex.(*lexer).doc = yyDollar[1].doc
			return 0
		}
	case 2:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line graphql_parser.y:123
		{
			yyVAL.str = yyDollar[1].str
		}
	case 3:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line graphql_parser.y:124
		{
			yyVAL.str = yyDollar[1].str
		}
	case 4:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line graphql_parser.y:125
		{
			yyVAL.str = yyDollar[1].str
		}
	case 5:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line graphql_parser.y:126
		{
			yyVAL.str = yyDollar[1].str
		}
	case 6:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line graphql_parser.y:127
		{
			yyVAL.str = yyDollar[1].str
		}
	case 7:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line graphql_parser.y:128
		{
			yyVAL.str = yyDollar[1].str
		}
	case 8:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line graphql_parser.y:129
		{
			yyVAL.str = yyDollar[1].str
		}
	case 9:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line graphql_parser.y:130
		{
			yyVAL.str = yyDollar[1].str
		}
	case 10:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line graphql_parser.y:131
		{
			yyVAL.str = yyDollar[1].str
		}
	case 11:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line graphql_parser.y:132
		{
			yyVAL.str = yyDollar[1].str
		}
	case 12:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line graphql_parser.y:133
		{
			yyVAL.str = yyDollar[1].str
		}
	case 13:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line graphql_parser.y:134
		{
			yyVAL.str = yyDollar[1].str
		}
	case 14:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line graphql_parser.y:135
		{
			yyVAL.str = yyDollar[1].str
		}
	case 15:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line graphql_parser.y:136
		{
			yyVAL.str = yyDollar[1].str
		}
	case 16:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line graphql_parser.y:137
		{
			yyVAL.str = yyDollar[1].str
		}
	case 17:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line graphql_parser.y:138
		{
			yyVAL.str = yyDollar[1].str
		}
	case 18:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line graphql_parser.y:139
		{
			yyVAL.str = yyDollar[1].str
		}
	case 19:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line graphql_parser.y:140
		{
			yyVAL.str = yyDollar[1].str
		}
	case 20:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line graphql_parser.y:143
		{
			yyVAL.str = yyDollar[1].str
		}
	case 21:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line graphql_parser.y:144
		{
			yyVAL.str = yyDollar[1].str
		}
	case 22:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line graphql_parser.y:147
		{
			yyVAL.str = ""
		}
	case 24:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line graphql_parser.y:155
		{
			yyVAL.doc = Document{Definitions: yyDollar[1].definitions}
		}
	case 25:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line graphql_parser.y:162
		{
			yyVAL.definitions = nil
		}
	case 26:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line graphql_parser.y:163
		{
			yyVAL.definitions = append(yyDollar[1].definitions, yyDollar[2].definition)
		}
	case 27:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line graphql_parser.y:167
		{
			yyVAL.definition = Definition{operation: yyDollar[1].operation}
		}
	case 28:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line graphql_parser.y:168
		{
			yyVAL.definition = Definition{fragment: yyDollar[1].fragment}
		}
	case 29:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line graphql_parser.y:174
		{
			yyVAL.operation = Operation{OpType: Query, SelectionSet: yyDollar[1].selectionSet}
		}
	case 30:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line graphql_parser.y:178
		{
			yyVAL.operation = Operation{OpType: yyDollar[1].opType, Name: yyDollar[2].str, SelectionSet: yyDollar[3].selectionSet}
		}
	case 31:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line graphql_parser.y:182
		{
			yyVAL.operation = Operation{OpType: yyDollar[1].opType, Name: yyDollar[2].str, Variables: yyDollar[3].variables, SelectionSet: yyDollar[4].selectionSet}
		}
	case 32:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line graphql_parser.y:186
		{
			yyVAL.operation = Operation{OpType: yyDollar[1].opType, Name: yyDollar[2].str, Directives: yyDollar[3].directives, SelectionSet: yyDollar[4].selectionSet}
		}
	case 33:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line graphql_parser.y:190
		{
			yyVAL.operation = Operation{OpType: yyDollar[1].opType, Name: yyDollar[2].str, Variables: yyDollar[3].variables, Directives: yyDollar[4].directives, SelectionSet: yyDollar[5].selectionSet}
		}
	case 34:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line graphql_parser.y:196
		{
			yyVAL.opType = Query
		}
	case 35:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line graphql_parser.y:197
		{
			yyVAL.opType = Mutation
		}
	case 36:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line graphql_parser.y:198
		{
			yyVAL.opType = Subscription
		}
	case 37:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line graphql_parser.y:202
		{
			yyVAL.variables = yyDollar[2].variables
		}
	case 38:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line graphql_parser.y:206
		{
			yyVAL.variables = nil
		}
	case 39:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line graphql_parser.y:207
		{
			yyVAL.variables = append(yyDollar[1].variables, yyDollar[2].variable)
		}
	case 40:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line graphql_parser.y:212
		{
			yyVAL.variable = Variable{Name: yyDollar[1].str, Type: yyDollar[3].str, Default: yyDollar[4].val}
		}
	case 41:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line graphql_parser.y:218
		{
			yyVAL.val = reflect.ValueOf(nil)
		}
	case 42:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line graphql_parser.y:219
		{
			yyVAL.val = yyDollar[1].val
		}
	case 43:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line graphql_parser.y:223
		{
			yyVAL.val = yyDollar[2].val
		}
	case 44:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line graphql_parser.y:226
		{
			yyVAL.selectionSet = yyDollar[2].selectionSet
		}
	case 45:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line graphql_parser.y:230
		{
			yyVAL.selectionSet = nil
		}
	case 47:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line graphql_parser.y:235
		{
			yyVAL.selectionSet = nil
		}
	case 48:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line graphql_parser.y:236
		{
			yyVAL.selectionSet = append(yyDollar[1].selectionSet, yyDollar[2].selection)
		}
	case 49:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line graphql_parser.y:240
		{
			yyVAL.selection = Selection{Field: yyDollar[1].field}
		}
	case 50:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line graphql_parser.y:241
		{
			yyVAL.selection = Selection{FragmentSpread: yyDollar[1].fragmentSpread}
		}
	case 51:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line graphql_parser.y:242
		{
			yyVAL.selection = Selection{InlineFragment: yyDollar[1].fragment}
		}
	case 52:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line graphql_parser.y:246
		{
			yyVAL.field = Field{
				Name:         yyDollar[1].str,
				Arguments:    yyDollar[2].arguments,
				Directives:   yyDollar[3].directives,
				SelectionSet: yyDollar[4].selectionSet,
			}
		}
	case 53:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line graphql_parser.y:254
		{
			yyVAL.field = Field{
				Name:         yyDollar[1].str,
				Arguments:    yyDollar[4].arguments,
				Directives:   yyDollar[5].directives,
				SelectionSet: yyDollar[6].selectionSet,
				Alias:        yyDollar[3].str,
			}
		}
	case 54:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line graphql_parser.y:266
		{
			yyVAL.arguments = yyDollar[2].arguments
		}
	case 55:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line graphql_parser.y:270
		{
			yyVAL.arguments = nil
		}
	case 56:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line graphql_parser.y:271
		{
			yyVAL.arguments = yyDollar[1].arguments
		}
	case 57:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line graphql_parser.y:275
		{
			yyVAL.arguments = nil
		}
	case 58:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line graphql_parser.y:276
		{
			yyVAL.arguments = append(yyDollar[1].arguments, yyDollar[2].argument)
		}
	case 59:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line graphql_parser.y:280
		{
			yyVAL.argument = Argument{Name: yyDollar[1].str, Value: yyDollar[3].val}
		}
	case 60:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line graphql_parser.y:285
		{
			yyVAL.fragmentSpread = FragmentSpread{Name: yyDollar[2].str, Directives: yyDollar[3].directives}
		}
	case 61:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line graphql_parser.y:290
		{
			yyVAL.fragment = Fragment{
				TypeCondition: yyDollar[3].str,
				Directives:    yyDollar[4].directives,
				SelectionSet:  yyDollar[5].selectionSet,
			}
		}
	case 62:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line graphql_parser.y:298
		{
			yyVAL.fragment = Fragment{
				Directives:   yyDollar[2].directives,
				SelectionSet: yyDollar[3].selectionSet,
			}
		}
	case 63:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line graphql_parser.y:308
		{
			yyVAL.fragment = Fragment{
				Name:          yyDollar[2].str,
				TypeCondition: yyDollar[4].str,
				Directives:    yyDollar[5].directives,
				SelectionSet:  yyDollar[6].selectionSet,
			}
		}
	case 64:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line graphql_parser.y:318
		{
			yyVAL.str = yyDollar[1].str
		}
	case 65:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line graphql_parser.y:323
		{
			yyVAL.val = yyDollar[1].val
		}
	case 66:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line graphql_parser.y:324
		{
			yyVAL.val = yyDollar[1].val
		}
	case 67:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line graphql_parser.y:325
		{
			yyVAL.val = reflect.ValueOf(yyDollar[1].str)
		}
	case 68:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line graphql_parser.y:326
		{
			yyVAL.val = reflect.ValueOf(yyDollar[1].vals)
		}
	case 69:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line graphql_parser.y:327
		{
			yyVAL.val = reflect.ValueOf(yyDollar[1].str)
		}
	case 70:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line graphql_parser.y:328
		{
			yyVAL.val = reflect.ValueOf(yyDollar[1].objectFields)
		}
	case 71:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line graphql_parser.y:329
		{
			yyVAL.val = reflect.ValueOf(nil)
		}
	case 72:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line graphql_parser.y:332
		{
			yyVAL.val = reflect.ValueOf(true)
		}
	case 73:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line graphql_parser.y:333
		{
			yyVAL.val = reflect.ValueOf(false)
		}
	case 74:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line graphql_parser.y:336
		{
			yyVAL.str = yyDollar[1].str
		}
	case 75:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line graphql_parser.y:337
		{
			yyVAL.str = yyDollar[1].str
		}
	case 76:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line graphql_parser.y:338
		{
			yyVAL.str = yyDollar[1].str
		}
	case 77:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line graphql_parser.y:339
		{
			yyVAL.str = yyDollar[1].str
		}
	case 78:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line graphql_parser.y:340
		{
			yyVAL.str = yyDollar[1].str
		}
	case 79:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line graphql_parser.y:341
		{
			yyVAL.str = yyDollar[1].str
		}
	case 80:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line graphql_parser.y:342
		{
			yyVAL.str = yyDollar[1].str
		}
	case 81:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line graphql_parser.y:343
		{
			yyVAL.str = yyDollar[1].str
		}
	case 82:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line graphql_parser.y:344
		{
			yyVAL.str = yyDollar[1].str
		}
	case 83:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line graphql_parser.y:345
		{
			yyVAL.str = yyDollar[1].str
		}
	case 84:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line graphql_parser.y:346
		{
			yyVAL.str = yyDollar[1].str
		}
	case 85:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line graphql_parser.y:347
		{
			yyVAL.str = yyDollar[1].str
		}
	case 86:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line graphql_parser.y:348
		{
			yyVAL.str = yyDollar[1].str
		}
	case 87:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line graphql_parser.y:349
		{
			yyVAL.str = yyDollar[1].str
		}
	case 88:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line graphql_parser.y:350
		{
			yyVAL.str = yyDollar[1].str
		}
	case 89:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line graphql_parser.y:351
		{
			yyVAL.str = yyDollar[1].str
		}
	case 90:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line graphql_parser.y:356
		{
			yyVAL.vals = yyDollar[2].vals
		}
	case 91:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line graphql_parser.y:359
		{
			yyVAL.vals = nil
		}
	case 92:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line graphql_parser.y:360
		{
			yyVAL.vals = append(yyDollar[1].vals, yyDollar[2].val)
		}
	case 93:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line graphql_parser.y:364
		{
			yyVAL.objectFields = yyDollar[2].objectFields
		}
	case 94:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line graphql_parser.y:367
		{
			yyVAL.objectFields = nil
		}
	case 95:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line graphql_parser.y:368
		{
			yyVAL.objectFields = append(yyDollar[1].objectFields, yyDollar[2].objectField)
		}
	case 96:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line graphql_parser.y:371
		{
			yyVAL.objectField = ObjectField{Key: yyDollar[1].str, Value: yyDollar[3].val}
		}
	case 97:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line graphql_parser.y:376
		{
			yyVAL.str = yyDollar[1].str
		}
	case 98:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line graphql_parser.y:379
		{
			yyVAL.str = yyDollar[1].str
		}
	case 99:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line graphql_parser.y:385
		{
			yyVAL.directives = nil
		}
	case 100:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line graphql_parser.y:386
		{
			yyVAL.directives = append(yyDollar[1].directives, yyDollar[2].directive)
		}
	case 101:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line graphql_parser.y:390
		{
			yyVAL.directives = append(yyDollar[1].directives, yyDollar[2].directive)
		}
	case 102:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line graphql_parser.y:394
		{
			yyVAL.directive = Directive{Name: yyDollar[2].str, Arguments: yyDollar[3].arguments}
		}
	}
	goto yystack /* stack new state and value */
}
