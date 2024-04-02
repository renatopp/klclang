package internal

import (
	"fmt"

	"github.com/renatopp/langtools/lexers"
	"github.com/renatopp/langtools/runes"
	"github.com/renatopp/langtools/tokens"
)

type KlcLexer struct {
	*lexers.BaseLexer
	pendingList []tokens.Token
}

func NewLexer(input []byte) *KlcLexer {
	k := &KlcLexer{}
	k.BaseLexer = lexers.NewBaseLexer(input, func(bl *lexers.BaseLexer) tokens.Token {
		return k.tokenizer()
	})
	k.pendingList = make([]tokens.Token, 0)
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

func (k *KlcLexer) tokenizer() tokens.Token {
	for {
		if k.HasTooManyErrors() {
			return tokens.NewToken(TUnknown, "", 0, 0)
		}

		if len(k.pendingList) > 0 {
			t := k.pendingList[0]
			k.pendingList = k.pendingList[1:]
			return t
		}

		c0 := k.PeekCharAt(0)
		c1 := k.PeekCharAt(1)

		switch {
		case c0.Rune == 0:
			return tokens.NewEofTokenAtChar(c0)

		// Comments
		case k.matches('-', '-'):
			k.EatChar()
			k.EatChar()
			return k.EatUntilEndOfLine().As(TComment)

		// Identifiers, Keywords and Logical Operators
		case runes.IsAlpha(c0.Rune):
			t := k.EatIdentifier()

			if t.IsOneOfLiterals("or", "and") {
				return t.As(TOperator)
			}

			if t.IsOneOfLiterals("to") {
				return t.As(TKeyword)
			}

			prev := k.PrevToken()
			if prev.IsType(TNumber) {
				k.pendingList = append(k.pendingList, t.As(TIdentifier))
				return tokens.NewToken(TOperator, "*", t.Line, t.Column)
			}

			return t.As(TIdentifier)

		// Numbers
		case runes.IsDigit(c0.Rune) || c0.Is('.') && runes.IsDigit(c1.Rune):
			return k.EatNumber().As(TNumber)

		case k.matches('+', '=') ||
			k.matches('-', '=') ||
			k.matches('*', '=') ||
			k.matches('/', '=') ||
			k.matches('%', '=') ||
			k.matches('^', '='):
			first := k.EatChar()
			second := k.EatChar()
			return tokens.NewToken(TAssign, string(first.Rune)+string(second.Rune), first.Line, first.Column)

		case k.matches('=', '=') ||
			k.matches('!', '=') ||
			k.matches('<', '=') ||
			k.matches('>', '=') ||
			k.matches('!', '=') ||
			k.matches('=', '='):
			first := k.EatChar()
			second := k.EatChar()
			return tokens.NewToken(TOperator, string(first.Rune)+string(second.Rune), first.Line, first.Column)

		case c0.Is('+') ||
			c0.Is('-') ||
			c0.Is('*') ||
			c0.Is('/') ||
			c0.Is('%') ||
			c0.Is('^') ||
			c0.Is('<') ||
			c0.Is('>') ||
			c0.Is('!'):
			return tokens.NewTokenAtChar(TOperator, string(k.EatChar().Rune), c0)

		case c0.Is('='):
			return tokens.NewTokenAtChar(TAssign, string(k.EatChar().Rune), c0)

		case c0.Is('('):
			t := tokens.NewTokenAtChar(TLParen, string(k.EatChar().Rune), c0)

			if k.PrevToken().IsType(TNumber) {
				k.pendingList = append(k.pendingList, t)
				return tokens.NewToken(TOperator, "*", t.Line, t.Column)
			}

			return t

		case c0.Is(')'):
			return tokens.NewTokenAtChar(TRParen, string(k.EatChar().Rune), c0)

		case c0.Is(','):
			return tokens.NewTokenAtChar(TComma, string(k.EatChar().Rune), c0)

		case c0.Is('\r'):
			k.EatChar()

		case c0.Is('\n') || c0.Is(';'):
			return tokens.NewToken(TEoe, string(k.EatChar().Rune), c0.Line, c0.Column)

		case runes.IsWhitespace(c0.Rune):
			k.EatWhitespaces()

		default:
			k.RegisterError(fmt.Sprintf("invalid character `%s`", string(c0.Rune)))
			k.EatChar()
			continue
			// t := tokens.NewTokenAtChar(TInvalid, string(c.Rune), c)

		}
	}
}
