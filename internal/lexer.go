package internal

import (
	"github.com/renatopp/langtools"
	"github.com/renatopp/langtools/runes"
	"github.com/renatopp/langtools/tokenizers"
)

type KlcLexer struct {
	*langtools.Lexer
}

func NewLexer(input []byte) *KlcLexer {
	k := &KlcLexer{}
	k.Lexer = langtools.NewLexer(input, k.tokenizer)
	return k
}

func (k *KlcLexer) matches(nexts ...rune) bool {
	for i, r := range nexts {
		c := k.PeekCharAt(i)
		if !c.Is(r) {
			return false
		}
	}
	return true
}

func (k *KlcLexer) tokenizer(l *langtools.Lexer) langtools.Token {
	for {
		if l.HasTooManyErrors() {
			return langtools.NewToken(langtools.TUnknown, "", 0, 0)
		}

		c0 := l.PeekCharAt(0)
		c1 := l.PeekCharAt(1)

		switch {
		case k.matches('-', '-'):
			l.EatChar()
			l.EatChar()
			return tokenizers.EatUntilEndOfLine(l, TComment)

		case runes.IsAlpha(c0.Rune):
			t := tokenizers.EatIdentifier(l, TIdentifier)

			if t.Literal == "or" ||
				t.Literal == "and" {
				return langtools.NewToken(TOperand, t.Literal, t.Line, t.Column)
			}

			if t.Literal == "to" {
				return langtools.NewToken(TKeyword, t.Literal, t.Line, t.Column)
			}

			return t

		case runes.IsDigit(c0.Rune) || c0.Is('.') && runes.IsDigit(c1.Rune):
			return tokenizers.EatNumber(l, TNumber)

		case k.matches('+', '=') ||
			k.matches('-', '=') ||
			k.matches('*', '=') ||
			k.matches('/', '=') ||
			k.matches('%', '=') ||
			k.matches('^', '='):
			first := l.EatChar()
			second := l.EatChar()
			return langtools.NewToken(TAssign, string(first.Rune)+string(second.Rune), first.Line, first.Column)

		case k.matches('=', '=') ||
			k.matches('!', '=') ||
			k.matches('<', '=') ||
			k.matches('>', '=') ||
			k.matches('!', '=') ||
			k.matches('=', '='):
			first := l.EatChar()
			second := l.EatChar()
			return langtools.NewToken(TOperand, string(first.Rune)+string(second.Rune), first.Line, first.Column)

		case c0.Is('+') ||
			c0.Is('-') ||
			c0.Is('*') ||
			c0.Is('/') ||
			c0.Is('%') ||
			c0.Is('^') ||
			c0.Is('<') ||
			c0.Is('>') ||
			c0.Is('!'):
			return langtools.NewToken(TOperand, string(l.EatChar().Rune), c0.Line, c0.Column)

		case c0.Is('='):
			return langtools.NewToken(TAssign, string(l.EatChar().Rune), c0.Line, c0.Column)

		case c0.Is('('):
			return langtools.NewToken(TLParen, string(l.EatChar().Rune), c0.Line, c0.Column)

		case c0.Is(')'):
			return langtools.NewToken(TRParen, string(l.EatChar().Rune), c0.Line, c0.Column)

		case c0.Is(','):
			return langtools.NewToken(TComma, string(l.EatChar().Rune), c0.Line, c0.Column)

		case c0.Is('\r'):
			l.EatChar()

		case c0.Is('\n') || c0.Is(';'):
			return langtools.NewToken(TEol, string(l.EatChar().Rune), c0.Line, c0.Column)

		case runes.IsWhitespace(c0.Rune):
			tokenizers.EatWhitespaces(l, langtools.TUnknown)

		default:
			c := l.EatChar()
			return langtools.NewToken(langtools.TUnknown, string(c.Rune), c.Line, c.Column)

		}
	}
}
