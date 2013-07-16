/*

Copyright (c) 2013 Go Authors. All rights reserved.
Use of this source code is governed by a BSD-style
license that can be found in the LICENSE file.

CAUTION: If this file is a Go source file (*.go), it was generated
automatically by '$ golex' from a *.l file - DO NOT EDIT in that case!

*/

%{

package main

import (
	"flag"
	"fmt"
	"log"
	"strconv"
	"unicode"

	"github.com/cznic/mathutil"
)

type lexer struct {
	c     int
	col   int
	errs  []error
	i     int
	lcol  int
	line  int
	ncol  int
	nline int
	sc    int
	src   string
	val   []byte
}

func newLexer(src string) (l *lexer) {
	l = &lexer{
		src:   src,
		nline: 1,
		ncol:  0,
	}
	l.next()
	return
}

func (l *lexer) next() int {
	if l.c != 0 {
		l.val = append(l.val, byte(l.c))
	}
	l.c = 0
	if l.i < len(l.src) {
		l.c = int(l.src[l.i])
		l.i++
	}
	switch l.c {
	case '\n':
		l.lcol = l.ncol
		l.nline++
		l.ncol = 0
	default:
		l.ncol++
	}
	return l.c
}

func (l *lexer) err(s string, arg ...interface{}) {
	err := fmt.Errorf(fmt.Sprintf("%d:%d ", l.line, l.col)+s, arg...)
	l.errs = append(l.errs, err)
}

func (l *lexer) Error(s string) {
	l.err(s)
}

func (l *lexer) Lex(lval *yySymType) int {
	const (
		INITIAL = iota
		S1
		S2
	)

	c0, c := 0, l.c
%}

int_lit         {decimal_lit}|{octal_lit}|{hex_lit}
decimal_lit     [1-9][0-9]*
octal_lit       0[0-7]*
hex_lit         0[xX][0-9a-fA-F]+

float_lit       {D}"."{D}?{E}?|{D}{E}|"."{D}{E}?
D        	[0-9]+
E        	[eE][-+]?[0-9]+

imaginary_ilit  {D}i
imaginary_lit   {float_lit}i

idchar0		[a-zA-Z_]
idchars		{idchar0}|[0-9]
ident		{idchar0}{idchars}*

%yyc c
%yyn c = l.next()
%yyt l.sc

%x S1 S2

%%
			l.val = l.val[:0]
			c0, l.line, l.col = l.c, l.nline, l.ncol
			
<*>\0			return 0

[ \t\n\r]+

{imaginary_ilit}	return l.int(lval, true)
{imaginary_lit}		return l.float(lval, true)
{int_lit}		return l.int(lval, false)
{float_lit}		return l.float(lval, false)

\"			l.sc = S1
`			l.sc = S2

'(\\.|[^'])*'		if ret := l.str(lval, ""); ret != STR {
				return ret
			}
			lval.item = int32(lval.item.(string)[0])
			return INTEGER

<S1>(\\.|[^\"])*\"	return l.str(lval, "\"")
<S2>([^`]|\n)*`		return l.str(lval, "`")

"&^"			return ANDNOT
"<<"			return LSH
">>"			return RSH

false			lval.item = false
			return BOOLEAN

true			lval.item = true
			return BOOLEAN

{ident}			lval.item = string(l.val)
			return IDENTIFIER

.			return c0

%%
			return int(unicode.ReplacementChar)
}

func (l *lexer) str(lval *yySymType, pref string) int {
	l.sc = 0
	s := pref + string(l.val)
	s, err := strconv.Unquote(s)
	if err != nil {
		l.err("string literal: %v", err)
		return int(unicode.ReplacementChar)
	}

	lval.item = s
	return STR
}

func (l *lexer) int(lval *yySymType, im bool) int {
	if im {
		l.val = l.val[:len(l.val)-1]
	}
	n, err := strconv.ParseUint(string(l.val), 0, 64)
	if err != nil {
		l.err("integer literal: %v", err)
		return int(unicode.ReplacementChar)
	}

	if im {
		lval.item = complex(0, float64(n))
		return IMAGINARY
	}

	switch {
	case n < mathutil.MaxInt:
		lval.item = int(n)
	default:
		lval.item = n
	}
	return INTEGER
}

func (l *lexer) float(lval *yySymType, im bool) int {
	if im {
		l.val = l.val[:len(l.val)-1]
	}
	n, err := strconv.ParseFloat(string(l.val), 64)
	if err != nil {
		l.err("float literal: %v", err)
		return int(unicode.ReplacementChar)
	}

	if im {
		lval.item = complex(0, n)
		return IMAGINARY
	}

	lval.item = n
	return FLOAT
}

func main() {
	oExpr := flag.String("e", "fmt.Printf(\"%d\\012\", -1 + 2.3*^3i | 4e2)", "The expression to parse")
	flag.Parse()
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	l := newLexer(*oExpr)
	if yyParse(l) != 0 {
		log.Fatal(l.errs)
	}

	fmt.Printf("AST for '%s'.\n", *oExpr)
	_dump()
}
