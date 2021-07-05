// Code generated by goyacc -o parse_ql.y.go parse_ql.y. DO NOT EDIT.

//line parse_ql.y:1
package parse

import (
	__yyfmt__ "fmt"

	"github.com/sophon-lab/temsearch/pkg/temql/labels"
)

//line parse_ql.y:3
//line parse_ql.y:7
type yySymType struct {
	yys      int
	item     Item
	node     Node
	matchers []*labels.Matcher
	matcher  *labels.Matcher
}

const LEFT_PAREN = 57346
const RIGHT_PAREN = 57347
const LEFT_BRACE = 57348
const RIGHT_BRACE = 57349
const ASSIGN = 57350
const LAND = 57351
const LOR = 57352
const EQL = 57353
const IDENTIFIER = 57354
const STRING = 57355
const METRIC_IDENTIFIER = 57356
const COMMA = 57357

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"LEFT_PAREN",
	"RIGHT_PAREN",
	"LEFT_BRACE",
	"RIGHT_BRACE",
	"ASSIGN",
	"LAND",
	"LOR",
	"EQL",
	"IDENTIFIER",
	"STRING",
	"METRIC_IDENTIFIER",
	"COMMA",
}

var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

//line parse_ql.y:94

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyPrivate = 57344

const yyLast = 20

var yyAct = [...]int{
	14, 20, 5, 9, 11, 12, 16, 7, 1, 8,
	3, 6, 19, 13, 10, 17, 2, 18, 4, 15,
}

var yyPact = [...]int{
	-12, -1000, -1000, -1000, 1, -1000, -1000, 2, -2, -1000,
	4, -1000, -1000, 2, -1000, -1, -1000, -1000, -1000, -1000,
	-1000,
}

var yyPgo = [...]int{
	0, 19, 18, 16, 11, 10, 9, 3, 8,
}

var yyR1 = [...]int{
	0, 8, 3, 1, 2, 4, 6, 6, 6, 7,
	7, 7, 7, 5,
}

var yyR2 = [...]int{
	0, 1, 1, 1, 1, 3, 3, 1, 2, 3,
	3, 2, 1, 2,
}

var yyChk = [...]int{
	-1000, -8, -3, -5, -2, 14, -4, 6, -6, -7,
	12, 2, 7, 15, 2, -1, 2, 11, -7, 13,
	2,
}

var yyDef = [...]int{
	0, -2, 1, 2, 0, 4, 13, 0, 0, 7,
	0, 12, 5, 0, 8, 0, 11, 3, 6, 9,
	10,
}

var yyTok1 = [...]int{
	1,
}

var yyTok2 = [...]int{
	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15,
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
//line parse_ql.y:42
		{
			yylex.(*parser).generatedParserResult = yyDollar[1].node
		}
	case 5:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parse_ql.y:55
		{
			yyVAL.node = &VectorSelector{
				LabelMatchers: yyDollar[2].matchers,
			}
		}
	case 6:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parse_ql.y:63
		{
			if yyDollar[1].matchers != nil {
				yyVAL.matchers = append(yyDollar[1].matchers, yyDollar[3].matcher)
			} else {
				yyVAL.matchers = yyDollar[1].matchers
			}
		}
	case 7:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parse_ql.y:71
		{
			yyVAL.matchers = []*labels.Matcher{yyDollar[1].matcher}
		}
	case 8:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parse_ql.y:73
		{
			yylex.(*parser).unexpected("label matching", "\",\" or \"}\"")
			yyVAL.matchers = yyDollar[1].matchers
		}
	case 9:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parse_ql.y:77
		{
			yyVAL.matcher = yylex.(*parser).newLabelMatcher(yyDollar[1].item, yyDollar[2].item, yyDollar[3].item)
		}
	case 10:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parse_ql.y:79
		{
			yylex.(*parser).unexpected("label matching", "string")
			yyVAL.matcher = nil
		}
	case 11:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parse_ql.y:81
		{
			yylex.(*parser).unexpected("label matching", "label matching operator")
			yyVAL.matcher = nil
		}
	case 12:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parse_ql.y:83
		{
			yylex.(*parser).unexpected("label matching", "identifier or \"}\"")
			yyVAL.matcher = nil
		}
	case 13:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parse_ql.y:87
		{
			vs := yyDollar[2].node.(*VectorSelector)
			vs.Name = yyDollar[1].item.Val
			yyVAL.node = vs
		}
	}
	goto yystack /* stack new state and value */
}