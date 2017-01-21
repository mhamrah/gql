//line graphql_parser.y:2
package query

import __yyfmt__ "fmt"

//line graphql_parser.y:2
import (
	"reflect"
)

//line graphql_parser.y:9
type yySymType struct {
	yys int
	str string
	val reflect.Value
	doc Document

	definitions         []Definition
	definition          Definition
	operationDefinition Operation
	opType              OpType
	selectionSet        []Selection
	selection           Selection
	field               Field
	arguments           []Argument
	argument            Argument
	fragment            Fragment
	directive           Directive
	directives          []Directive
	variables           []Variable
	variable            Variable
	fragmentSpread      FragmentSpread
}

const NAME = 57346
const VALUE = 57347
const QUERY = 57348
const MUTATION = 57349
const SUBSCRIPTION = 57350
const FRAGMENT = 57351
const SPREAD = 57352
const ON = 57353
const VARIABLE = 57354
const FRAGMENT_NAME = 57355
const type_name = 57356

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"NAME",
	"VALUE",
	"QUERY",
	"MUTATION",
	"SUBSCRIPTION",
	"FRAGMENT",
	"SPREAD",
	"ON",
	"VARIABLE",
	"FRAGMENT_NAME",
	"type_name",
	"'('",
	"')'",
	"':'",
	"'='",
	"'{'",
	"'}'",
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
	-1, 34,
	21, 46,
	-2, 47,
}

const yyNprod = 49
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 83

var yyAct = [...]int{

	63, 62, 39, 6, 21, 36, 35, 57, 10, 11,
	12, 8, 28, 9, 17, 35, 9, 19, 29, 30,
	32, 9, 20, 75, 41, 41, 9, 38, 23, 34,
	71, 66, 45, 58, 44, 48, 31, 22, 37, 46,
	43, 50, 42, 64, 52, 56, 15, 54, 78, 55,
	77, 60, 59, 69, 61, 51, 49, 14, 27, 26,
	67, 74, 73, 68, 47, 33, 70, 18, 72, 5,
	13, 65, 76, 53, 40, 25, 24, 16, 7, 3,
	4, 2, 1,
}
var yyPact = [...]int{

	-1000, -1000, 2, -1000, -1000, -1000, -1000, 53, 33, -1000,
	-1000, -1000, -1000, 7, -1000, 26, 8, -1000, -3, -3,
	-1000, -15, 24, -1000, -1000, -1000, -1000, -1000, 10, 29,
	-1000, -3, -1000, 23, -1000, 52, -1000, -1000, 51, -1000,
	-1000, -1000, -1000, 24, -6, -1000, -1000, -1000, 16, 9,
	-6, 9, -6, 27, -15, -1000, -1000, -1000, 49, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, 13, -6, 5, -1000,
	-6, 45, -1000, -1000, -1000, 43, -1000, -1000, -1000,
}
var yyPgo = [...]int{

	0, 82, 81, 80, 79, 78, 0, 1, 77, 76,
	75, 2, 74, 73, 71, 70, 69, 7, 67, 65,
	64, 63, 62, 61, 59, 58, 4, 17, 5,
}
var yyR1 = [...]int{

	0, 1, 15, 15, 2, 2, 4, 4, 3, 3,
	3, 3, 3, 5, 5, 5, 18, 19, 19, 20,
	22, 22, 23, 6, 7, 7, 8, 8, 9, 9,
	9, 10, 10, 12, 11, 11, 13, 13, 14, 24,
	25, 25, 16, 28, 21, 26, 26, 27, 17,
}
var yyR2 = [...]int{

	0, 1, 0, 1, 0, 2, 1, 1, 1, 3,
	4, 4, 5, 1, 1, 1, 3, 0, 2, 4,
	0, 1, 2, 3, 0, 1, 0, 2, 1, 1,
	1, 6, 4, 3, 0, 1, 0, 2, 3, 3,
	5, 3, 6, 1, 1, 0, 2, 2, 3,
}
var yyChk = [...]int{

	-1000, -1, -2, -4, -3, -16, -6, -5, 9, 19,
	6, 7, 8, -15, 4, 13, -8, -6, -18, -27,
	15, -26, 11, 20, -9, -10, -24, -25, 4, 10,
	-6, -27, -6, -19, -17, 21, -28, 14, 17, -11,
	-12, 15, 13, 11, -26, -6, 16, -20, 12, 4,
	-26, 4, -26, -13, -26, -28, -6, -17, 17, -11,
	-6, -11, -7, -6, 16, -14, 4, -26, -21, 4,
	-26, 17, -6, -22, -23, 18, -7, 5, 5,
}
var yyDef = [...]int{

	4, -2, 1, 5, 6, 7, 8, 2, 0, 26,
	13, 14, 15, 45, 3, 0, 0, 9, 45, 0,
	17, 0, 0, 23, 27, 28, 29, 30, 34, 45,
	10, 0, 11, 0, -2, 0, 45, 43, 0, 45,
	35, 36, 45, 0, 0, 12, 16, 18, 0, 34,
	0, 34, 24, 0, 39, 45, 41, 46, 0, 48,
	42, 45, 32, 25, 33, 37, 0, 0, 20, 44,
	24, 0, 40, 19, 21, 0, 31, 38, 22,
}
var yyTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	15, 16, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 17, 3,
	3, 18, 3, 3, 21, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 19, 3, 20,
}
var yyTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14,
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
		//line graphql_parser.y:76
		{
			yyVAL.doc = Document{Definitions: yyDollar[1].definitions}
			yylex.(*lexer).doc = yyVAL.doc
			return 0
		}
	case 2:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line graphql_parser.y:84
		{
			yyVAL.str = ""
		}
	case 3:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line graphql_parser.y:85
		{
			yyVAL.str = yyDollar[1].str
		}
	case 4:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line graphql_parser.y:89
		{
			yyVAL.definitions = nil
		}
	case 5:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line graphql_parser.y:90
		{
			yyVAL.definitions = append(yyDollar[1].definitions, yyDollar[2].definition)
		}
	case 6:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line graphql_parser.y:94
		{
			yyVAL.definition = Definition{operation: yyDollar[1].operationDefinition}
		}
	case 7:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line graphql_parser.y:95
		{
			yyVAL.definition = Definition{fragment: yyDollar[1].fragment}
		}
	case 8:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line graphql_parser.y:101
		{
			yyVAL.operationDefinition = Operation{OpType: Query, SelectionSet: yyDollar[1].selectionSet}
		}
	case 9:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line graphql_parser.y:105
		{
			yyVAL.operationDefinition = Operation{OpType: yyDollar[1].opType, Name: yyDollar[2].str, SelectionSet: yyDollar[3].selectionSet}
		}
	case 10:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line graphql_parser.y:109
		{
			yyVAL.operationDefinition = Operation{OpType: yyDollar[1].opType, Name: yyDollar[2].str, Variables: yyDollar[3].variables, SelectionSet: yyDollar[4].selectionSet}
		}
	case 11:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line graphql_parser.y:113
		{
			yyVAL.operationDefinition = Operation{OpType: yyDollar[1].opType, Name: yyDollar[2].str, Directives: yyDollar[3].directives, SelectionSet: yyDollar[4].selectionSet}
		}
	case 12:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line graphql_parser.y:117
		{
			yyVAL.operationDefinition = Operation{OpType: yyDollar[1].opType, Name: yyDollar[2].str, Variables: yyDollar[3].variables, Directives: yyDollar[4].directives, SelectionSet: yyDollar[5].selectionSet}
		}
	case 13:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line graphql_parser.y:123
		{
			yyVAL.opType = Query
		}
	case 14:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line graphql_parser.y:124
		{
			yyVAL.opType = Mutation
		}
	case 15:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line graphql_parser.y:125
		{
			yyVAL.opType = Subscription
		}
	case 16:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line graphql_parser.y:129
		{
			yyVAL.variables = yyDollar[2].variables
		}
	case 17:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line graphql_parser.y:133
		{
			yyVAL.variables = nil
		}
	case 18:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line graphql_parser.y:134
		{
			yyVAL.variables = append(yyDollar[1].variables, yyDollar[2].variable)
		}
	case 19:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line graphql_parser.y:139
		{
			yyVAL.variable = Variable{Name: yyDollar[1].str, Type: yyDollar[3].str, Default: yyDollar[4].val}
		}
	case 20:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line graphql_parser.y:145
		{
			yyVAL.val = reflect.ValueOf(nil)
		}
	case 21:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line graphql_parser.y:146
		{
			yyVAL.val = yyDollar[1].val
		}
	case 22:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line graphql_parser.y:150
		{
			yyVAL.val = yyDollar[2].val
		}
	case 23:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line graphql_parser.y:154
		{
			yyVAL.selectionSet = yyDollar[2].selectionSet
		}
	case 24:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line graphql_parser.y:158
		{
			yyVAL.selectionSet = nil
		}
	case 26:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line graphql_parser.y:163
		{
			yyVAL.selectionSet = nil
		}
	case 27:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line graphql_parser.y:164
		{
			yyVAL.selectionSet = append(yyDollar[1].selectionSet, yyDollar[2].selection)
		}
	case 28:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line graphql_parser.y:168
		{
			yyVAL.selection = Selection{Field: yyDollar[1].field}
		}
	case 29:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line graphql_parser.y:169
		{
			yyVAL.selection = Selection{FragmentSpread: yyDollar[1].fragmentSpread}
		}
	case 30:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line graphql_parser.y:170
		{
			yyVAL.selection = Selection{InlineFragment: yyDollar[1].fragment}
		}
	case 31:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line graphql_parser.y:174
		{
			yyVAL.field = Field{
				Name:         yyDollar[1].str,
				Arguments:    yyDollar[4].arguments,
				Directives:   yyDollar[5].directives,
				SelectionSet: yyDollar[6].selectionSet,
				Alias:        yyDollar[3].str,
			}
		}
	case 32:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line graphql_parser.y:183
		{
			yyVAL.field = Field{
				Name:         yyDollar[1].str,
				Arguments:    yyDollar[2].arguments,
				Directives:   yyDollar[3].directives,
				SelectionSet: yyDollar[4].selectionSet,
			}
		}
	case 33:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line graphql_parser.y:195
		{
			yyVAL.arguments = yyDollar[2].arguments
		}
	case 34:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line graphql_parser.y:199
		{
			yyVAL.arguments = nil
		}
	case 35:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line graphql_parser.y:200
		{
			yyVAL.arguments = yyDollar[1].arguments
		}
	case 36:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line graphql_parser.y:204
		{
			yyVAL.arguments = nil
		}
	case 37:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line graphql_parser.y:205
		{
			yyVAL.arguments = append(yyDollar[1].arguments, yyDollar[2].argument)
		}
	case 38:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line graphql_parser.y:209
		{
			yyVAL.argument = Argument{Name: yyDollar[1].str, Value: yyDollar[3].val}
		}
	case 39:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line graphql_parser.y:214
		{
			yyVAL.fragmentSpread = FragmentSpread{Name: yyDollar[2].str, Directives: yyDollar[3].directives}
		}
	case 40:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line graphql_parser.y:219
		{
			yyVAL.fragment = Fragment{
				TypeCondition: yyDollar[3].str,
				Directives:    yyDollar[4].directives,
				SelectionSet:  yyDollar[5].selectionSet,
			}
		}
	case 41:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line graphql_parser.y:227
		{
			yyVAL.fragment = Fragment{
				Directives:   yyDollar[2].directives,
				SelectionSet: yyDollar[3].selectionSet,
			}
		}
	case 42:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line graphql_parser.y:237
		{
			yyVAL.fragment = Fragment{
				Name:          yyDollar[2].str,
				TypeCondition: yyDollar[4].str,
				Directives:    yyDollar[5].directives,
				SelectionSet:  yyDollar[6].selectionSet,
			}
		}
	case 43:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line graphql_parser.y:248
		{
			yyVAL.str = yyDollar[1].str
		}
	case 44:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line graphql_parser.y:256
		{
			yyVAL.str = yyDollar[1].str
		}
	case 45:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line graphql_parser.y:262
		{
			yyVAL.directives = nil
		}
	case 46:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line graphql_parser.y:263
		{
			yyVAL.directives = append(yyDollar[1].directives, yyDollar[2].directive)
		}
	case 47:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line graphql_parser.y:267
		{
			yyVAL.directives = append(yyDollar[1].directives, yyDollar[2].directive)
		}
	case 48:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line graphql_parser.y:271
		{
			yyVAL.directive = Directive{Name: yyDollar[2].str, Arguments: yyDollar[3].arguments}
		}
	}
	goto yystack /* stack new state and value */
}
