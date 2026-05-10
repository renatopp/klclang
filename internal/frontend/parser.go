package frontend

import (
	"slices"

	"github.com/renatopp/klclang/internal/core"
	"github.com/renatopp/x/strx"
)

func Parse(script *core.Script, tokens []*core.Token) (core.Node, error) {
	if len(tokens) == 0 {
		return nil, nil
	}

	p := &parser{}
	p.script = script
	p.scanner = NewScanner(tokens, nil)
	p.solver = NewPrattSolver(p.peek)
	p.lastToken = tokens[len(tokens)-1]

	val := p.solver
	val.RegisterPrecedence(core.TokenEqual, 30)
	// val.RegisterPrecedence(core.TokenNotEqual, 30)
	// val.RegisterPrecedence(core.TokenOr, 60)
	// val.RegisterPrecedence(core.TokenAnd, 70)
	// val.RegisterPrecedence(core.TokenLess, 80)
	// val.RegisterPrecedence(core.TokenLessEqual, 80)
	// val.RegisterPrecedence(core.TokenGreater, 80)
	// val.RegisterPrecedence(core.TokenGreaterEqual, 80)
	// val.RegisterPrecedence(core.TokenSpaceShip, 85)
	// val.RegisterPrecedence(core.TokenPlus, 90)
	// val.RegisterPrecedence(core.TokenMinus, 90)
	// val.RegisterPrecedence(core.TokenStar, 100)
	// val.RegisterPrecedence(core.TokenSlash, 100)
	// val.RegisterPrecedence(core.TokenCaret, 110)
	// val.RegisterPrecedence(core.TokenPercent, 120)
	// val.RegisterPrecedence(core.TokenLeftBrace, 129)
	// val.RegisterPrecedence(core.TokenLeftParen, 130)
	// val.RegisterPrecedence(core.TokenLeftBracket, 131)
	// val.RegisterPrecedence(core.TokenBang, 140)
	// val.RegisterPrecedence(core.TokenDot, 150)

	// val.RegisterPrefixFn(core.TokenInt, p.parseLiteral)
	// val.RegisterPrefixFn(core.TokenHex, p.parseLiteral)
	// val.RegisterPrefixFn(core.TokenOct, p.parseLiteral)
	// val.RegisterPrefixFn(core.TokenBin, p.parseLiteral)
	// val.RegisterPrefixFn(core.TokenFloat, p.parseLiteral)
	// val.RegisterPrefixFn(core.TokenString, p.parseLiteral)
	// val.RegisterPrefixFn(core.TokenTrue, p.parseLiteral)
	// val.RegisterPrefixFn(core.TokenFalse, p.parseLiteral)
	// val.RegisterPrefixFn(core.TokenPlus, p.parseUnaryOp)
	// val.RegisterPrefixFn(core.TokenMinus, p.parseUnaryOp)
	// val.RegisterPrefixFn(core.TokenBang, p.parseUnaryOp)
	// val.RegisterInfixFn(core.TokenPlus, p.parseBinaryOp)
	// val.RegisterInfixFn(core.TokenMinus, p.parseBinaryOp)
	// val.RegisterInfixFn(core.TokenStar, p.parseBinaryOp)
	// val.RegisterInfixFn(core.TokenSlash, p.parseBinaryOp)
	// val.RegisterInfixFn(core.TokenPercent, p.parseBinaryOp)
	// val.RegisterInfixFn(core.TokenCaret, p.parseBinaryOp)
	// val.RegisterInfixFn(core.TokenGreater, p.parseBinaryOp)
	// val.RegisterInfixFn(core.TokenGreaterEqual, p.parseBinaryOp)
	// val.RegisterInfixFn(core.TokenLess, p.parseBinaryOp)
	// val.RegisterInfixFn(core.TokenLessEqual, p.parseBinaryOp)
	// val.RegisterInfixFn(core.TokenSpaceShip, p.parseBinaryOp)
	// val.RegisterInfixFn(core.TokenEqual, p.parseBinaryOp)
	// val.RegisterInfixFn(core.TokenNotEqual, p.parseBinaryOp)
	// val.RegisterInfixFn(core.TokenAnd, p.parseBinaryOp)
	// val.RegisterInfixFn(core.TokenOr, p.parseBinaryOp)

	return core.WithRecover(p.parse)
}

type parser struct {
	script     *core.Script
	scanner    *Scanner[*core.Token]
	solver     *PrattSolver // pratt solver for value expressions
	lastToken  *core.Token  // last token in the token stream
	prevToken  *core.Token  // previous token parsed
	currentDoc *core.Token  // current comment token being processed
}

func (p *parser) parse() core.Node {
	// module := p.parseModule()
	// script.Module = p.script
	// return module
	return nil
}

// ---------------------------------------------------------------------------
// HELPER FUNCTIONS
// ---------------------------------------------------------------------------

// peek returns the next significant (non-comment) token without consuming it.
func (p *parser) peek() *core.Token {
	return p.peekAt(0)
}

// peekAt returns the token at the specified lookahead position, skipping comments.
func (p *parser) peekAt(at int) *core.Token {
	for {
		next := p.scanner.PeekAt(at)
		if next == nil {
			return p.lastToken
		}

		if next.IsKind(core.TokenComment) {
			at++
			continue
		}

		return next
	}
}

// eat consumes and returns the next significant (non-comment) token. Comments
// are stored in the parser's currentDoc field.
func (p *parser) eat() *core.Token {
	for {
		next := p.scanner.Eat()
		if next == nil {
			return p.lastToken
		}

		if next.IsKind(core.TokenComment) {
			p.currentDoc = next
			continue
		}

		if next.IsKind(core.TokenSeparator) {
			p.currentDoc = nil
		}
		return next
	}
}

// eatN consumes and returns the next n significant (non-comment) core.Token
func (p *parser) eatN(n int) []*core.Token {
	tokens := make([]*core.Token, 0, n)
	for i := 0; i < n; i++ {
		tokens = append(tokens, p.eat())
	}
	return tokens
}

// skip consumes and returns the next significant (non-comment) token if it
// matches one of the specified kinds.
func (p *parser) skip(kinds ...core.TokenKind) []*core.Token {
	return p.skipN(1, kinds...)
}

// skipN consumes and returns the next n significant (non-comment) tokens if
// they match one of the specified kinds.
func (p *parser) skipN(n int, kinds ...core.TokenKind) []*core.Token {
	tokens := []*core.Token{}
	peek := p.peek()
	for slices.Contains(kinds, peek.Kind) {
		tokens = append(tokens, p.eat())
		peek = p.peek()
	}
	return tokens
}

// expect checks that the next significant (non-comment) token matches one of the
// specified kinds. If it does, it returns the token. If not, it reports an error.
func (p *parser) expect(kinds ...core.TokenKind) *core.Token {
	next := p.peek()
	for !slices.Contains(kinds, next.Kind) {
		kindsStr := []string{}
		for _, k := range kinds {
			kindsStr = append(kindsStr, string(k))
		}
		core.ThrowAt(next.Span, core.ErrorSyntax, "expected %s, got %s",
			strx.Join(kindsStr, "or"),
			next.Kind,
		)
	}
	return next
}

// expectEat checks that the next significant (non-comment) token matches one
// of the specified kinds. If it does, it consumes and returns the token. If
// not, it reports an error.
func (p *parser) expectEat(kinds ...core.TokenKind) *core.Token {
	p.expect(kinds...)
	return p.eat()
}

// expectValue checks that the next significant (non-comment) token can be
// parsed as a value expression. If it can, it parses and returns the value
// expression. If not, it reports an error.
func (p *parser) expectValue(kinds ...core.TokenKind) core.Node {
	value := p.parseValue(kinds...)
	if value == nil {
		core.ThrowAt(p.peek().Span, core.ErrorSyntax, "expected value expression")
	}
	return value
}

// parseValue parses a value expression with optional precedence.
func (p *parser) parseValue(kind ...core.TokenKind) core.Node {
	prec := 0
	if len(kind) > 0 {
		prec = p.solver.PrecedenceOf(kind[0])
	}

	peek := p.peek()
	if peek == p.prevToken {
		core.ThrowAt(peek.Span, core.ErrorInternal, "internal infinite recursion detected")
	}

	p.prevToken = peek
	return p.solver.Solve(prec)
}

// ---------------------------------------------------------------------------
// PARSING FUNCTIONS
// ---------------------------------------------------------------------------

// func (p *parser) parseModule() *nodes.Module {
// 	imports := []*nodes.Import{}
// 	decls := []core.Node{}
// 	for {
// 		peek := p.peek()
// 		switch {
// 		case peek.IsKind(core.TokenEof):
// 			return nodes.NewModule(peek, imports, decls)

// 		case peek.IsKind(core.TokenSeparator):
// 			p.eat() // skip separators

// 		case peek.IsKind(core.TokenImport):
// 			imports = append(imports, p.parseImport())

// 		case peek.IsKind(core.TokenLet):
// 			decls = append(decls, p.parseLetDecl())

// 		default:
// 			core.ThrowAt(peek.Span, core.ErrorSyntax, "unexpected token '%s' at module level", peek.Kind)
// 		}
// 	}
// }

// func (p *parser) parseImport() *nodes.Import {
// 	tok := p.expectEat(core.TokenImport)
// 	path := p.expectEat(core.TokenString)
// 	return nodes.NewImport(tok, nodes.NewString(path, path.Literal))
// }

// func (p *parser) parseLetDecl() *nodes.LetExpr {
// 	letTok := p.expectEat(core.TokenLet)         // consume 'let' token
// 	docs := p.currentDoc                         // capture any preceding comment as documentation
// 	nameTok := p.expectEat(core.TokenValueIdent) // consume identifier token
// 	assignTok := p.expectEat(core.TokenAssign)   // consume '=' token
// 	p.skip(core.TokenSeparator)                  // skip any separators
// 	valueNode := p.expectValue()                 // parse the value expression
// 	letNode := nodes.NewLetExpr(
// 		letTok,
// 		assignTok,
// 		nodes.NewValueIdent(nameTok, nameTok.Literal),
// 		valueNode,
// 	)
// 	letNode.WithDocs(docs)
// 	return letNode
// }

// func (p *parser) parseLiteral() core.Node {
// 	tok := p.eat()
// 	switch tok.Kind {
// 	case core.TokenInt:
// 		val, err := strconv.ParseInt(tok.Literal, 10, 64)
// 		if err != nil {
// 			core.ThrowAt(tok.Span, core.ErrorInternal, "invalid integer literal '%s'", tok.Literal)
// 		}
// 		return nodes.NewInteger(tok, val)

// 	case core.TokenHex:
// 		val, err := strconv.ParseInt(tok.Literal, 16, 64)
// 		if err != nil {
// 			core.ThrowAt(tok.Span, core.ErrorInternal, "invalid hex literal '%s'", tok.Literal)
// 		}
// 		return nodes.NewInteger(tok, val)

// 	case core.TokenOct:
// 		val, err := strconv.ParseInt(tok.Literal, 8, 64)
// 		if err != nil {
// 			core.ThrowAt(tok.Span, core.ErrorInternal, "invalid octal literal '%s'", tok.Literal)
// 		}
// 		return nodes.NewInteger(tok, val)

// 	case core.TokenBin:
// 		val, err := strconv.ParseInt(tok.Literal, 2, 64)
// 		if err != nil {
// 			core.ThrowAt(tok.Span, core.ErrorInternal, "invalid binary literal '%s'", tok.Literal)
// 		}
// 		return nodes.NewInteger(tok, val)

// 	case core.TokenFloat:
// 		val, err := strconv.ParseFloat(tok.Literal, 64)
// 		if err != nil {
// 			core.ThrowAt(tok.Span, core.ErrorInternal, "invalid float literal '%s'", tok.Literal)
// 		}
// 		return nodes.NewFloat(tok, val)

// 	case core.TokenString:
// 		return nodes.NewString(tok, tok.Literal)

// 	case core.TokenTrue:
// 		return nodes.NewBool(tok, true)

// 	case core.TokenFalse:
// 		return nodes.NewBool(tok, false)
// 	default:
// 		core.ThrowAt(tok.Span, core.ErrorSyntax, "unexpected literal token '%s'", tok.Kind)
// 		return nil
// 	}
// }

// func (p *parser) parseUnaryOp() core.Node {
// 	opTok := p.eat()
// 	operand := p.solver.Prefix()
// 	if operand == nil {
// 		core.ThrowAt(opTok.Span, core.ErrorSyntax, "expected expression after unary operator '%s'", opTok.Literal)
// 	}
// 	return nodes.NewUnaryOp(opTok, opTok.Literal, operand)
// }

// func (p *parser) parseBinaryOp(left core.Node) core.Node {
// 	opTok := p.eat()
// 	right := p.expectValue()
// 	return nodes.NewBinaryOp(opTok, left, opTok.Literal, right)
// }
