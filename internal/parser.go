package internal

import (
	"strconv"
	"strings"

	"github.com/renatopp/klclang/internal/ast"
	"github.com/renatopp/langtools/asts"
	"github.com/renatopp/langtools/parsers"
	"github.com/renatopp/langtools/tokens"
)

type KlcParser struct {
	*parsers.PrattParser

	previousComment    string
	previousNodeInLine asts.Node
}

func NewParser(lexer *KlcLexer) *KlcParser {
	k := &KlcParser{}
	k.PrattParser = parsers.NewPrattParser(lexer)
	k.PrattParser.IsEndOfExpr = k.isEndOfExpr
	k.PrattParser.PrecedenceFn = k.precedence

	k.RegisterPrefixFn(TNumber, k.prefixNumber)
	k.RegisterPrefixFn(TIdentifier, k.prefixIdentifier)
	k.RegisterPrefixFn(TLParen, k.prefixParen)
	k.RegisterPrefixFn(TOperator, k.prefixOperator)
	k.RegisterInfixFn(TAssign, k.infixAssignment)
	k.RegisterInfixFn(TOperator, k.infixOperator)
	k.RegisterInfixFn(TKeyword, k.infixKeyword)
	k.RegisterInfixFn(TLParen, k.infixParen)

	return k
}

func (k *KlcParser) isEndOfExpr(t tokens.Token) bool {
	return t.Type == TEoe || t.Type == TEof
}

const PRECEDENCE_UNARY = 36

func (k *KlcParser) precedence(t tokens.Token) int {
	switch {
	case t.IsType(TComma):
		return 10

	case t.IsType(TAssign):
		return 15

	case t.IsType(TOperator) && t.IsOneOfLiterals("or"):
		return 20
	case t.IsType(TOperator) && t.IsOneOfLiterals("and"):
		return 21
	case t.IsType(TOperator) && t.IsOneOfLiterals("!"):
		return 22

	case t.IsType(TOperator) && t.IsOneOfLiterals("==", "!="):
		return 25
	case t.IsType(TOperator) && t.IsOneOfLiterals(">", "<", ">=", "<="):
		return 26

	case t.IsType(TKeyword) && t.IsOneOfLiterals("to"):
		return 29
	case t.IsType(TOperator) && t.IsOneOfLiterals("+", "-"):
		return 30
	case t.IsType(TOperator) && t.IsOneOfLiterals("*", "/"):
		return 31
	case t.IsType(TOperator) && t.IsOneOfLiterals("%"):
		return 32
	case t.IsType(TOperator) && t.IsOneOfLiterals("^"):
		return 34
	// UNARY

	case t.IsType(TLParen):
		return 40

	default:
		return 0
	}
}

func (k *KlcParser) Parse() asts.Node {
	first := k.Lexer.PeekToken()
	statements := []asts.Node{}

	cur := k.Lexer.PeekToken()
	for !cur.IsType(TEof) {
		switch {
		case cur.IsType(TEoe):
			if cur.IsLiteral("\n") {
				k.previousNodeInLine = nil
			}
			k.Lexer.EatToken()

		case cur.IsType(TComment):
			k.appendComment(cur.Literal)

			if k.previousNodeInLine != nil {
				switch node := k.previousNodeInLine.(type) {
				case *ast.Assignment:
					node.Documentation = k.previousComment
				case *ast.FunctionDef:
					node.Documentation = k.previousComment
				}
			}

			k.Lexer.EatToken()

		default:
			node := k.ParseExpression(0)
			if node == nil {
				k.RegisterError("invalid expression")
				return nil
			}

			k.previousComment = ""
			k.previousNodeInLine = node
			statements = append(statements, node)
		}
		cur = k.Lexer.PeekToken()
	}

	return &ast.Block{
		Token:      first,
		Statements: statements,
	}
}

func (k *KlcParser) parseExpressionList() []asts.Node {
	list := []asts.Node{}

	for {
		expr := k.ParseExpression(0)
		if expr == nil {
			break
		}

		list = append(list, expr)

		cur := k.Lexer.PeekToken()
		if cur.IsType(TComma) {
			k.Lexer.EatToken()
			continue
		}
	}

	return list
}

func (k *KlcParser) prefixNumber() asts.Node {
	t := k.Lexer.EatToken()
	v, err := strconv.ParseFloat(t.Literal, 64)
	if err != nil {
		k.RegisterErrorWithToken(err.Error(), t)
	}

	return &ast.Number{
		Token: t,
		Value: v,
	}
}

func (k *KlcParser) prefixIdentifier() asts.Node {
	t := k.Lexer.EatToken()
	return &ast.Identifier{
		Token: t,
		Name:  t.Literal,
	}
}

func (k *KlcParser) prefixParen() asts.Node {
	t := k.Lexer.EatToken()
	expr := k.ParseExpression(0)

	if expr == nil {
		k.RegisterErrorWithToken("expected expression", t)
		return nil
	}

	if !k.ExpectType(TRParen) {
		return nil
	}

	k.Lexer.EatToken()
	return expr
}

func (k *KlcParser) prefixOperator() asts.Node {
	cur := k.Lexer.EatToken()

	if !cur.IsOneOfLiterals("-", "+", "!") {
		k.RegisterErrorWithToken("expected unary operator", cur)
		return nil
	}

	return &ast.UnaryOperator{
		Token:      cur,
		Operator:   cur.Literal,
		Expression: k.ParseExpression(PRECEDENCE_UNARY),
	}
}

func (k *KlcParser) infixAssignment(left asts.Node) asts.Node {
	cur := k.Lexer.EatToken()
	right := k.ParseExpression(k.PrecedenceFn(cur))

	if isNodeAn[*ast.Identifier](left) {
		return &ast.Assignment{
			Token:         cur,
			Operator:      cur.Literal,
			Documentation: k.previousComment,
			Identifier:    left.(*ast.Identifier),
			Expression:    right,
		}
	}

	if isNodeAn[*ast.FunctionCall](left) {
		left := left.(*ast.FunctionCall)

		for _, arg := range left.Arguments {
			if !isNodeAn[*ast.Identifier](arg) && !isNodeAn[*ast.Number](arg) {
				k.RegisterErrorWithToken("function definition only accept identifiers and numbers as parameters", arg.GetToken())
				return nil
			}
		}

		return &ast.FunctionDef{
			Token:         left.Token,
			Name:          left.Target.Name,
			Documentation: k.previousComment,
			Params:        left.Arguments,
			Body:          right,
		}
	}

	k.RegisterErrorWithToken("expected identifier or function definition at left", left.GetToken())
	return nil
}

func (k *KlcParser) infixOperator(left asts.Node) asts.Node {
	cur := k.Lexer.EatToken()
	right := k.ParseExpression(k.PrecedenceFn(cur))

	if right == nil {
		k.RegisterErrorWithToken("expected expression", cur)
		return nil
	}

	return &ast.BinaryOperator{
		Token:    cur,
		Operator: cur.Literal,
		Left:     left,
		Right:    right,
	}
}

func (k *KlcParser) infixKeyword(left asts.Node) asts.Node {
	t := k.Lexer.EatToken()
	right := k.ParseExpression(k.PrecedenceFn(t))

	if right == nil {
		k.RegisterErrorWithToken("expected expression", t)
		return nil
	}

	return &ast.BinaryOperator{
		Token:    t,
		Operator: "/",
		Left:     left,
		Right:    right,
	}
}

func (k *KlcParser) infixParen(left asts.Node) asts.Node {
	t := k.Lexer.EatToken()

	if !isNodeAn[*ast.Identifier](left) {
		k.RegisterErrorWithToken("expected identifier at left", left.GetToken())
		return nil
	}

	right := k.parseExpressionList()
	if !k.ExpectType(TRParen) {
		return nil
	}

	k.Lexer.EatToken()
	return &ast.FunctionCall{
		Token:     t,
		Target:    left.(*ast.Identifier),
		Arguments: right,
	}
}

func (k *KlcParser) appendComment(msg string) {
	msg = strings.TrimSpace(msg)
	switch {
	case k.previousComment == "":
		k.previousComment = msg

	case msg == "":
		k.previousComment = k.previousComment + "\n"

	case strings.HasSuffix(k.previousComment, "\n"):
		k.previousComment = k.previousComment + msg

	default:
		k.previousComment = k.previousComment + " " + msg
	}
}

func isNodeAn[T asts.Node](node asts.Node) bool {
	_, ok := node.(T)
	return ok
}
