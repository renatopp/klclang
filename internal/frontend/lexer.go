package frontend

import (
	"strconv"
	"strings"

	"github.com/renatopp/klclang/internal/core"
	"github.com/renatopp/x/strx"
)

func Lex(module *core.Module, content []byte) ([]*core.Token, error) {
	l := &lexer{
		module:         module,
		scanner:        NewScanner([]rune(string(content)), rune(0)),
		indentStack:    ds.NewStack[*identinfo](),
		syntheticQueue: ds.NewQueue[*core.Token](),
		result:         make([]*core.Token, 0),
		previousRune:   0,
		currentLine:    1,
		currentColumn:  1,
		previousLine:   1,
		previousColumn: 1,
	}
	return core.WithRecover(l.lex)
}

type identinfo struct {
	indent int  // indent level
	ommit  bool // if the indent was caused by a {, [, ( or comma, the indent can be ommited
}

type lexer struct {
	module         *core.Module
	scanner        *Scanner[rune]
	indentStack    ds.Stack[*identinfo]
	syntheticQueue ds.Queue[*core.Token]
	result         []*core.Token
	currentLine    int
	currentColumn  int
	previousRune   int
	previousLine   int
	previousColumn int
}

// lex performs the lexical analysis of the input content, producing a list of tokens.
func (l *lexer) lex() []*core.Token {
	for !l.last().IsKind(core.TokenEof) {
		l.result = append(l.result, l.next())
	}

	return l.result
}

// next returns the next token from the input stream, handling indentation
// levels and generating synthetic tokens as needed.
func (l *lexer) next() *core.Token {
	for {
		if l.syntheticQueue.Size() > 0 {
			return l.syntheticQueue.Pop()
		}

		l.anchorPrevious()
		c0 := l.scanner.PeekAt(0)
		c1 := l.scanner.PeekAt(1)
		c2 := l.scanner.PeekAt(2)

		switch {
		// ------------------------------------------------------------------------
		// EOF | DEDENT
		// ------------------------------------------------------------------------
		// If the end of file is reached, we need to close all the open indents by
		// generating the corresponding } tokens.
		case runes.IsEof(c0):
			for l.indentStack.Size() > 0 {
				popped := l.indentStack.Pop()
				if !popped.ommit {
					return l.tok(core.TokenRightBrace, "}")
				}
			}
			return l.tok(core.TokenEof, "")

		// ------------------------------------------------------------------------
		// SPACES
		// ------------------------------------------------------------------------
		// Spaces are not relevant for the parser, but it may be relevant for error
		// reporting, code formatting tools or analysis.
		case runes.IsSpace(c0):
			l.eatSpaces()
			// return l.tok(core.TokenSpace, l.eatSpaces())

		// ------------------------------------------------------------------------
		// SEPARATORS
		// ------------------------------------------------------------------------
		// Handle newlines and semicolons as instruction separators. If they are
		// followed by an increase or decrease in the indentation level, we need
		// to generate the corresponding { or } tokens for brace illisions.
		case runes.IsOneOf(c0, '\n', ';'):
			sep, indent := l.eatSeparators()

			// indent increase
			prevIndent := l.indentStack.TopOr(&identinfo{})
			if indent > prevIndent.indent {
				curIndent := &identinfo{
					indent: indent,
					ommit: l.last().IsKind( // ommit if caused by {, [, ( or comma
						core.TokenLeftBrace,
						core.TokenLeftBracket,
						core.TokenLeftParen,
						core.TokenComma,
					),
				}
				l.indentStack.Push(curIndent)
				if !curIndent.ommit {
					return l.tok(core.TokenLeftBrace, "{")
				} else {
					return l.tok(core.TokenSeparator, sep)
				}
			}

			// indent decrease
			for indent < prevIndent.indent {
				popped := l.indentStack.Pop()
				if !popped.ommit {
					l.syntheticQueue.Push(l.tok(core.TokenRightBrace, "}"))
				}
				prevIndent = l.indentStack.TopOr(&identinfo{})
			}

			// check valid indent
			if indent > prevIndent.indent {
				l.anchorPrevious()
				core.ThrowAt(l.span(), core.ErrorSyntax, "invalid indentation level")
			}

			// return dedent
			if l.syntheticQueue.Size() > 0 {
				return l.syntheticQueue.Pop()
			}

			// return separator
			return l.tok(core.TokenSeparator, sep)

		// ------------------------------------------------------------------------
		// COMMENT
		// ------------------------------------------------------------------------
		// Comments start with # and go until the end of the line. Multiple lines
		// starting with # without empty lines between them are considered as part
		// of the same comment.
		case runes.IsOneOf(c0, '#'):
			comment := l.tok(core.TokenComment, l.eatComment())
			// Merge consecutive comments into a single token
			for runes.IsOneOf(l.scanner.Peek(), '#') {
				other := l.tok(core.TokenComment, l.eatComment())
				comment.Absorb(other)
			}
			return comment

		// ------------------------------------------------------------------------
		// IMPORT
		// ------------------------------------------------------------------------
		case l.last().IsKind(core.TokenImport):
			literal := l.eatComment() // import until the end of the line
			return l.tok(core.TokenString, strings.TrimSpace(literal))

		// ------------------------------------------------------------------------
		// IDENTIFIERS and KEYWORDS
		// ------------------------------------------------------------------------
		// Identifiers and keywords start with a letter or an underscore, and can
		// be followed by letters, digits or underscores.
		case runes.IsAlpha(c0) || runes.IsOneOf(c0, '_'):
			literal := l.eatIdentifier()

			// Keywords
			if kind := l.kindFromLiteral(literal); kind != core.TokenUnknown {
				return l.tok(kind, literal)
			}

			// Type identifiers
			if core.IsTypeName(literal) {
				return l.tok(core.TokenTypeIdent, literal)
			}

			// Value identifiers
			return l.tok(core.TokenValueIdent, literal)

		// ------------------------------------------------------------------------
		// NUMBERS
		// ------------------------------------------------------------------------
		// Numbers are tricky because they can be in different bases and formats.
		// We need to handle integers, floats, hex, octal and binary numbers.
		// The following cases are handled:
		//
		// - Hexadecimal: 0xFF or 0XFF
		// - Octal: 0o77 or 0O77
		// - Binary: 0b11 or 0B11
		// - Float: 123.45, .45, 123., 123e10, 123.45e10, 123e+10, 123e-10
		// - Integer: 12345
		case runes.IsDigit(c0) || c0 == '.' && runes.IsDigit(c1):
			// Number as value identifiers such as tuple accesss
			if c0 == '.' && l.last().IsKind(
				core.TokenValueIdent,
				core.TokenTypeIdent,
				core.TokenRightBrace,
				core.TokenRightBracket,
				core.TokenRightParen,
			) {
				token := l.tok(core.TokenDot, string(l.eat()))
				l.anchorPrevious()
				l.syntheticQueue.Push(l.tok(core.TokenValueIdent, l.eatIdentifier()))
				return token
			}

			// Hexadecimal
			if c0 == '0' && runes.IsOneOf(c1, 'x', 'X') {
				return l.tok(core.TokenHex, l.eatHexadecimal())
			}

			// Octal
			if c0 == '0' && runes.IsOneOf(c1, 'o', 'O') {
				return l.tok(core.TokenOct, l.eatOctal())
			}

			// Binary
			if c0 == '0' && runes.IsOneOf(c1, 'b', 'B') {
				return l.tok(core.TokenBin, l.eatBinary())
			}

			// Float
			number := l.eatNumber()
			if strings.Contains(number, ".") || strings.Contains(number, "e") {
				return l.tok(core.TokenFloat, number)
			}

			// Integer
			return l.tok(core.TokenInt, number)

		// ------------------------------------------------------------------------
		// STRING LITERALS
		// ------------------------------------------------------------------------
		// Strings can be defined using double quotes (").
		case runes.IsOneOf(c0, '"'):
			return l.tok(core.TokenString, l.eatString())

		// ------------------------------------------------------------------------
		// OPERATORS AND ERROR CASES
		// ------------------------------------------------------------------------
		default:
			s1 := string(c0)
			s2 := s1 + string(c1)
			s3 := s2 + string(c2)

			// 3-char operators
			if tok := l.kindFromLiteral(s3); tok != core.TokenUnknown {
				l.eat()
				l.eat()
				l.eat()
				return l.tok(tok, s3)
			}

			// 2-char operators
			if tok := l.kindFromLiteral(s2); tok != core.TokenUnknown {
				l.eat()
				l.eat()
				return l.tok(tok, s2)
			}

			// 1-char operators
			if tok := l.kindFromLiteral(s1); tok != core.TokenUnknown {
				l.eat()
				return l.tok(tok, s1)
			}

			// Unknown character
			l.eat()
			core.ThrowAt(l.span(), core.ErrorSyntax, "invalid character: '%s'", strx.Escape(s1))
		}
	}
}

// kindFromLiteral returns the token kind for a given literal string. If the
// literal does not correspond to any known token kind, it returns
// core.TokenUnknown.
func (l *lexer) kindFromLiteral(literal string) core.TokenKind {
	kind, ok := core.TokenKindMap[literal]
	if ok {
		return kind
	}
	return core.TokenUnknown
}

// span returns the current span of the lexer, from the previous rune to the
// current cursor position.
func (l *lexer) span() *core.Span {
	return &core.Span{
		Module:     l.module,
		From:       l.previousRune,
		To:         l.scanner.Cursor(),
		FromLine:   l.currentLine,
		FromColumn: l.currentColumn,
		ToLine:     l.previousLine,
		ToColumn:   l.previousColumn,
	}
}

// anchorPrevious updates the previous rune, line and column to the current
// position of the lexer.
func (l *lexer) anchorPrevious() {
	l.previousRune = l.scanner.Cursor()
	l.previousLine = l.currentLine
	l.previousColumn = l.currentColumn
}

// tok creates a new token with the given kind and literal, using the current
// span of the lexer.
func (l *lexer) tok(kind core.TokenKind, literal string) *core.Token {
	return core.NewToken(l.span(), kind, literal)
}

// last returns the last token generated by the lexer.
func (l *lexer) last() *core.Token {
	if len(l.result) == 0 {
		return nil
	}
	return l.result[len(l.result)-1]
}

// eat consumes a single character from the input, updating the current line
func (l *lexer) eat() rune {
	r := l.scanner.Eat()
	l.currentColumn++
	if r == '\n' {
		l.currentLine++
		l.currentColumn = 1
	}
	return r
}

// eatSpaces consumes all the characters that composes a space. Including space, tab and
// carriage return.
func (l *lexer) eatSpaces() string {
	res := ""
	for {
		c := l.scanner.Peek()
		if !runes.IsSpace(c) {
			break
		}
		res += string(c)
		l.eat()
	}
	return res
}

// eatNewlines consumes all the characters that composes a newline. Including only the
// newline character (without the carriage return).
func (l *lexer) eatNewlines() string {
	res := ""
	for {
		c := l.scanner.Peek()
		if runes.IsNewline(c) {
			res += string(c)
		} else if runes.IsSpace(c) {
			// pass
		} else {
			break

		}
		l.eat()
	}
	return res
}

// eatSeparators consumes all the characters that composes a separator (newline and ;).
func (l *lexer) eatSeparators() (string, int) {
	res := ""
	indent := -1 // disabled, only enable if char is \n
out:
	for {
		c := l.scanner.Peek()
		switch {
		case runes.IsOneOf(c, '\n'):
			indent = 0
			res += string(c)
		case runes.IsOneOf(c, ';'):
			indent = -1 // disable
			res += string(c)
		case runes.IsSpace(c):
			if indent >= 0 {
				indent++
			}
		default:
			break out
		}
		l.eat()
	}
	return res, max(0, indent)
}

// eatComment consumes all the characters until the end of the line or the end of file.
func (l *lexer) eatComment() string {
	res := ""
	for {
		c := l.eat()
		if runes.IsOneOf(c, '\r') {
			continue
		}
		res += string(c)
		if runes.IsOneOf(c, '\n', 0) {
			break
		}
	}
	return res
}

// eatIdentifier consumes all the characters that composes an common identifier. Including
// letters, digits and underscores.
func (l *lexer) eatIdentifier() string {
	res := ""
	for {
		c := l.scanner.Peek()
		if !runes.IsAlphaNumeric(c) && c != '_' {
			break
		}

		res += string(c)
		l.eat()
	}
	return res
}

// eatNumber consumes all the characters that composes all common cases of numbers.
// Including:
//
// - Integers: `123_000`
// - Floats: `123.32`
// - Exponents: `123e32`
// - Floats with Exponents: `123.32e32`
// - Exponents with signal: `123e+32`
func (l *lexer) eatNumber() string {
	res := ""
	dot := false
	exp := false
	for {
		c := l.scanner.Peek()
		switch {
		case c == '_':
			l.eat()
			continue

		case c == '.':
			if dot || exp {
				core.ThrowAt(l.span(), core.ErrorSyntax, "unexpected '.' character")
			}
			dot = true
			res += string(c)

		case runes.IsOneOf(c, 'f', 'F'):
			if exp {
				core.ThrowAt(l.span(), core.ErrorSyntax, "unexpected 'f' character")
			}
			if !dot {
				res += ".0"
			}
			l.eat() // ignore the 'f' character
			return res

		case runes.IsOneOf(c, 'e', 'E'):
			if exp {
				core.ThrowAt(l.span(), core.ErrorSyntax, "unexpected 'e' character")
			}
			exp = true
			res += string(c)

			next := l.scanner.PeekAt(1)
			if runes.IsOneOf(next, '+', '-') {
				l.eat()
				res += string(next)
			}

		case runes.IsDigit(c):
			res += string(c)

		default:
			return res
		}

		l.eat()
	}
}

// eatHexadecimal consumes all the characters that composes a hexadecimal number. Considering
// `0x` or `0X` as optional prefix.
func (l *lexer) eatHexadecimal() string {
	res := ""
	c0 := l.scanner.PeekAt(0)
	c1 := l.scanner.PeekAt(1)
	if runes.IsOneOf(c0, 'x', 'X') {
		l.eat()
	}
	if runes.IsOneOf(c1, 'x', 'X') {
		l.eat()
		l.eat()
	}
	for {
		c := l.scanner.Peek()
		if !runes.IsHexadecimal(c) {
			break
		}
		res += string(c)
		l.eat()
	}
	return res
}

// eatOctal consumes all the characters that composes an octal number. Considering `0`
// as optional prefix.
func (l *lexer) eatOctal() string {
	res := ""
	c0 := l.scanner.PeekAt(0)
	c1 := l.scanner.PeekAt(1)
	if runes.IsOneOf(c0, 'o', 'O') {
		l.eat()
	}
	if runes.IsOneOf(c1, 'o', 'O') {
		l.eat()
		l.eat()
	}
	for {
		c := l.scanner.Peek()
		if !runes.IsOctal(c) {
			break
		}
		res += string(c)
		l.eat()
	}
	return res
}

// eatBinary consumes all the characters that composes a binary number. Considering `0b`
// or `0B` as optional prefix.
func (l *lexer) eatBinary() string {
	res := ""
	c0 := l.scanner.PeekAt(0)
	c1 := l.scanner.PeekAt(1)
	if runes.IsOneOf(c0, 'b', 'B') {
		l.eat()
	}
	if runes.IsOneOf(c1, 'b', 'B') {
		l.eat()
		l.eat()
	}
	for {
		c := l.scanner.Peek()
		if !runes.IsBinary(c) {
			break
		}
		res += string(c)
		l.eat()
	}
	return res
}

// eatString consumes all the characters that composes a string. This function will
// consider the first character as the delimiter of the string, and will stop
// when it finds the same character again. If the string is not closed, it will
// register an error.
//
// This function will process escape characters such as \n, \t, \", etc.
func (l *lexer) eatString() string {
	res := ""
	escaping := false
	first := l.eat()
	for {
		c := l.scanner.Peek()

		if runes.IsOneOf(c, '\r') {
			l.eat()
			continue
		}

		if runes.IsEof(c) {
			core.ThrowAt(l.span(), core.ErrorSyntax, "unexpected end of file")
			break
		} else if runes.IsOneOf(c, '\n') {
			core.ThrowAt(l.span(), core.ErrorSyntax, "unexpected new line")
			break
		}

		if !escaping && c == first {
			break
		}

		if !escaping && c == '\\' {
			escaping = true
			l.eat()
			continue
		}

		if escaping && c != first {
			escaping = false
			r, err := strconv.Unquote(`"\` + string(c) + `"`)
			if err != nil {
				core.ThrowAt(l.span(), core.ErrorSyntax, "%v", err.Error())
			}
			c = []rune(r)[0]
		}

		res += string(c)
		l.eat()
	}
	l.eat()
	return res
}

// eatRawString consumes all the characters that composes a raw string. This function will
// consider the first character as the delimiter of the string, and will stop
// when it finds the same character again. If the string is not closed, it will
// register an error.
//
// This function will record the string as it was written, meaning that any
// escape character will be kept as it is, including new lines. If you need to
// ignore new lines, use `EatString` instead.
func (l *lexer) eatRawString() string {
	res := ""
	escaping := false
	first := l.eat()
	for {
		c := l.scanner.Peek()

		if runes.IsOneOf(c, '\r') {
			l.eat()
			continue
		}

		if runes.IsEof(c) {
			core.ThrowAt(l.span(), core.ErrorSyntax, "unexpected end of file")
			break
		}

		if !escaping && c == first {
			break
		}

		if !escaping && c == '\\' {
			escaping = true
			l.eat()
			continue
		}

		if escaping && c != first {
			escaping = false
			r, err := strconv.Unquote(`"\` + string(c) + `"`)
			if err != nil {
				core.ThrowAt(l.span(), core.ErrorSyntax, "%v", err.Error())
			}
			c = []rune(r)[0]
		}

		res += string(c)
		l.eat()
	}
	l.eat()
	return res
}
