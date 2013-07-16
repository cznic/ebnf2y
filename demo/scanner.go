/*

Copyright (c) 2013 Go Authors. All rights reserved.
Use of this source code is governed by a BSD-style
license that can be found in the LICENSE file.

CAUTION: If this file is a Go source file (*.go), it was generated
automatically by '$ golex' from a *.l file - DO NOT EDIT in that case!

*/

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

yystate0:

	l.val = l.val[:0]
	c0, l.line, l.col = l.c, l.nline, l.ncol

	switch yyt := l.sc; yyt {
	default:
		panic(fmt.Errorf(`invalid start condition %d`, yyt))
	case 0: // start condition: INITIAL
		goto yystart1
	case 1: // start condition: S1
		goto yystart44
	case 2: // start condition: S2
		goto yystart49
	}

	goto yystate1 // silence unused label error
yystate1:
	c = l.next()
yystart1:
	switch {
	default:
		goto yystate3 // c >= '\x01' && c <= '\b' || c == '\v' || c == '\f' || c >= '\x0e' && c <= '\x1f' || c == '!' || c >= '#' && c <= '%%' || c >= '(' && c <= '-' || c == '/' || c == ':' || c == ';' || c == '=' || c == '?' || c == '@' || c >= '[' && c <= '^' || c >= '{' && c <= 'ÿ'
	case c == '"':
		goto yystate6
	case c == '&':
		goto yystate7
	case c == '.':
		goto yystate14
	case c == '0':
		goto yystate20
	case c == '<':
		goto yystate28
	case c == '>':
		goto yystate30
	case c == '\'':
		goto yystate9
	case c == '\n':
		goto yystate5
	case c == '\t' || c == '\r' || c == ' ':
		goto yystate4
	case c == '\x00':
		goto yystate2
	case c == '`':
		goto yystate34
	case c == 'f':
		goto yystate35
	case c == 't':
		goto yystate40
	case c >= '1' && c <= '9':
		goto yystate26
	case c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'e' || c >= 'g' && c <= 's' || c >= 'u' && c <= 'z':
		goto yystate32
	}

yystate2:
	c = l.next()
	goto yyrule1

yystate3:
	c = l.next()
	goto yyrule18

yystate4:
	c = l.next()
	switch {
	default:
		goto yyrule2
	case c == '\t' || c == '\n' || c == '\r' || c == ' ':
		goto yystate5
	}

yystate5:
	c = l.next()
	switch {
	default:
		goto yyrule2
	case c == '\t' || c == '\n' || c == '\r' || c == ' ':
		goto yystate5
	}

yystate6:
	c = l.next()
	goto yyrule7

yystate7:
	c = l.next()
	switch {
	default:
		goto yyrule18
	case c == '^':
		goto yystate8
	}

yystate8:
	c = l.next()
	goto yyrule12

yystate9:
	c = l.next()
	switch {
	default:
		goto yyrule18
	case c == '\'':
		goto yystate11
	case c == '\\':
		goto yystate12
	case c >= '\x01' && c <= '&' || c >= '(' && c <= '[' || c >= ']' && c <= 'ÿ':
		goto yystate10
	}

yystate10:
	c = l.next()
	switch {
	default:
		goto yyabort
	case c == '\'':
		goto yystate11
	case c == '\\':
		goto yystate12
	case c >= '\x01' && c <= '&' || c >= '(' && c <= '[' || c >= ']' && c <= 'ÿ':
		goto yystate10
	}

yystate11:
	c = l.next()
	goto yyrule9

yystate12:
	c = l.next()
	switch {
	default:
		goto yyabort
	case c == '\'':
		goto yystate13
	case c == '\\':
		goto yystate12
	case c >= '\x01' && c <= '&' || c >= '(' && c <= '[' || c >= ']' && c <= 'ÿ':
		goto yystate10
	}

yystate13:
	c = l.next()
	switch {
	default:
		goto yyrule9
	case c == '\'':
		goto yystate11
	case c == '\\':
		goto yystate12
	case c >= '\x01' && c <= '&' || c >= '(' && c <= '[' || c >= ']' && c <= 'ÿ':
		goto yystate10
	}

yystate14:
	c = l.next()
	switch {
	default:
		goto yyrule18
	case c >= '0' && c <= '9':
		goto yystate15
	}

yystate15:
	c = l.next()
	switch {
	default:
		goto yyrule6
	case c == 'E' || c == 'e':
		goto yystate16
	case c == 'i':
		goto yystate19
	case c >= '0' && c <= '9':
		goto yystate15
	}

yystate16:
	c = l.next()
	switch {
	default:
		goto yyabort
	case c == '+' || c == '-':
		goto yystate17
	case c >= '0' && c <= '9':
		goto yystate18
	}

yystate17:
	c = l.next()
	switch {
	default:
		goto yyabort
	case c >= '0' && c <= '9':
		goto yystate18
	}

yystate18:
	c = l.next()
	switch {
	default:
		goto yyrule6
	case c == 'i':
		goto yystate19
	case c >= '0' && c <= '9':
		goto yystate18
	}

yystate19:
	c = l.next()
	goto yyrule4

yystate20:
	c = l.next()
	switch {
	default:
		goto yyrule5
	case c == '.':
		goto yystate15
	case c == '8' || c == '9':
		goto yystate22
	case c == 'E' || c == 'e':
		goto yystate16
	case c == 'X' || c == 'x':
		goto yystate24
	case c == 'i':
		goto yystate23
	case c >= '0' && c <= '7':
		goto yystate21
	}

yystate21:
	c = l.next()
	switch {
	default:
		goto yyrule5
	case c == '.':
		goto yystate15
	case c == '8' || c == '9':
		goto yystate22
	case c == 'E' || c == 'e':
		goto yystate16
	case c == 'i':
		goto yystate23
	case c >= '0' && c <= '7':
		goto yystate21
	}

yystate22:
	c = l.next()
	switch {
	default:
		goto yyabort
	case c == '.':
		goto yystate15
	case c == 'E' || c == 'e':
		goto yystate16
	case c == 'i':
		goto yystate23
	case c >= '0' && c <= '9':
		goto yystate22
	}

yystate23:
	c = l.next()
	goto yyrule3

yystate24:
	c = l.next()
	switch {
	default:
		goto yyabort
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'F' || c >= 'a' && c <= 'f':
		goto yystate25
	}

yystate25:
	c = l.next()
	switch {
	default:
		goto yyrule5
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'F' || c >= 'a' && c <= 'f':
		goto yystate25
	}

yystate26:
	c = l.next()
	switch {
	default:
		goto yyrule5
	case c == '.':
		goto yystate15
	case c == 'E' || c == 'e':
		goto yystate16
	case c == 'i':
		goto yystate23
	case c >= '0' && c <= '9':
		goto yystate27
	}

yystate27:
	c = l.next()
	switch {
	default:
		goto yyrule5
	case c == '.':
		goto yystate15
	case c == 'E' || c == 'e':
		goto yystate16
	case c == 'i':
		goto yystate23
	case c >= '0' && c <= '9':
		goto yystate27
	}

yystate28:
	c = l.next()
	switch {
	default:
		goto yyrule18
	case c == '<':
		goto yystate29
	}

yystate29:
	c = l.next()
	goto yyrule13

yystate30:
	c = l.next()
	switch {
	default:
		goto yyrule18
	case c == '>':
		goto yystate31
	}

yystate31:
	c = l.next()
	goto yyrule14

yystate32:
	c = l.next()
	switch {
	default:
		goto yyrule17
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate33
	}

yystate33:
	c = l.next()
	switch {
	default:
		goto yyrule17
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate33
	}

yystate34:
	c = l.next()
	goto yyrule8

yystate35:
	c = l.next()
	switch {
	default:
		goto yyrule17
	case c == 'a':
		goto yystate36
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'b' && c <= 'z':
		goto yystate33
	}

yystate36:
	c = l.next()
	switch {
	default:
		goto yyrule17
	case c == 'l':
		goto yystate37
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'k' || c >= 'm' && c <= 'z':
		goto yystate33
	}

yystate37:
	c = l.next()
	switch {
	default:
		goto yyrule17
	case c == 's':
		goto yystate38
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'r' || c >= 't' && c <= 'z':
		goto yystate33
	}

yystate38:
	c = l.next()
	switch {
	default:
		goto yyrule17
	case c == 'e':
		goto yystate39
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z':
		goto yystate33
	}

yystate39:
	c = l.next()
	switch {
	default:
		goto yyrule15
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate33
	}

yystate40:
	c = l.next()
	switch {
	default:
		goto yyrule17
	case c == 'r':
		goto yystate41
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'q' || c >= 's' && c <= 'z':
		goto yystate33
	}

yystate41:
	c = l.next()
	switch {
	default:
		goto yyrule17
	case c == 'u':
		goto yystate42
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 't' || c >= 'v' && c <= 'z':
		goto yystate33
	}

yystate42:
	c = l.next()
	switch {
	default:
		goto yyrule17
	case c == 'e':
		goto yystate43
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z':
		goto yystate33
	}

yystate43:
	c = l.next()
	switch {
	default:
		goto yyrule16
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate33
	}

	goto yystate44 // silence unused label error
yystate44:
	c = l.next()
yystart44:
	switch {
	default:
		goto yystate45 // c >= '\x01' && c <= '!' || c >= '#' && c <= '[' || c >= ']' && c <= 'ÿ'
	case c == '"':
		goto yystate46
	case c == '\\':
		goto yystate47
	case c == '\x00':
		goto yystate2
	}

yystate45:
	c = l.next()
	switch {
	default:
		goto yyabort
	case c == '"':
		goto yystate46
	case c == '\\':
		goto yystate47
	case c >= '\x01' && c <= '!' || c >= '#' && c <= '[' || c >= ']' && c <= 'ÿ':
		goto yystate45
	}

yystate46:
	c = l.next()
	goto yyrule10

yystate47:
	c = l.next()
	switch {
	default:
		goto yyabort
	case c == '"':
		goto yystate48
	case c == '\\':
		goto yystate47
	case c >= '\x01' && c <= '!' || c >= '#' && c <= '[' || c >= ']' && c <= 'ÿ':
		goto yystate45
	}

yystate48:
	c = l.next()
	switch {
	default:
		goto yyrule10
	case c == '"':
		goto yystate46
	case c == '\\':
		goto yystate47
	case c >= '\x01' && c <= '!' || c >= '#' && c <= '[' || c >= ']' && c <= 'ÿ':
		goto yystate45
	}

	goto yystate49 // silence unused label error
yystate49:
	c = l.next()
yystart49:
	switch {
	default:
		goto yystate50 // c >= '\x01' && c <= '_' || c >= 'a' && c <= 'ÿ'
	case c == '\x00':
		goto yystate2
	case c == '`':
		goto yystate51
	}

yystate50:
	c = l.next()
	switch {
	default:
		goto yyabort
	case c == '`':
		goto yystate51
	case c >= '\x01' && c <= '_' || c >= 'a' && c <= 'ÿ':
		goto yystate50
	}

yystate51:
	c = l.next()
	goto yyrule11

yyrule1: // \0
	{
		return 0
	}
yyrule2: // [ \t\n\r]+

	goto yystate0
yyrule3: // {imaginary_ilit}
	{
		return l.int(lval, true)
	}
yyrule4: // {imaginary_lit}
	{
		return l.float(lval, true)
	}
yyrule5: // {int_lit}
	{
		return l.int(lval, false)
	}
yyrule6: // {float_lit}
	{
		return l.float(lval, false)
	}
yyrule7: // \"
	{
		l.sc = S1
		goto yystate0
	}
yyrule8: // `
	{
		l.sc = S2
		goto yystate0
	}
yyrule9: // '(\\.|[^'])*'
	{
		if ret := l.str(lval, ""); ret != STR {
			return ret
		}
		lval.item = int32(lval.item.(string)[0])
		return INTEGER
	}
yyrule10: // (\\.|[^\"])*\"
	{
		return l.str(lval, "\"")
	}
yyrule11: // ([^`]|\n)*`
	{
		return l.str(lval, "`")
	}
yyrule12: // "&^"
	{
		return ANDNOT
	}
yyrule13: // "<<"
	{
		return LSH
	}
yyrule14: // ">>"
	{
		return RSH
	}
yyrule15: // false
	{
		lval.item = false
		return BOOLEAN
	}
yyrule16: // true
	{
		lval.item = true
		return BOOLEAN
	}
yyrule17: // {ident}
	{
		lval.item = string(l.val)
		return IDENTIFIER
	}
yyrule18: // .
	{
		return c0
	}
	panic("unreachable")

	goto yyabort // silence unused label error

yyabort: // no lexem recognized
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
