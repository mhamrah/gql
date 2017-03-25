//line parser.y:2
package parser

import __yyfmt__ "fmt"

//line parser.y:2
import (
	"github.com/mhamrah/gql"
	"reflect"
)

//line parser.y:10
type yySymType struct {
	yys  int
	str  string
	strs []string
	val  reflect.Value
	vals []reflect.Value
	doc  gql.Document

	definitions []gql.Definition
	definition  gql.Definition

	directives []gql.Directive
	directive  gql.Directive

	arguments map[string]gql.Argument
	argument  gql.Argument

	variables []gql.Variable
	variable  gql.Variable

	selectionSet []gql.Selection
	selection    gql.Selection

	operation      gql.Operation
	opType         gql.OpType
	field          gql.Field
	fragment       gql.Fragment
	fragmentSpread gql.FragmentSpread
	objectField    gql.ObjectField
	objectFields   []gql.ObjectField

	schema                   gql.Schema
	operationTypeDefinition  gql.OperationTypeDefinition
	operationTypeDefinitions []gql.OperationTypeDefinition
	typeDefinition           gql.TypeDefinition
	typeDefinitions          map[string]gql.TypeDefinition
	inputDefinitions         map[string]gql.InputValueDefinition
	inputDefinition          gql.InputValueDefinition
	fieldDef                 gql.FieldDefinition
	fieldsDef                []gql.FieldDefinition
	typeDesc                 gql.TypeDescription
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
	-1, 91,
	36, 105,
	-2, 106,
}

const yyNprod = 145
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 572

var yyAct = [...]int{

	94, 174, 135, 181, 127, 161, 140, 69, 134, 162,
	141, 132, 12, 33, 11, 126, 92, 67, 31, 62,
	14, 153, 64, 34, 95, 64, 56, 57, 58, 59,
	60, 61, 178, 124, 108, 107, 106, 54, 64, 64,
	64, 105, 64, 65, 63, 213, 64, 177, 13, 64,
	14, 176, 16, 16, 15, 15, 76, 10, 17, 17,
	98, 223, 96, 170, 78, 86, 81, 82, 83, 84,
	68, 14, 152, 109, 14, 167, 159, 87, 89, 138,
	93, 137, 116, 98, 88, 101, 70, 102, 114, 91,
	104, 80, 24, 131, 23, 7, 129, 119, 111, 113,
	117, 99, 18, 118, 22, 21, 128, 121, 133, 136,
	93, 112, 123, 125, 151, 150, 130, 122, 29, 103,
	79, 145, 20, 30, 27, 128, 19, 128, 139, 25,
	146, 148, 133, 26, 28, 136, 6, 157, 85, 110,
	147, 156, 4, 155, 142, 187, 215, 160, 128, 221,
	166, 75, 148, 136, 74, 185, 214, 168, 183, 175,
	186, 171, 164, 172, 163, 173, 115, 90, 66, 9,
	144, 179, 136, 180, 157, 211, 209, 120, 97, 73,
	212, 192, 193, 194, 190, 195, 197, 198, 199, 200,
	188, 202, 201, 203, 204, 205, 189, 206, 207, 184,
	196, 182, 72, 55, 5, 8, 3, 32, 208, 2,
	191, 218, 1, 0, 0, 0, 222, 216, 219, 217,
	0, 0, 0, 0, 0, 0, 0, 224, 192, 193,
	194, 190, 195, 197, 198, 199, 200, 188, 202, 201,
	203, 204, 205, 189, 206, 207, 184, 196, 182, 0,
	0, 0, 0, 0, 0, 208, 0, 191, 36, 37,
	38, 39, 40, 42, 43, 44, 45, 46, 47, 35,
	48, 49, 50, 51, 52, 53, 0, 41, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 165, 36, 37,
	38, 39, 40, 42, 43, 44, 45, 46, 47, 35,
	48, 49, 50, 51, 52, 53, 0, 41, 0, 77,
	0, 0, 0, 0, 0, 0, 71, 36, 37, 38,
	39, 40, 42, 43, 44, 45, 46, 47, 35, 48,
	49, 50, 51, 52, 53, 0, 41, 0, 0, 0,
	0, 0, 0, 0, 0, 220, 36, 37, 38, 39,
	40, 42, 43, 44, 45, 46, 47, 35, 48, 49,
	50, 51, 52, 53, 0, 41, 0, 0, 0, 0,
	0, 0, 0, 0, 169, 36, 37, 38, 39, 40,
	42, 43, 44, 45, 46, 47, 35, 48, 49, 50,
	51, 52, 53, 0, 41, 0, 0, 0, 0, 0,
	0, 0, 0, 158, 36, 37, 38, 39, 40, 42,
	43, 44, 45, 46, 47, 35, 48, 49, 50, 51,
	52, 53, 0, 41, 0, 0, 0, 0, 0, 0,
	0, 0, 154, 36, 37, 38, 39, 40, 42, 43,
	44, 45, 46, 47, 35, 48, 49, 50, 51, 52,
	53, 0, 41, 0, 0, 0, 0, 0, 0, 0,
	0, 149, 36, 37, 38, 39, 40, 42, 43, 44,
	45, 46, 47, 35, 48, 49, 50, 51, 52, 53,
	0, 41, 0, 0, 0, 0, 210, 36, 37, 38,
	39, 40, 42, 43, 44, 45, 46, 47, 35, 48,
	49, 50, 51, 52, 53, 0, 41, 0, 0, 0,
	0, 143, 36, 37, 38, 39, 40, 42, 43, 44,
	45, 46, 47, 35, 48, 49, 50, 51, 52, 53,
	0, 41, 36, 37, 38, 39, 40, 42, 43, 44,
	45, 46, 47, 100, 48, 49, 50, 51, 52, 53,
	0, 41, 36, 37, 38, 39, 40, 42, 43, 44,
	45, 46, 47, 0, 48, 49, 50, 51, 52, 53,
	0, 41,
}
var yyPact = [...]int{

	-1000, -1000, -1000, 40, -1000, -1000, -1000, 113, -1000, -1000,
	-1000, -1000, 508, 548, -1000, -1000, -1000, -1000, 113, -1000,
	-1000, -1000, -1000, -1000, -1000, 508, 508, 508, 508, 508,
	508, 13, 43, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, 71, 284, -1000, 82, -1000, -1000,
	-1000, -1000, -1000, -1000, 508, -1000, 19, 19, -1000, -14,
	508, -1000, -1000, -1000, -1000, -1000, 33, 528, -14, -1000,
	508, 10, 6, 4, 3, 41, 56, -1000, 19, -1000,
	60, -1000, -1000, -1000, -1000, -1000, 508, -1000, -1000, -1000,
	508, -11, 2, 508, -1000, 508, 508, 508, 508, -1000,
	-1000, 52, -1000, -1000, -1000, -1000, 50, -11, -11, 56,
	483, -14, -1000, -1000, 508, -1000, 429, -1000, 45, -16,
	-1000, 400, -1000, -1000, 371, -1000, 47, 508, 254, -1000,
	-1000, -1000, -1000, -1000, -1000, 46, -11, 342, -1000, -1000,
	34, -1000, 508, 508, -1000, -1000, -14, -1000, -1000, 254,
	-1000, 21, 12, -3, -1000, 254, -11, 224, -1000, -1000,
	254, 458, -1000, 21, -1000, -1000, 224, -1000, -1000, 11,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, 177, 313, -14, -14, -1000, -1000,
	-1000, -1000, 32, 224, -1000,
}
var yyPgo = [...]int{

	0, 212, 209, 23, 0, 207, 206, 205, 204, 12,
	10, 6, 203, 202, 179, 24, 178, 177, 170, 169,
	19, 168, 167, 166, 5, 9, 164, 162, 160, 1,
	159, 3, 158, 156, 155, 154, 151, 7, 17, 16,
	149, 146, 145, 142, 139, 138, 136, 126, 95, 122,
	120, 119, 4, 15, 115, 114, 8, 2, 105, 104,
	96, 94, 11, 93, 92,
}
var yyR1 = [...]int{

	0, 1, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	4, 4, 5, 5, 2, 6, 6, 8, 8, 7,
	7, 7, 7, 7, 9, 9, 9, 21, 22, 22,
	23, 29, 29, 30, 10, 11, 11, 12, 12, 13,
	13, 13, 14, 14, 16, 15, 15, 17, 17, 18,
	35, 36, 36, 19, 39, 31, 31, 31, 31, 31,
	31, 31, 32, 32, 28, 28, 28, 28, 28, 28,
	28, 28, 28, 28, 28, 28, 28, 28, 28, 28,
	34, 33, 33, 42, 41, 41, 40, 24, 24, 24,
	25, 26, 27, 27, 37, 37, 38, 20, 43, 43,
	46, 45, 45, 44, 48, 48, 48, 48, 48, 48,
	48, 47, 49, 50, 50, 51, 51, 52, 53, 53,
	54, 54, 55, 56, 56, 57, 58, 59, 60, 60,
	61, 62, 63, 63, 64,
}
var yyR2 = [...]int{

	0, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 0, 1, 2, 0, 2, 1, 1, 1,
	3, 4, 4, 5, 1, 1, 1, 3, 0, 2,
	4, 0, 1, 2, 3, 0, 1, 0, 2, 1,
	1, 1, 4, 6, 3, 0, 1, 0, 2, 3,
	3, 5, 3, 6, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	3, 0, 2, 3, 0, 2, 3, 1, 1, 1,
	1, 3, 2, 2, 0, 2, 2, 3, 2, 1,
	5, 0, 2, 3, 0, 2, 2, 2, 2, 2,
	2, 3, 7, 0, 2, 1, 2, 5, 1, 2,
	0, 1, 3, 1, 2, 5, 6, 5, 1, 3,
	6, 2, 1, 2, 6,
}
var yyChk = [...]int{

	-1000, -1, -2, -6, -43, -8, -46, -48, -7, -19,
	17, -10, -9, 8, 31, 14, 12, 18, -48, -47,
	-49, -58, -59, -61, -64, 16, 20, 11, 21, 5,
	10, -37, -5, -4, -3, 15, 4, 5, 6, 7,
	8, 23, 9, 10, 11, 12, 13, 14, 16, 17,
	18, 19, 20, 21, -3, -12, -4, -4, -4, -4,
	-4, -4, -20, 31, 36, -10, -21, -38, 27, -37,
	15, 32, -13, -14, -35, -36, -4, 25, -37, -50,
	9, -37, -37, -37, -37, -45, -4, -10, -38, -10,
	-22, -20, -39, -25, -4, -15, 29, -16, 27, -3,
	15, -37, -37, -51, -25, 31, 30, 31, 31, 32,
	-44, -9, -15, -10, 28, -23, 22, -37, -37, -4,
	-17, -37, -39, -10, 31, -25, -53, -52, -4, -60,
	-25, -63, -62, -4, -56, -57, -4, 29, 29, -10,
	-11, -10, -15, 28, -18, -4, -37, -53, -52, 32,
	-54, -55, 27, 37, 32, -62, -37, -57, 32, 29,
	-25, -24, -25, -26, -27, 33, -37, 29, -10, 32,
	29, -56, -25, -24, -29, -30, 30, 35, 35, -24,
	-11, -31, 24, -32, 22, -34, -28, -42, 13, 19,
	7, 33, 4, 5, 6, 8, 23, 9, 10, 11,
	12, 15, 14, 16, 17, 18, 20, 21, 31, -24,
	28, -29, -31, 34, -33, -41, -37, -37, 34, -31,
	32, -40, -4, 29, -31,
}
var yyDef = [...]int{

	25, -2, 1, 114, 24, 26, 114, 109, 27, 28,
	104, 29, 22, 0, 47, 34, 35, 36, 108, 115,
	116, 117, 118, 119, 120, 0, 0, 0, 0, 0,
	0, 0, 104, 23, 20, 21, 2, 3, 4, 5,
	6, 7, 8, 9, 10, 11, 12, 13, 14, 15,
	16, 17, 18, 19, 0, 0, 104, 123, 104, 104,
	104, 104, 105, 111, 0, 30, 104, 0, 38, 0,
	0, 44, 48, 49, 50, 51, 55, 104, 121, 104,
	0, 0, 0, 0, 0, 0, 55, 31, 0, 32,
	0, -2, 104, 64, 100, 104, 0, 56, 57, 104,
	0, 0, 0, 124, 125, 0, 0, 0, 0, 110,
	112, 0, 107, 33, 37, 39, 0, 0, 45, 55,
	0, 60, 104, 62, 0, 126, 0, 128, 130, 137,
	138, 0, 142, 104, 0, 133, 0, 0, 0, 63,
	52, 46, 104, 54, 58, 0, 0, 0, 129, 136,
	0, 131, 0, 0, 140, 143, 141, 134, 144, 0,
	113, 41, 97, 98, 99, 0, 45, 0, 61, 122,
	0, 0, 139, 41, 40, 42, 0, 102, 103, 0,
	53, 59, 65, 66, 67, 68, 69, 70, 71, 72,
	73, 91, 74, 75, 76, 77, 78, 79, 80, 81,
	82, 83, 84, 85, 86, 87, 88, 89, 94, 104,
	132, 104, 43, 101, 0, 0, 127, 135, 90, 92,
	93, 95, 0, 0, 96,
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
		//line parser.y:149
		{
			// TODO check error
			yylex.(*lexer).doc = yyDollar[1].doc
			return 0
		}
	case 2:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:157
		{
			yyVAL.str = yyDollar[1].str
		}
	case 3:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:158
		{
			yyVAL.str = yyDollar[1].str
		}
	case 4:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:159
		{
			yyVAL.str = yyDollar[1].str
		}
	case 5:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:160
		{
			yyVAL.str = yyDollar[1].str
		}
	case 6:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:161
		{
			yyVAL.str = yyDollar[1].str
		}
	case 7:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:162
		{
			yyVAL.str = yyDollar[1].str
		}
	case 8:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:163
		{
			yyVAL.str = yyDollar[1].str
		}
	case 9:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:164
		{
			yyVAL.str = yyDollar[1].str
		}
	case 10:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:165
		{
			yyVAL.str = yyDollar[1].str
		}
	case 11:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:166
		{
			yyVAL.str = yyDollar[1].str
		}
	case 12:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:167
		{
			yyVAL.str = yyDollar[1].str
		}
	case 13:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:168
		{
			yyVAL.str = yyDollar[1].str
		}
	case 14:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:169
		{
			yyVAL.str = yyDollar[1].str
		}
	case 15:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:170
		{
			yyVAL.str = yyDollar[1].str
		}
	case 16:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:171
		{
			yyVAL.str = yyDollar[1].str
		}
	case 17:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:172
		{
			yyVAL.str = yyDollar[1].str
		}
	case 18:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:173
		{
			yyVAL.str = yyDollar[1].str
		}
	case 19:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:174
		{
			yyVAL.str = yyDollar[1].str
		}
	case 20:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:177
		{
			yyVAL.str = yyDollar[1].str
		}
	case 21:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:178
		{
			yyVAL.str = yyDollar[1].str
		}
	case 22:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.y:181
		{
			yyVAL.str = ""
		}
	case 24:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:189
		{
			yyVAL.doc = gql.Document{Definitions: yyDollar[1].definitions, Schema: yyDollar[2].schema}
		}
	case 25:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.y:196
		{
			yyVAL.definitions = nil
		}
	case 26:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:197
		{
			yyVAL.definitions = append(yyDollar[1].definitions, yyDollar[2].definition)
		}
	case 27:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:201
		{
			yyVAL.definition = gql.Definition{Operation: yyDollar[1].operation}
		}
	case 28:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:202
		{
			yyVAL.definition = gql.Definition{Fragment: yyDollar[1].fragment}
		}
	case 29:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:208
		{
			yyVAL.operation = gql.Operation{OpType: gql.Query, SelectionSet: yyDollar[1].selectionSet}
		}
	case 30:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:212
		{
			yyVAL.operation = gql.Operation{OpType: yyDollar[1].opType, Name: yyDollar[2].str, SelectionSet: yyDollar[3].selectionSet}
		}
	case 31:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:216
		{
			yyVAL.operation = gql.Operation{OpType: yyDollar[1].opType, Name: yyDollar[2].str, Variables: yyDollar[3].variables, SelectionSet: yyDollar[4].selectionSet}
		}
	case 32:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:220
		{
			yyVAL.operation = gql.Operation{OpType: yyDollar[1].opType, Name: yyDollar[2].str, Directives: yyDollar[3].directives, SelectionSet: yyDollar[4].selectionSet}
		}
	case 33:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.y:224
		{
			yyVAL.operation = gql.Operation{OpType: yyDollar[1].opType, Name: yyDollar[2].str, Variables: yyDollar[3].variables, Directives: yyDollar[4].directives, SelectionSet: yyDollar[5].selectionSet}
		}
	case 34:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:230
		{
			yyVAL.opType = gql.Query
		}
	case 35:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:231
		{
			yyVAL.opType = gql.Mutation
		}
	case 36:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:232
		{
			yyVAL.opType = gql.Subscription
		}
	case 37:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:236
		{
			yyVAL.variables = yyDollar[2].variables
		}
	case 38:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.y:240
		{
			yyVAL.variables = nil
		}
	case 39:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:241
		{
			yyVAL.variables = append(yyDollar[1].variables, yyDollar[2].variable)
		}
	case 40:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:246
		{
			yyVAL.variable = gql.Variable{Name: yyDollar[1].str, Type: yyDollar[3].typeDesc, Default: yyDollar[4].val}
		}
	case 41:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.y:252
		{
			yyVAL.val = reflect.ValueOf(nil)
		}
	case 42:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:253
		{
			yyVAL.val = yyDollar[1].val
		}
	case 43:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:257
		{
			yyVAL.val = yyDollar[2].val
		}
	case 44:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:260
		{
			yyVAL.selectionSet = yyDollar[2].selectionSet
		}
	case 45:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.y:264
		{
			yyVAL.selectionSet = nil
		}
	case 47:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.y:269
		{
			yyVAL.selectionSet = nil
		}
	case 48:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:270
		{
			yyVAL.selectionSet = append(yyDollar[1].selectionSet, yyDollar[2].selection)
		}
	case 49:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:274
		{
			yyVAL.selection = gql.Selection{Field: yyDollar[1].field}
		}
	case 50:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:275
		{
			yyVAL.selection = gql.Selection{FragmentSpread: yyDollar[1].fragmentSpread}
		}
	case 51:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:276
		{
			yyVAL.selection = gql.Selection{InlineFragment: yyDollar[1].fragment}
		}
	case 52:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:280
		{
			yyVAL.field = gql.Field{
				Name:         yyDollar[1].str,
				Arguments:    yyDollar[2].arguments,
				Directives:   yyDollar[3].directives,
				SelectionSet: yyDollar[4].selectionSet,
			}
		}
	case 53:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.y:288
		{
			yyVAL.field = gql.Field{
				Name:         yyDollar[1].str,
				Arguments:    yyDollar[4].arguments,
				Directives:   yyDollar[5].directives,
				SelectionSet: yyDollar[6].selectionSet,
				Alias:        yyDollar[3].str,
			}
		}
	case 54:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:299
		{
			yyVAL.arguments = yyDollar[2].arguments
		}
	case 55:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.y:302
		{
			yyVAL.arguments = nil
		}
	case 56:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:303
		{
			yyVAL.arguments = yyDollar[1].arguments
		}
	case 57:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.y:306
		{
			yyVAL.arguments = make(map[string]gql.Argument)
		}
	case 58:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:307
		{
			yyDollar[1].arguments[yyDollar[2].argument.Name] = yyDollar[2].argument
			yyVAL.arguments = yyDollar[1].arguments
		}
	case 59:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:310
		{
			yyVAL.argument = gql.Argument{Name: yyDollar[1].str, Value: yyDollar[3].val}
		}
	case 60:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:315
		{
			yyVAL.fragmentSpread = gql.FragmentSpread{Name: yyDollar[2].str, Directives: yyDollar[3].directives}
		}
	case 61:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.y:320
		{
			yyVAL.fragment = gql.Fragment{
				TypeCondition: yyDollar[3].str,
				Directives:    yyDollar[4].directives,
				SelectionSet:  yyDollar[5].selectionSet,
			}
		}
	case 62:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:328
		{
			yyVAL.fragment = gql.Fragment{
				Directives:   yyDollar[2].directives,
				SelectionSet: yyDollar[3].selectionSet,
			}
		}
	case 63:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.y:338
		{
			yyVAL.fragment = gql.Fragment{
				Name:          yyDollar[2].str,
				TypeCondition: yyDollar[4].str,
				Directives:    yyDollar[5].directives,
				SelectionSet:  yyDollar[6].selectionSet,
			}
		}
	case 64:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:348
		{
			yyVAL.str = yyDollar[1].str
		}
	case 65:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:353
		{
			yyVAL.val = yyDollar[1].val
		}
	case 66:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:354
		{
			yyVAL.val = yyDollar[1].val
		}
	case 67:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:355
		{
			yyVAL.val = reflect.ValueOf(yyDollar[1].str)
		}
	case 68:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:356
		{
			yyVAL.val = reflect.ValueOf(yyDollar[1].vals)
		}
	case 69:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:357
		{
			yyVAL.val = reflect.ValueOf(yyDollar[1].str)
		}
	case 70:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:358
		{
			yyVAL.val = reflect.ValueOf(yyDollar[1].objectFields)
		}
	case 71:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:359
		{
			yyVAL.val = reflect.ValueOf(nil)
		}
	case 72:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:362
		{
			yyVAL.val = reflect.ValueOf(true)
		}
	case 73:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:363
		{
			yyVAL.val = reflect.ValueOf(false)
		}
	case 74:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:366
		{
			yyVAL.str = yyDollar[1].str
		}
	case 75:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:367
		{
			yyVAL.str = yyDollar[1].str
		}
	case 76:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:368
		{
			yyVAL.str = yyDollar[1].str
		}
	case 77:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:369
		{
			yyVAL.str = yyDollar[1].str
		}
	case 78:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:370
		{
			yyVAL.str = yyDollar[1].str
		}
	case 79:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:371
		{
			yyVAL.str = yyDollar[1].str
		}
	case 80:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:372
		{
			yyVAL.str = yyDollar[1].str
		}
	case 81:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:373
		{
			yyVAL.str = yyDollar[1].str
		}
	case 82:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:374
		{
			yyVAL.str = yyDollar[1].str
		}
	case 83:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:375
		{
			yyVAL.str = yyDollar[1].str
		}
	case 84:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:376
		{
			yyVAL.str = yyDollar[1].str
		}
	case 85:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:377
		{
			yyVAL.str = yyDollar[1].str
		}
	case 86:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:378
		{
			yyVAL.str = yyDollar[1].str
		}
	case 87:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:379
		{
			yyVAL.str = yyDollar[1].str
		}
	case 88:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:380
		{
			yyVAL.str = yyDollar[1].str
		}
	case 89:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:381
		{
			yyVAL.str = yyDollar[1].str
		}
	case 90:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:386
		{
			yyVAL.vals = yyDollar[2].vals
		}
	case 91:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.y:389
		{
			yyVAL.vals = nil
		}
	case 92:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:390
		{
			yyVAL.vals = append(yyDollar[1].vals, yyDollar[2].val)
		}
	case 93:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:394
		{
			yyVAL.objectFields = yyDollar[2].objectFields
		}
	case 94:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.y:397
		{
			yyVAL.objectFields = nil
		}
	case 95:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:398
		{
			yyVAL.objectFields = append(yyDollar[1].objectFields, yyDollar[2].objectField)
		}
	case 96:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:401
		{
			yyVAL.objectField = gql.ObjectField{Key: yyDollar[1].str, Value: yyDollar[3].val}
		}
	case 97:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:406
		{
			yyVAL.typeDesc = gql.TypeDescription{Name: yyDollar[1].str}
		}
	case 98:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:407
		{
			yyVAL.typeDesc = yyDollar[1].typeDesc
		}
	case 99:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:408
		{
			yyVAL.typeDesc = yyDollar[1].typeDesc
		}
	case 100:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:411
		{
			yyVAL.str = yyDollar[1].str
		}
	case 101:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:415
		{
			yyDollar[2].typeDesc.Flags |= gql.List
			yyVAL.typeDesc = yyDollar[2].typeDesc
		}
	case 102:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:422
		{
			yyVAL.typeDesc = gql.TypeDescription{Name: yyDollar[1].str, Flags: gql.Required}
		}
	case 103:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:426
		{
			yyDollar[1].typeDesc.Flags |= gql.Required
			yyVAL.typeDesc = yyDollar[1].typeDesc
		}
	case 104:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.y:434
		{
			yyVAL.directives = nil
		}
	case 105:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:435
		{
			yyVAL.directives = append(yyDollar[1].directives, yyDollar[2].directive)
		}
	case 106:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:438
		{
			yyVAL.directives = append(yyDollar[1].directives, yyDollar[2].directive)
		}
	case 107:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:441
		{
			yyVAL.directive = gql.Directive{Name: yyDollar[2].str, Arguments: yyDollar[3].arguments}
		}
	case 108:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:446
		{
			yyVAL.schema = gql.Schema{OperationTypeDefinitions: yyDollar[1].operationTypeDefinitions, TypeDefinitions: yyDollar[2].typeDefinitions}
		}
	case 109:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:447
		{
			yyVAL.schema = gql.Schema{TypeDefinitions: yyDollar[1].typeDefinitions}
		}
	case 110:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.y:452
		{
			yyVAL.operationTypeDefinitions = yyDollar[4].operationTypeDefinitions
		}
	case 111:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.y:457
		{
			yyVAL.operationTypeDefinitions = nil
		}
	case 112:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:458
		{
			yyVAL.operationTypeDefinitions = append(yyDollar[1].operationTypeDefinitions, yyDollar[2].operationTypeDefinition)
		}
	case 113:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:462
		{
			yyVAL.operationTypeDefinition = gql.OperationTypeDefinition{OpType: yyDollar[1].opType, Name: yyDollar[3].str}
		}
	case 114:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.y:467
		{
			yyVAL.typeDefinitions = gql.BuiltinDefinitions()
		}
	case 115:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:468
		{
			yyDollar[1].typeDefinitions[yyDollar[2].typeDefinition.NamedType()] = yyDollar[2].typeDefinition
			yyVAL.typeDefinitions = yyDollar[1].typeDefinitions
		}
	case 116:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:469
		{
			yyDollar[1].typeDefinitions[yyDollar[2].typeDefinition.NamedType()] = yyDollar[2].typeDefinition
			yyVAL.typeDefinitions = yyDollar[1].typeDefinitions
		}
	case 117:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:470
		{
			yyDollar[1].typeDefinitions[yyDollar[2].typeDefinition.NamedType()] = yyDollar[2].typeDefinition
			yyVAL.typeDefinitions = yyDollar[1].typeDefinitions
		}
	case 118:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:471
		{
			yyDollar[1].typeDefinitions[yyDollar[2].typeDefinition.NamedType()] = yyDollar[2].typeDefinition
			yyVAL.typeDefinitions = yyDollar[1].typeDefinitions
		}
	case 119:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:472
		{
			yyDollar[1].typeDefinitions[yyDollar[2].typeDefinition.NamedType()] = yyDollar[2].typeDefinition
			yyVAL.typeDefinitions = yyDollar[1].typeDefinitions
		}
	case 120:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:473
		{
			yyDollar[1].typeDefinitions[yyDollar[2].typeDefinition.NamedType()] = yyDollar[2].typeDefinition
			yyVAL.typeDefinitions = yyDollar[1].typeDefinitions
		}
	case 121:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:476
		{
			yyVAL.typeDefinition = gql.ScalarDefinition{Name: yyDollar[2].str, Directives: yyDollar[3].directives}
		}
	case 122:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line parser.y:480
		{
			yyVAL.typeDefinition = gql.ObjectDefinition{
				Name:       yyDollar[2].str,
				Implements: yyDollar[3].strs,
				Fields:     yyDollar[6].fieldsDef,
				Directives: yyDollar[4].directives,
			}
		}
	case 123:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.y:490
		{
			yyVAL.strs = nil
		}
	case 124:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:491
		{
			yyVAL.strs = yyDollar[2].strs
		}
	case 125:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:494
		{
			yyVAL.strs = []string{yyDollar[1].str}
		}
	case 126:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:495
		{
			yyVAL.strs = append(yyDollar[1].strs, yyDollar[2].str)
		}
	case 127:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.y:500
		{
			yyVAL.fieldDef = gql.FieldDefinition{Name: yyDollar[1].str, Arguments: yyDollar[2].inputDefinitions, Type: yyDollar[4].typeDesc, Directives: yyDollar[5].directives}
		}
	case 128:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:505
		{
			yyVAL.fieldsDef = []gql.FieldDefinition{yyDollar[1].fieldDef}
		}
	case 129:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:506
		{
			yyVAL.fieldsDef = append(yyDollar[1].fieldsDef, yyDollar[2].fieldDef)
		}
	case 130:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.y:509
		{
			yyVAL.inputDefinitions = nil
		}
	case 131:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:510
		{
			yyVAL.inputDefinitions = yyDollar[1].inputDefinitions
		}
	case 132:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:513
		{
			yyVAL.inputDefinitions = yyDollar[2].inputDefinitions
		}
	case 133:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:516
		{
			yyVAL.inputDefinitions = make(map[string]gql.InputValueDefinition)
			yyVAL.inputDefinitions[yyDollar[1].inputDefinition.Name] = yyDollar[1].inputDefinition
		}
	case 134:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:517
		{
			yyDollar[1].inputDefinitions[yyDollar[2].inputDefinition.Name] = yyDollar[2].inputDefinition
			yyVAL.inputDefinitions = yyDollar[1].inputDefinitions
		}
	case 135:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.y:521
		{
			yyVAL.inputDefinition = gql.InputValueDefinition{Name: yyDollar[1].str, Type: yyDollar[3].typeDesc, Default: yyDollar[4].val, Directives: yyDollar[5].directives}
		}
	case 136:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.y:527
		{
			yyVAL.typeDefinition = gql.InterfaceDefinition{Name: yyDollar[2].str, Fields: yyDollar[5].fieldsDef, Directives: yyDollar[3].directives}
		}
	case 137:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.y:533
		{
			yyVAL.typeDefinition = gql.UnionDefinition{Name: yyDollar[2].str, Types: yyDollar[5].strs, Directives: yyDollar[3].directives}
		}
	case 138:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:538
		{
			yyVAL.strs = []string{yyDollar[1].str}
		}
	case 139:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:539
		{
			yyVAL.strs = append(yyDollar[1].strs, yyDollar[3].str)
		}
	case 140:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.y:543
		{
			yyVAL.typeDefinition = gql.EnumDefinition{Name: yyDollar[2].str, Values: yyDollar[5].strs, Directives: yyDollar[3].directives}
		}
	case 141:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:548
		{
			yyVAL.str = yyDollar[1].str
		}
	case 142:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:551
		{
			yyVAL.strs = []string{yyDollar[1].str}
		}
	case 143:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:552
		{
			yyVAL.strs = append(yyDollar[1].strs, yyDollar[2].str)
		}
	case 144:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.y:556
		{
			yyVAL.typeDefinition = gql.InputObjectDefinition{
				Name:       yyDollar[2].str,
				InputDefs:  yyDollar[5].inputDefinitions,
				Directives: yyDollar[3].directives,
			}
		}
	}
	goto yystack /* stack new state and value */
}
