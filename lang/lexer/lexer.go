package lexer

import (
	"klc/lang/token"
	"strings"
)

type Lexer struct {
	input  []rune
	cursor int

	cursorLine   int
	cursorColumn int
	eof          *token.Token

	queue []*token.Token
	cur   *token.Token

	sbuilder *strings.Builder
}

func New(input string) *Lexer {
	l := &Lexer{
		input:        []rune(input),
		cursorLine:   1,
		cursorColumn: 1,
		sbuilder:     &strings.Builder{},
	}
	l.next()
	return l
}

func (l *Lexer) All() []*token.Token {
	var tokens = []*token.Token{}

	for {
		t := l.Current()
		tokens = append(tokens, t)

		if t.Is(token.Eof) || t.Is(token.Invalid) {
			break
		}
		l.Next()
	}

	return tokens
}

func (l *Lexer) Current() *token.Token {
	if l.cur == nil {
		l.Next()
	}

	return l.cur
}

func (l *Lexer) Next() *token.Token {
	if len(l.queue) > 0 {
		l.cur = l.queue[0]
		l.queue = l.queue[1:]

		if len(l.queue) == 0 {
			l.next()
		}

		return l.cur
	}

	return l.eof
}

func (l *Lexer) Peek() *token.Token {
	if len(l.queue) > 0 {
		return l.queue[0]
	}

	return l.eof
}

func (l *Lexer) PeekN(n int) *token.Token {
	if len(l.queue) <= n {
		if l.eof != nil {
			return l.eof
		}

		for i := 0; i < n-len(l.queue); i++ {
			l.next()

			if l.eof != nil {
				return l.eof
			}
		}
	}

	return l.queue[n]
}

// ----------------------------------------------------------------------------
// Parsing
// ----------------------------------------------------------------------------
func (l *Lexer) next() {
	if l.eof != nil {
		return
	}

	var t *token.Token
	c := l.char()
	if isIgnored(c) {
		l.skipIgnored()
		c = l.char()
	}

	switch c {
	case 0: // eof
		t = l.parseEof()
	case '\n':
		t = l.parseNewline()
	case '#':
		t = l.parseComment()
	case '{':
		t = l.parseLbrace()
	case '}':
		t = l.parseRbrace()
	case '(':
		t = l.parseLparen()
	case ')':
		t = l.parseRparen()
	case '[':
		t = l.parseLbracket()
	case ']':
		t = l.parseRbracket()
	case ';':
		t = l.parseSemicolon()
	case ',':
		t = l.parseComma()
	case ':':
		t = l.parseColon()
	case '?':
		t = l.parseQuestion()
	case '.':
		t = l.parseDot()
	case '=':
		t = l.parseAssignment()
	case '\\':
		t = l.parseBackslash()
		if t == nil { // ignored the line break
			l.next()
			return
		}
	case '\'':
		t = l.parseString()
	default:
		if isLetter(c) {
			t = l.parseIdentifier()
		} else if isDigit(c) {
			t = l.parseNumber()
		} else {
			t = l.parseOperator()
		}

		if t == nil {
			t = l.parseInvalid()
		}
	}

	if t != nil {
		l.queue = append(l.queue, t)
	}
}

func (l *Lexer) at() token.At {
	return token.At{
		Cursor: l.cursor,
		Line:   l.cursorLine,
		Column: l.cursorColumn,
	}
}

func (l *Lexer) skip() {
	l.cursor++
	l.cursorColumn++
}

func (l *Lexer) char() rune {
	if l.cursor >= len(l.input) {
		return 0
	}

	return l.input[l.cursor]
}

func (l *Lexer) peekChar(i int) rune {
	if (l.cursor + i) >= len(l.input) {
		return 0
	}

	return l.input[l.cursor+i]
}

func (l *Lexer) skipIgnored() {
	for isIgnored(l.char()) {
		l.skip()
	}
}

func (l *Lexer) parseEof() *token.Token {
	t := token.Create(token.Eof, "", l.at())
	l.eof = t
	return t
}

func (l *Lexer) parseInvalid() *token.Token {
	t := token.Create(token.Invalid, string(l.char()), l.at())
	return t
}

func (l *Lexer) parseNewline() *token.Token {
	t := token.Create(token.Newline, string(l.char()), l.at())
	l.skip()
	l.cursorLine++
	l.cursorColumn = 1
	return t
}

func (l *Lexer) parseComment() *token.Token {
	l.sbuilder.Reset()
	t := token.CreateEmpty(token.Comment, l.at())
	c := l.char()

	for c != '\n' && c != 0 {
		l.sbuilder.WriteRune(l.char())
		l.skip()
		c = l.char()
	}

	t.Literal = l.sbuilder.String()
	return t
}

func (l *Lexer) parseLbrace() *token.Token {
	t := token.Create(token.Lbrace, string(l.char()), l.at())
	l.skip()
	return t
}

func (l *Lexer) parseRbrace() *token.Token {
	t := token.Create(token.Rbrace, string(l.char()), l.at())
	l.skip()
	return t
}

func (l *Lexer) parseLparen() *token.Token {
	t := token.Create(token.Lparen, string(l.char()), l.at())
	l.skip()
	return t
}

func (l *Lexer) parseRparen() *token.Token {
	t := token.Create(token.Rparen, string(l.char()), l.at())
	l.skip()
	return t
}

func (l *Lexer) parseLbracket() *token.Token {
	t := token.Create(token.Lbracket, string(l.char()), l.at())
	l.skip()
	return t
}

func (l *Lexer) parseRbracket() *token.Token {
	t := token.Create(token.Rbracket, string(l.char()), l.at())
	l.skip()
	return t
}

func (l *Lexer) parseSemicolon() *token.Token {
	t := token.Create(token.Semicolon, string(l.char()), l.at())
	l.skip()
	return t
}

func (l *Lexer) parseComma() *token.Token {
	t := token.Create(token.Comma, string(l.char()), l.at())
	l.skip()
	return t
}

func (l *Lexer) parseColon() *token.Token {
	t := token.Create(token.Colon, string(l.char()), l.at())
	l.skip()
	return t
}

func (l *Lexer) parseQuestion() *token.Token {
	t := token.Create(token.Question, string(l.char()), l.at())
	l.skip()
	return t
}

func (l *Lexer) parseDot() *token.Token {
	pos := l.at()

	left := l.char()
	l.skip()

	mid := l.char()
	if mid != '.' {
		return token.Create(token.Dot, string(left), l.at())
	}

	l.skip()
	right := l.char()
	if right != '.' {
		return l.parseInvalid()
	}

	l.skip()
	return token.Create(token.Spread, "...", pos)
}

func (l *Lexer) parseAssignment() *token.Token {
	pos := l.at()

	left := l.char()
	l.skip()

	right := l.char()
	if right == '>' {
		l.skip()
		return token.Create(token.Arrow, "=>", pos)

	} else if isDoubleOperator(left, right) {
		l.skip()
		return token.Create(token.Operator, string(left)+string(right), pos)
	}

	return token.Create(token.Assignment, string(left), pos)
}

func (l *Lexer) parseBackslash() *token.Token {
	l.skip()
	l.skipIgnored()
	if l.char() == '\n' {
		l.skip()
		l.cursorLine++
		l.cursorColumn = 1
		return nil
	}

	return l.parseInvalid()
}

func (l *Lexer) parseOperator() *token.Token {
	pos := l.at()
	left := l.char()

	if isOperator(left) {
		l.skip()

		right := l.char()
		if isDoubleOperator(left, right) {
			l.skip()
			return token.Create(token.Operator, string(left)+string(right), pos)
		}
		if isCompositeAssignment(left, right) {
			l.skip()
			return token.Create(token.Assignment, string(left), pos)
		}

		return token.Create(token.Operator, string(left), pos)
	}

	return nil
}

func (l *Lexer) parseString() *token.Token {
	l.sbuilder.Reset()
	pos := l.at()

	l.skip()

	escaping := false
	for {
		c := l.char()
		if escaping {
			l.sbuilder.WriteRune(c)
			escaping = false
			l.skip()
			continue
		}

		if !escaping && c == '\\' {
			escaping = true
			l.skip()
			continue
		}

		if c == '\'' {
			l.skip()
			break
		}

		if c == '\n' || c == 0 {
			return l.parseInvalid()
		}

		l.sbuilder.WriteRune(c)
		l.skip()
	}

	return token.Create(token.String, l.sbuilder.String(), pos)
}

func (l *Lexer) parseNumber() *token.Token {
	l.sbuilder.Reset()
	pos := l.at()

	l.sbuilder.WriteRune(l.char())
	l.skip()

	c := l.char()
	hasDot := false
	hasExp := false
	for {
		if c == '.' {
			if hasDot {
				break
			}
			hasDot = true
			l.sbuilder.WriteRune(c)
			l.skip()
		} else if !hasExp && (c == 'e' || c == 'E' || c == '.') {
			nxt := l.peekChar(1)
			if isDigit(nxt) || nxt == '+' || nxt == '-' {
				hasDot = false
				hasExp = true
				l.sbuilder.WriteRune(c)
				l.skip()
			} else {
				break
			}

			c = l.char()
			if c == '+' || c == '-' {
				l.sbuilder.WriteRune(c)
				l.skip()
			}
		}

		c = l.char()
		if isDigit(c) {
			l.sbuilder.WriteRune(c)
			l.skip()
		} else {
			break
		}

		c = l.char()
	}

	return token.Create(token.Number, l.sbuilder.String(), pos)
}

func (l *Lexer) parseIdentifier() *token.Token {
	l.sbuilder.Reset()
	pos := l.at()

	l.sbuilder.WriteRune(l.char())
	l.skip()

	c := l.char()
	for isLetter(c) || isDigit(c) {
		l.sbuilder.WriteRune(c)
		l.skip()
		c = l.char()
	}

	return token.Create(token.Identifier, l.sbuilder.String(), pos)
}

// ----------------------------------------------------------------------------
// Helpers
// ----------------------------------------------------------------------------
func isIgnored(r rune) bool {
	return r == ' ' || r == '\t' || r == '\r'
}

func isLetter(r rune) bool {
	return (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || r == '_'
}

func isDigit(r rune) bool {
	return r >= '0' && r <= '9'
}

func isOperator(a rune) bool {
	switch a {
	case '+', '-', '*', '/', '%', '!', '>', '<', '&', '|', '^':
		return true
	}

	return false
}

func isDoubleOperator(a rune, b rune) bool {
	switch a {
	case '+':
		return b == '+'
	case '*':
		return b == '*'
	case '/':
		return b == '/'
	case '!':
		return b == '=' || b == '|' || b == '&'
	case '>':
		return b == '='
	case '<':
		return b == '='
	case '=':
		return b == '='
	case '|':
		return b == '|'
	case '&':
		return b == '&'
	case '^':
		return b == '^'
	}

	return false
}

func isCompositeAssignment(a rune, b rune) bool {
	return b == '=' && (a == '+' || a == '-' || a == '*' || a == '/')
}
