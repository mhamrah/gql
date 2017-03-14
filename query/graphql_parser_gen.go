//line graphql_parser.y:2
package query

import __yyfmt__ "fmt"

//line graphql_parser.y:2
import (
	"reflect"
)

//line graphql_parser.y:9
type yySymType struct {
	yys  int
	str  string
	strs []string
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

	schema                   Schema
	operationTypeDefinition  OperationTypeDefinition
	operationTypeDefinitions []OperationTypeDefinition
	typeDefinition           TypeDefinition
	typeDefinitions          []TypeDefinition
	inputDefinitions         []InputValueDefinition
	inputDefinition          InputValueDefinition
	fieldDef                 FieldDefinition
	fieldsDef                []FieldDefinition
	typeDesc                 TypeDescription
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
	"'!'",
	"'@'",
	"'|'",
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
	-1, 83,
	36, 106,
	-2, 107,
}

const yyNprod = 145
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 586

var yyAct = [...]int{

	86, 168, 62, 175, 144, 143, 148, 136, 147, 18,
	141, 126, 20, 135, 87, 84, 13, 11, 55, 161,
	21, 57, 60, 133, 120, 127, 122, 121, 57, 10,
	57, 57, 57, 41, 119, 56, 57, 172, 13, 57,
	57, 171, 15, 69, 14, 58, 209, 170, 16, 222,
	71, 72, 73, 74, 75, 76, 61, 12, 78, 160,
	13, 15, 101, 14, 90, 204, 88, 16, 167, 153,
	85, 124, 123, 93, 94, 90, 97, 98, 99, 100,
	13, 83, 80, 6, 63, 79, 81, 109, 96, 111,
	110, 91, 48, 104, 113, 103, 108, 140, 116, 85,
	47, 138, 106, 118, 46, 45, 105, 159, 114, 158,
	117, 95, 44, 131, 17, 43, 5, 132, 77, 115,
	137, 102, 142, 145, 134, 4, 128, 139, 181, 211,
	146, 152, 218, 68, 137, 125, 137, 67, 179, 210,
	177, 142, 169, 156, 145, 164, 180, 155, 165, 150,
	149, 163, 107, 82, 59, 9, 137, 130, 154, 112,
	173, 145, 89, 156, 174, 66, 205, 65, 206, 42,
	7, 8, 3, 19, 208, 2, 207, 1, 186, 187,
	188, 184, 189, 191, 192, 193, 194, 182, 196, 195,
	197, 198, 199, 183, 200, 201, 178, 190, 176, 0,
	0, 0, 0, 0, 0, 202, 145, 185, 215, 214,
	165, 0, 219, 212, 216, 220, 0, 221, 0, 0,
	0, 0, 0, 0, 0, 0, 223, 186, 187, 188,
	184, 189, 191, 192, 193, 194, 182, 196, 195, 197,
	198, 199, 183, 200, 201, 178, 190, 176, 0, 0,
	0, 0, 0, 0, 202, 0, 185, 23, 24, 25,
	26, 27, 29, 30, 31, 32, 33, 34, 22, 35,
	36, 37, 38, 39, 40, 0, 28, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 151, 23, 24, 25,
	26, 27, 29, 30, 31, 32, 33, 34, 22, 35,
	36, 37, 38, 39, 40, 0, 28, 0, 70, 0,
	0, 0, 0, 0, 0, 64, 23, 24, 25, 26,
	27, 29, 30, 31, 32, 33, 34, 22, 35, 36,
	37, 38, 39, 40, 0, 28, 0, 0, 0, 0,
	0, 0, 0, 0, 217, 23, 24, 25, 26, 27,
	29, 30, 31, 32, 33, 34, 22, 35, 36, 37,
	38, 39, 40, 0, 28, 0, 0, 0, 0, 0,
	0, 0, 0, 203, 23, 24, 25, 26, 27, 29,
	30, 31, 32, 33, 34, 22, 35, 36, 37, 38,
	39, 40, 0, 28, 0, 0, 0, 0, 0, 0,
	0, 0, 166, 23, 24, 25, 26, 27, 29, 30,
	31, 32, 33, 34, 22, 35, 36, 37, 38, 39,
	40, 0, 28, 0, 0, 0, 0, 0, 0, 0,
	0, 162, 23, 24, 25, 26, 27, 29, 30, 31,
	32, 33, 34, 22, 35, 36, 37, 38, 39, 40,
	0, 28, 0, 0, 0, 0, 0, 0, 0, 0,
	157, 23, 24, 25, 26, 27, 29, 30, 31, 32,
	33, 34, 22, 35, 36, 37, 38, 39, 40, 0,
	28, 0, 0, 0, 0, 213, 23, 24, 25, 26,
	27, 29, 30, 31, 32, 33, 34, 22, 35, 36,
	37, 38, 39, 40, 0, 28, 0, 0, 0, 0,
	129, 23, 24, 25, 26, 27, 29, 30, 31, 32,
	33, 34, 22, 35, 36, 37, 38, 39, 40, 0,
	28, 23, 24, 25, 26, 27, 29, 30, 31, 32,
	33, 34, 92, 35, 36, 37, 38, 39, 40, 0,
	28, 23, 24, 25, 26, 27, 29, 30, 31, 32,
	33, 34, 0, 35, 36, 37, 38, 39, 40, 53,
	28, 0, 0, 0, 54, 51, 0, 0, 0, 0,
	49, 0, 0, 0, 50, 52,
}
var yyPact = [...]int{

	66, -1000, -1000, 49, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, 507, 547, -1000, -1000, -1000, -1000, 564, 4, 29,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, 69, 283, -1000, -1000, -1000, -1000, -1000, -1000, 507,
	507, 507, 507, 507, 507, -1000, -1000, 507, -1000, 7,
	7, -1000, 0, 507, -1000, -1000, -1000, -1000, -1000, 37,
	527, -1000, 79, -1000, -1000, -1000, -1000, 30, 48, -1000,
	7, -1000, 74, -1000, -1000, -1000, -1000, -1000, 507, -1000,
	-1000, -1000, 507, -15, 0, -1000, 507, 3, -6, -4,
	-5, -1000, -1000, 43, -1000, -1000, -1000, -1000, 42, -15,
	-15, 48, 482, 0, -1000, -1000, -8, 507, -1000, 507,
	507, 507, 507, 507, 253, -1000, -1000, -1000, -1000, -1000,
	-1000, 40, -15, 507, -1000, 428, -1000, 32, -18, -1000,
	399, -1000, -1000, 370, -1000, 39, -1000, 17, 6, 2,
	-1000, 253, -15, 223, -1000, 341, -1000, -1000, 36, -1000,
	507, 507, -1000, -1000, 0, -1000, -1000, 253, -1000, -1000,
	223, -1000, -1000, 12, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, 253, 457, -1000, 17, -1000, -1000,
	174, 312, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 20,
	0, 0, 223, -1000,
}
var yyPgo = [...]int{

	0, 177, 175, 20, 0, 173, 172, 171, 170, 17,
	25, 11, 169, 167, 165, 14, 162, 159, 157, 155,
	18, 154, 153, 152, 8, 6, 150, 149, 146, 1,
	142, 3, 140, 139, 138, 137, 133, 2, 22, 15,
	132, 129, 128, 125, 121, 118, 116, 115, 114, 112,
	111, 110, 7, 13, 109, 107, 5, 4, 105, 104,
	101, 100, 10, 97, 92,
}
var yyR1 = [...]int{

	0, 1, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	4, 4, 5, 5, 2, 2, 6, 6, 8, 8,
	7, 7, 7, 7, 7, 9, 9, 9, 21, 22,
	22, 23, 29, 29, 30, 10, 11, 11, 12, 12,
	13, 13, 13, 14, 14, 16, 15, 15, 17, 17,
	18, 35, 36, 36, 19, 39, 31, 31, 31, 31,
	31, 31, 31, 32, 32, 28, 28, 28, 28, 28,
	28, 28, 28, 28, 28, 28, 28, 28, 28, 28,
	28, 34, 33, 33, 42, 41, 41, 40, 24, 24,
	24, 25, 26, 27, 27, 37, 37, 38, 20, 43,
	46, 45, 45, 44, 48, 48, 48, 48, 48, 48,
	48, 47, 49, 50, 50, 51, 51, 52, 53, 53,
	54, 54, 55, 56, 56, 57, 58, 59, 60, 60,
	61, 62, 63, 63, 64,
}
var yyR2 = [...]int{

	0, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 0, 1, 1, 1, 0, 2, 1, 1,
	1, 3, 4, 4, 5, 1, 1, 1, 3, 0,
	2, 4, 0, 1, 2, 3, 0, 1, 0, 2,
	1, 1, 1, 4, 6, 3, 0, 1, 0, 2,
	3, 3, 5, 3, 6, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 3, 0, 2, 3, 0, 2, 3, 1, 1,
	1, 1, 3, 2, 2, 0, 2, 2, 3, 2,
	5, 0, 2, 3, 0, 2, 2, 2, 2, 2,
	2, 3, 7, 0, 2, 1, 2, 5, 1, 2,
	0, 1, 3, 1, 2, 5, 6, 5, 1, 3,
	6, 2, 1, 2, 6,
}
var yyChk = [...]int{

	-1000, -1, -2, -6, -43, -46, 17, -8, -7, -19,
	-10, -9, 8, 31, 14, 12, 18, -48, -37, -5,
	-4, -3, 15, 4, 5, 6, 7, 8, 23, 9,
	10, 11, 12, 13, 14, 16, 17, 18, 19, 20,
	21, -3, -12, -47, -49, -58, -59, -61, -64, 16,
	20, 11, 21, 5, 10, -20, 31, 36, -10, -21,
	-38, 27, -37, 15, 32, -13, -14, -35, -36, -4,
	25, -4, -4, -4, -4, -4, -4, -45, -4, -10,
	-38, -10, -22, -20, -39, -25, -4, -15, 29, -16,
	27, -3, 15, -37, -37, -50, 9, -37, -37, -37,
	-37, 32, -44, -9, -15, -10, 28, -23, 22, -37,
	-37, -4, -17, -37, -39, -10, -37, -51, -25, 31,
	30, 31, 31, 29, 29, -10, -11, -10, -15, 28,
	-18, -4, -37, 31, -25, -53, -52, -4, -60, -25,
	-63, -62, -4, -56, -57, -4, -25, -24, -25, -26,
	-27, 33, -37, 29, -10, -53, -52, 32, -54, -55,
	27, 37, 32, -62, -37, -57, 32, 29, -29, -30,
	30, 35, 35, -24, -11, -31, 24, -32, 22, -34,
	-28, -42, 13, 19, 7, 33, 4, 5, 6, 8,
	23, 9, 10, 11, 12, 15, 14, 16, 17, 18,
	20, 21, 31, 32, 29, -56, -25, -24, -31, 34,
	-33, -41, -24, 28, -29, 34, -31, 32, -40, -4,
	-37, -37, 29, -31,
}
var yyDef = [...]int{

	26, -2, 1, 24, 25, 114, 105, 27, 28, 29,
	30, 22, 0, 48, 35, 36, 37, 109, 0, 105,
	23, 20, 21, 2, 3, 4, 5, 6, 7, 8,
	9, 10, 11, 12, 13, 14, 15, 16, 17, 18,
	19, 0, 0, 115, 116, 117, 118, 119, 120, 0,
	0, 0, 0, 0, 0, 106, 111, 0, 31, 105,
	0, 39, 0, 0, 45, 49, 50, 51, 52, 56,
	105, 105, 123, 105, 105, 105, 105, 0, 56, 32,
	0, 33, 0, -2, 105, 65, 101, 105, 0, 57,
	58, 105, 0, 0, 121, 105, 0, 0, 0, 0,
	0, 110, 112, 0, 108, 34, 38, 40, 0, 0,
	46, 56, 0, 61, 105, 63, 0, 124, 125, 0,
	0, 0, 0, 0, 0, 64, 53, 47, 105, 55,
	59, 0, 0, 0, 126, 0, 128, 130, 137, 138,
	0, 142, 105, 0, 133, 0, 113, 42, 98, 99,
	100, 0, 46, 0, 62, 0, 129, 136, 0, 131,
	0, 0, 140, 143, 141, 134, 144, 0, 41, 43,
	0, 103, 104, 0, 54, 60, 66, 67, 68, 69,
	70, 71, 72, 73, 74, 92, 75, 76, 77, 78,
	79, 80, 81, 82, 83, 84, 85, 86, 87, 88,
	89, 90, 95, 122, 0, 0, 139, 42, 44, 102,
	0, 0, 105, 132, 105, 91, 93, 94, 96, 0,
	127, 135, 0, 97,
}
var yyTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 35, 3, 3, 3, 3, 3, 3,
	27, 28, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 29, 3,
	3, 30, 3, 3, 36, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 33, 3, 34, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 31, 37, 32,
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
		//line graphql_parser.y:148
		{
			// TODO check error
			yylex.(*lexer).doc = yyDollar[1].doc
			return 0
		}
	case 2:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line graphql_parser.y:156
		{
			yyVAL.str = yyDollar[1].str
		}
	case 3:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line graphql_parser.y:157
		{
			yyVAL.str = yyDollar[1].str
		}
	case 4:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line graphql_parser.y:158
		{
			yyVAL.str = yyDollar[1].str
		}
	case 5:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line graphql_parser.y:159
		{
			yyVAL.str = yyDollar[1].str
		}
	case 6:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line graphql_parser.y:160
		{
			yyVAL.str = yyDollar[1].str
		}
	case 7:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line graphql_parser.y:161
		{
			yyVAL.str = yyDollar[1].str
		}
	case 8:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line graphql_parser.y:162
		{
			yyVAL.str = yyDollar[1].str
		}
	case 9:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line graphql_parser.y:163
		{
			yyVAL.str = yyDollar[1].str
		}
	case 10:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line graphql_parser.y:164
		{
			yyVAL.str = yyDollar[1].str
		}
	case 11:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line graphql_parser.y:165
		{
			yyVAL.str = yyDollar[1].str
		}
	case 12:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line graphql_parser.y:166
		{
			yyVAL.str = yyDollar[1].str
		}
	case 13:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line graphql_parser.y:167
		{
			yyVAL.str = yyDollar[1].str
		}
	case 14:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line graphql_parser.y:168
		{
			yyVAL.str = yyDollar[1].str
		}
	case 15:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line graphql_parser.y:169
		{
			yyVAL.str = yyDollar[1].str
		}
	case 16:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line graphql_parser.y:170
		{
			yyVAL.str = yyDollar[1].str
		}
	case 17:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line graphql_parser.y:171
		{
			yyVAL.str = yyDollar[1].str
		}
	case 18:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line graphql_parser.y:172
		{
			yyVAL.str = yyDollar[1].str
		}
	case 19:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line graphql_parser.y:173
		{
			yyVAL.str = yyDollar[1].str
		}
	case 20:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line graphql_parser.y:176
		{
			yyVAL.str = yyDollar[1].str
		}
	case 21:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line graphql_parser.y:177
		{
			yyVAL.str = yyDollar[1].str
		}
	case 22:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line graphql_parser.y:180
		{
			yyVAL.str = ""
		}
	case 24:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line graphql_parser.y:188
		{
			yyVAL.doc = Document{Definitions: yyDollar[1].definitions}
		}
	case 25:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line graphql_parser.y:192
		{
			yyVAL.doc = Document{Schema: yyDollar[1].schema}
		}
	case 26:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line graphql_parser.y:199
		{
			yyVAL.definitions = nil
		}
	case 27:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line graphql_parser.y:200
		{
			yyVAL.definitions = append(yyDollar[1].definitions, yyDollar[2].definition)
		}
	case 28:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line graphql_parser.y:204
		{
			yyVAL.definition = Definition{operation: yyDollar[1].operation}
		}
	case 29:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line graphql_parser.y:205
		{
			yyVAL.definition = Definition{fragment: yyDollar[1].fragment}
		}
	case 30:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line graphql_parser.y:211
		{
			yyVAL.operation = Operation{OpType: Query, SelectionSet: yyDollar[1].selectionSet}
		}
	case 31:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line graphql_parser.y:215
		{
			yyVAL.operation = Operation{OpType: yyDollar[1].opType, Name: yyDollar[2].str, SelectionSet: yyDollar[3].selectionSet}
		}
	case 32:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line graphql_parser.y:219
		{
			yyVAL.operation = Operation{OpType: yyDollar[1].opType, Name: yyDollar[2].str, Variables: yyDollar[3].variables, SelectionSet: yyDollar[4].selectionSet}
		}
	case 33:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line graphql_parser.y:223
		{
			yyVAL.operation = Operation{OpType: yyDollar[1].opType, Name: yyDollar[2].str, Directives: yyDollar[3].directives, SelectionSet: yyDollar[4].selectionSet}
		}
	case 34:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line graphql_parser.y:227
		{
			yyVAL.operation = Operation{OpType: yyDollar[1].opType, Name: yyDollar[2].str, Variables: yyDollar[3].variables, Directives: yyDollar[4].directives, SelectionSet: yyDollar[5].selectionSet}
		}
	case 35:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line graphql_parser.y:233
		{
			yyVAL.opType = Query
		}
	case 36:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line graphql_parser.y:234
		{
			yyVAL.opType = Mutation
		}
	case 37:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line graphql_parser.y:235
		{
			yyVAL.opType = Subscription
		}
	case 38:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line graphql_parser.y:239
		{
			yyVAL.variables = yyDollar[2].variables
		}
	case 39:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line graphql_parser.y:243
		{
			yyVAL.variables = nil
		}
	case 40:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line graphql_parser.y:244
		{
			yyVAL.variables = append(yyDollar[1].variables, yyDollar[2].variable)
		}
	case 41:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line graphql_parser.y:249
		{
			yyVAL.variable = Variable{Name: yyDollar[1].str, Type: yyDollar[3].typeDesc, Default: yyDollar[4].val}
		}
	case 42:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line graphql_parser.y:255
		{
			yyVAL.val = reflect.ValueOf(nil)
		}
	case 43:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line graphql_parser.y:256
		{
			yyVAL.val = yyDollar[1].val
		}
	case 44:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line graphql_parser.y:260
		{
			yyVAL.val = yyDollar[2].val
		}
	case 45:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line graphql_parser.y:263
		{
			yyVAL.selectionSet = yyDollar[2].selectionSet
		}
	case 46:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line graphql_parser.y:267
		{
			yyVAL.selectionSet = nil
		}
	case 48:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line graphql_parser.y:272
		{
			yyVAL.selectionSet = nil
		}
	case 49:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line graphql_parser.y:273
		{
			yyVAL.selectionSet = append(yyDollar[1].selectionSet, yyDollar[2].selection)
		}
	case 50:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line graphql_parser.y:277
		{
			yyVAL.selection = Selection{Field: yyDollar[1].field}
		}
	case 51:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line graphql_parser.y:278
		{
			yyVAL.selection = Selection{FragmentSpread: yyDollar[1].fragmentSpread}
		}
	case 52:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line graphql_parser.y:279
		{
			yyVAL.selection = Selection{InlineFragment: yyDollar[1].fragment}
		}
	case 53:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line graphql_parser.y:283
		{
			yyVAL.field = Field{
				Name:         yyDollar[1].str,
				Arguments:    yyDollar[2].arguments,
				Directives:   yyDollar[3].directives,
				SelectionSet: yyDollar[4].selectionSet,
			}
		}
	case 54:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line graphql_parser.y:291
		{
			yyVAL.field = Field{
				Name:         yyDollar[1].str,
				Arguments:    yyDollar[4].arguments,
				Directives:   yyDollar[5].directives,
				SelectionSet: yyDollar[6].selectionSet,
				Alias:        yyDollar[3].str,
			}
		}
	case 55:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line graphql_parser.y:303
		{
			yyVAL.arguments = yyDollar[2].arguments
		}
	case 56:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line graphql_parser.y:307
		{
			yyVAL.arguments = nil
		}
	case 57:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line graphql_parser.y:308
		{
			yyVAL.arguments = yyDollar[1].arguments
		}
	case 58:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line graphql_parser.y:312
		{
			yyVAL.arguments = nil
		}
	case 59:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line graphql_parser.y:313
		{
			yyVAL.arguments = append(yyDollar[1].arguments, yyDollar[2].argument)
		}
	case 60:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line graphql_parser.y:317
		{
			yyVAL.argument = Argument{Name: yyDollar[1].str, Value: yyDollar[3].val}
		}
	case 61:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line graphql_parser.y:322
		{
			yyVAL.fragmentSpread = FragmentSpread{Name: yyDollar[2].str, Directives: yyDollar[3].directives}
		}
	case 62:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line graphql_parser.y:327
		{
			yyVAL.fragment = Fragment{
				TypeCondition: yyDollar[3].str,
				Directives:    yyDollar[4].directives,
				SelectionSet:  yyDollar[5].selectionSet,
			}
		}
	case 63:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line graphql_parser.y:335
		{
			yyVAL.fragment = Fragment{
				Directives:   yyDollar[2].directives,
				SelectionSet: yyDollar[3].selectionSet,
			}
		}
	case 64:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line graphql_parser.y:345
		{
			yyVAL.fragment = Fragment{
				Name:          yyDollar[2].str,
				TypeCondition: yyDollar[4].str,
				Directives:    yyDollar[5].directives,
				SelectionSet:  yyDollar[6].selectionSet,
			}
		}
	case 65:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line graphql_parser.y:355
		{
			yyVAL.str = yyDollar[1].str
		}
	case 66:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line graphql_parser.y:360
		{
			yyVAL.val = yyDollar[1].val
		}
	case 67:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line graphql_parser.y:361
		{
			yyVAL.val = yyDollar[1].val
		}
	case 68:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line graphql_parser.y:362
		{
			yyVAL.val = reflect.ValueOf(yyDollar[1].str)
		}
	case 69:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line graphql_parser.y:363
		{
			yyVAL.val = reflect.ValueOf(yyDollar[1].vals)
		}
	case 70:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line graphql_parser.y:364
		{
			yyVAL.val = reflect.ValueOf(yyDollar[1].str)
		}
	case 71:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line graphql_parser.y:365
		{
			yyVAL.val = reflect.ValueOf(yyDollar[1].objectFields)
		}
	case 72:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line graphql_parser.y:366
		{
			yyVAL.val = reflect.ValueOf(nil)
		}
	case 73:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line graphql_parser.y:369
		{
			yyVAL.val = reflect.ValueOf(true)
		}
	case 74:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line graphql_parser.y:370
		{
			yyVAL.val = reflect.ValueOf(false)
		}
	case 75:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line graphql_parser.y:373
		{
			yyVAL.str = yyDollar[1].str
		}
	case 76:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line graphql_parser.y:374
		{
			yyVAL.str = yyDollar[1].str
		}
	case 77:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line graphql_parser.y:375
		{
			yyVAL.str = yyDollar[1].str
		}
	case 78:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line graphql_parser.y:376
		{
			yyVAL.str = yyDollar[1].str
		}
	case 79:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line graphql_parser.y:377
		{
			yyVAL.str = yyDollar[1].str
		}
	case 80:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line graphql_parser.y:378
		{
			yyVAL.str = yyDollar[1].str
		}
	case 81:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line graphql_parser.y:379
		{
			yyVAL.str = yyDollar[1].str
		}
	case 82:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line graphql_parser.y:380
		{
			yyVAL.str = yyDollar[1].str
		}
	case 83:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line graphql_parser.y:381
		{
			yyVAL.str = yyDollar[1].str
		}
	case 84:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line graphql_parser.y:382
		{
			yyVAL.str = yyDollar[1].str
		}
	case 85:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line graphql_parser.y:383
		{
			yyVAL.str = yyDollar[1].str
		}
	case 86:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line graphql_parser.y:384
		{
			yyVAL.str = yyDollar[1].str
		}
	case 87:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line graphql_parser.y:385
		{
			yyVAL.str = yyDollar[1].str
		}
	case 88:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line graphql_parser.y:386
		{
			yyVAL.str = yyDollar[1].str
		}
	case 89:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line graphql_parser.y:387
		{
			yyVAL.str = yyDollar[1].str
		}
	case 90:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line graphql_parser.y:388
		{
			yyVAL.str = yyDollar[1].str
		}
	case 91:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line graphql_parser.y:393
		{
			yyVAL.vals = yyDollar[2].vals
		}
	case 92:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line graphql_parser.y:396
		{
			yyVAL.vals = nil
		}
	case 93:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line graphql_parser.y:397
		{
			yyVAL.vals = append(yyDollar[1].vals, yyDollar[2].val)
		}
	case 94:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line graphql_parser.y:401
		{
			yyVAL.objectFields = yyDollar[2].objectFields
		}
	case 95:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line graphql_parser.y:404
		{
			yyVAL.objectFields = nil
		}
	case 96:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line graphql_parser.y:405
		{
			yyVAL.objectFields = append(yyDollar[1].objectFields, yyDollar[2].objectField)
		}
	case 97:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line graphql_parser.y:408
		{
			yyVAL.objectField = ObjectField{Key: yyDollar[1].str, Value: yyDollar[3].val}
		}
	case 98:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line graphql_parser.y:413
		{
			yyVAL.typeDesc = TypeDescription{Name: yyDollar[1].str}
		}
	case 99:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line graphql_parser.y:414
		{
			yyVAL.typeDesc = yyDollar[1].typeDesc
		}
	case 100:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line graphql_parser.y:415
		{
			yyVAL.typeDesc = yyDollar[1].typeDesc
		}
	case 101:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line graphql_parser.y:418
		{
			yyVAL.str = yyDollar[1].str
		}
	case 102:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line graphql_parser.y:422
		{
			yyDollar[2].typeDesc.Flags |= List
			yyVAL.typeDesc = yyDollar[2].typeDesc
		}
	case 103:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line graphql_parser.y:429
		{
			yyVAL.typeDesc = TypeDescription{Name: yyDollar[1].str, Flags: Required}
		}
	case 104:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line graphql_parser.y:433
		{
			yyDollar[1].typeDesc.Flags |= Required
			yyVAL.typeDesc = yyDollar[1].typeDesc
		}
	case 105:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line graphql_parser.y:441
		{
			yyVAL.directives = nil
		}
	case 106:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line graphql_parser.y:442
		{
			yyVAL.directives = append(yyDollar[1].directives, yyDollar[2].directive)
		}
	case 107:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line graphql_parser.y:445
		{
			yyVAL.directives = append(yyDollar[1].directives, yyDollar[2].directive)
		}
	case 108:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line graphql_parser.y:448
		{
			yyVAL.directive = Directive{Name: yyDollar[2].str, Arguments: yyDollar[3].arguments}
		}
	case 109:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line graphql_parser.y:453
		{
			yyVAL.schema = Schema{OperationTypeDefinitions: yyDollar[1].operationTypeDefinitions, TypeDefinitions: yyDollar[2].typeDefinitions}
		}
	case 110:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line graphql_parser.y:457
		{
			yyVAL.operationTypeDefinitions = yyDollar[4].operationTypeDefinitions
		}
	case 111:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line graphql_parser.y:462
		{
			yyVAL.operationTypeDefinitions = nil
		}
	case 112:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line graphql_parser.y:463
		{
			yyVAL.operationTypeDefinitions = append(yyDollar[1].operationTypeDefinitions, yyDollar[2].operationTypeDefinition)
		}
	case 113:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line graphql_parser.y:467
		{
			yyVAL.operationTypeDefinition = OperationTypeDefinition{OpType: yyDollar[1].opType, Name: yyDollar[3].str}
		}
	case 114:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line graphql_parser.y:472
		{
			yyVAL.typeDefinitions = nil
		}
	case 115:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line graphql_parser.y:473
		{
			yyVAL.typeDefinitions = append(yyDollar[1].typeDefinitions, yyDollar[2].typeDefinition)
		}
	case 116:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line graphql_parser.y:474
		{
			yyVAL.typeDefinitions = append(yyDollar[1].typeDefinitions, yyDollar[2].typeDefinition)
		}
	case 117:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line graphql_parser.y:475
		{
			yyVAL.typeDefinitions = append(yyDollar[1].typeDefinitions, yyDollar[2].typeDefinition)
		}
	case 118:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line graphql_parser.y:476
		{
			yyVAL.typeDefinitions = append(yyDollar[1].typeDefinitions, yyDollar[2].typeDefinition)
		}
	case 119:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line graphql_parser.y:477
		{
			yyVAL.typeDefinitions = append(yyDollar[1].typeDefinitions, yyDollar[2].typeDefinition)
		}
	case 120:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line graphql_parser.y:478
		{
			yyVAL.typeDefinitions = append(yyDollar[1].typeDefinitions, yyDollar[2].typeDefinition)
		}
	case 121:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line graphql_parser.y:481
		{
			yyVAL.typeDefinition = ScalarDefinition{Name: yyDollar[2].str, Directives: yyDollar[3].directives}
		}
	case 122:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line graphql_parser.y:485
		{
			yyVAL.typeDefinition = ObjectDefinition{
				Name:       yyDollar[2].str,
				Implements: yyDollar[3].strs,
				Fields:     yyDollar[6].fieldsDef,
				Directives: yyDollar[4].directives,
			}
		}
	case 123:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line graphql_parser.y:495
		{
			yyVAL.strs = nil
		}
	case 124:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line graphql_parser.y:496
		{
			yyVAL.strs = yyDollar[2].strs
		}
	case 125:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line graphql_parser.y:499
		{
			yyVAL.strs = []string{yyDollar[1].str}
		}
	case 126:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line graphql_parser.y:500
		{
			yyVAL.strs = append(yyDollar[1].strs, yyDollar[2].str)
		}
	case 127:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line graphql_parser.y:505
		{
			yyVAL.fieldDef = FieldDefinition{Name: yyDollar[1].str, Arguments: yyDollar[2].inputDefinitions, Type: yyDollar[4].typeDesc, Directives: yyDollar[5].directives}
		}
	case 128:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line graphql_parser.y:510
		{
			yyVAL.fieldsDef = []FieldDefinition{yyDollar[1].fieldDef}
		}
	case 129:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line graphql_parser.y:511
		{
			yyVAL.fieldsDef = append(yyDollar[1].fieldsDef, yyDollar[2].fieldDef)
		}
	case 130:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line graphql_parser.y:514
		{
			yyVAL.inputDefinitions = nil
		}
	case 131:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line graphql_parser.y:515
		{
			yyVAL.inputDefinitions = yyDollar[1].inputDefinitions
		}
	case 132:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line graphql_parser.y:518
		{
			yyVAL.inputDefinitions = yyDollar[2].inputDefinitions
		}
	case 133:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line graphql_parser.y:521
		{
			yyVAL.inputDefinitions = []InputValueDefinition{yyDollar[1].inputDefinition}
		}
	case 134:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line graphql_parser.y:522
		{
			yyVAL.inputDefinitions = append(yyDollar[1].inputDefinitions, yyDollar[2].inputDefinition)
		}
	case 135:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line graphql_parser.y:526
		{
			yyVAL.inputDefinition = InputValueDefinition{Name: yyDollar[1].str, Type: yyDollar[3].typeDesc, Default: yyDollar[4].val, Directives: yyDollar[5].directives}
		}
	case 136:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line graphql_parser.y:532
		{
			yyVAL.typeDefinition = InterfaceDefinition{Name: yyDollar[2].str, Fields: yyDollar[5].fieldsDef, Directives: yyDollar[3].directives}
		}
	case 137:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line graphql_parser.y:538
		{
			yyVAL.typeDefinition = UnionDefinition{Name: yyDollar[2].str, Types: yyDollar[5].strs, Directives: yyDollar[3].directives}
		}
	case 138:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line graphql_parser.y:543
		{
			yyVAL.strs = []string{yyDollar[1].str}
		}
	case 139:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line graphql_parser.y:544
		{
			yyVAL.strs = append(yyDollar[1].strs, yyDollar[3].str)
		}
	case 140:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line graphql_parser.y:548
		{
			yyVAL.typeDefinition = EnumDefinition{Name: yyDollar[2].str, Values: yyDollar[5].strs, Directives: yyDollar[3].directives}
		}
	case 141:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line graphql_parser.y:553
		{
			yyVAL.str = yyDollar[1].str
		}
	case 142:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line graphql_parser.y:556
		{
			yyVAL.strs = []string{yyDollar[1].str}
		}
	case 143:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line graphql_parser.y:557
		{
			yyVAL.strs = append(yyDollar[1].strs, yyDollar[2].str)
		}
	case 144:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line graphql_parser.y:561
		{
			yyVAL.typeDefinition = InputObjectDefinition{
				Name:       yyDollar[2].str,
				InputDefs:  yyDollar[5].inputDefinitions,
				Directives: yyDollar[3].directives,
			}
		}
	}
	goto yystack /* stack new state and value */
}
