package parser

import (
	"klc/lang/ast"
	"klc/lang/lexer"
	"klc/lang/parser/order"
	"klc/lang/token"
	"strconv"
)

var priorities = map[string]int{
	"+":  order.Addition,
	"-":  order.Subtraction,
	"*":  order.Multiplication,
	"/":  order.Division,
	"%":  order.Mod,
	"**": order.Exponentiation,
	"//": order.Division,
	"==": order.Comparison,
	"!=": order.Comparison,
	">":  order.Comparison,
	"<":  order.Comparison,
	">=": order.Comparison,
	"<=": order.Comparison,
	"!":  order.Not,
	"&&": order.And,
	"||": order.Or,
	"++": order.Concat,
}

func priorityOf(t *token.Token) int {
	switch t.Type {
	case token.Operator:
		op := t.Literal
		v, ok := priorities[op]
		if !ok {
			return order.Lowest
		}
		return v
	case token.Assignment:
		return order.Assign
	case token.Lparen:
		return order.Calls
	case token.Lbracket:
		return order.Indexing
	case token.Dot:
		return order.Chain
	default:
		return order.Lowest
	}
}

type prefixFn func() ast.Node
type infixFn func(ast.Node) ast.Node

type Parser struct {
	lexer    *lexer.Lexer
	root     *ast.Program
	prefixes map[token.Type]prefixFn
	infixes  map[token.Type]infixFn
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{
		lexer:   l,
		root:    &ast.Program{},
		infixes: make(map[token.Type]infixFn),
	}

	p.prefixes = map[token.Type]prefixFn{
		token.Identifier: p.parseIdentifier,
		token.String:     p.parseString,
		token.Number:     p.parseNumber,
		token.Operator:   p.parseUnaryOperator,
		token.Lparen:     p.parseGrouping,
		token.Keyword:    p.parseKeyword,
		token.Lbracket:   p.parseList,
	}

	p.infixes = map[token.Type]infixFn{
		token.Operator:   p.parseBinaryOperator,
		token.Lparen:     p.parseFunctionCall,
		token.Lbracket:   p.parseIndexing,
		token.Assignment: p.parseAssignment,
		token.Dot:        p.parseChain,
	}

	return p
}

func (p *Parser) Parse() *ast.Program {
	prg := &ast.Program{}
	prg.Root = p.parseBlock()
	return prg
}

func (p *Parser) expect(t token.Type) {
	if !p.lexer.Current().Is(t) {
		panic("expected " + string(t) + ", got " + p.lexer.Current().ToString())
	}
}

func (p *Parser) expectNext(t token.Type) {
	if !p.lexer.Peek().Is(t) {
		panic("expected " + string(t) + ", got " + p.lexer.Peek().ToString())
	}
}

// ----------------------------------------------------------------------------
// PARSING
// ----------------------------------------------------------------------------
func (p *Parser) parseBlock() ast.Node {
	block := &ast.Block{}

	t := p.lexer.Current()
	if t.Is(token.Lbrace) {
		t = p.lexer.Next()
	}

	for !isEndOfBlock(t) {
		s := p.parseStatement()
		if s != nil {
			block.Statements = append(block.Statements, s)
		}
		t = p.lexer.Current()
	}

	if t.Is(token.Rbrace) {
		p.lexer.Next()
	}

	return block
}

func (p *Parser) parseStatement() ast.Node {
	t := p.lexer.Current()
	for t.Is(token.Newline) {
		t = p.lexer.Next()
	}

	if t.Is(token.Eof) {
		return nil
	} else if t.Is(token.Lbrace) {
		return p.parseBlock()
	} else if t.Is(token.Question) {
		return p.parseIfReturn()
	}

	e := p.parseExpression(order.Lowest)

	if e == nil {
		panic("invalid token " + t.ToString())
	}

	if !isEndOfStatement(p.lexer.Current()) {
		panic("expected end of statement, got " + p.lexer.Current().ToString())
	}

	return e
}

func (p *Parser) parseIfReturn() ast.Node {
	p.lexer.Next()
	cond := p.parseExpression(order.Lowest)
	if cond == nil {
		panic("invalid expression " + p.lexer.Current().ToString())
	}

	t := p.lexer.Current()
	if !t.Is(token.Assignment) {
		return &ast.IfReturn{
			Condition: ast.TrueCondition(),
			Return:    cond,
		}
	}

	if t.Literal != "=" {
		panic("expected assignment, got " + t.ToString())
	}

	p.lexer.Next()
	ret := p.parseExpression(order.Lowest)
	if cond == nil {
		panic("invalid expression " + p.lexer.Current().ToString())
	}

	return &ast.IfReturn{
		Condition: cond,
		Return:    ret,
	}
}

func (p *Parser) parseExpression(priority int) ast.Node {
	t := p.lexer.Current()

	for t.Is(token.Newline) {
		t = p.lexer.Next()
	}

	prefix := p.prefixes[t.Type]
	if prefix == nil {
		return nil
	}

	root := prefix()

	nt := p.lexer.Current()
	for nt.Is(token.Newline) {
		nt = p.lexer.Next()
	}
	for !isEndOfExpression(nt) && priorityOf(nt) >= priority {
		infix := p.infixes[nt.Type]
		if infix == nil {
			return root
		}

		root = infix(root)
		nt = p.lexer.Current()
	}

	return root
}

func (p *Parser) parseIdentifier() ast.Node {
	t := p.lexer.Current()
	p.lexer.Next()
	return &ast.Identifier{
		Token: t,
		Value: t.Literal,
	}
}

func (p *Parser) parseNumber() ast.Node {
	t := p.lexer.Current()
	v, e := strconv.ParseFloat(t.Literal, 64)

	if e != nil {
		panic(e)
	}

	p.lexer.Next()
	return &ast.Number{
		Token: t,
		Value: v,
	}
}

func (p *Parser) parseString() ast.Node {
	t := p.lexer.Current()
	p.lexer.Next()
	return &ast.String{
		Token: t,
		Value: t.Literal,
	}
}

func (p *Parser) parseUnaryOperator() ast.Node {
	t := p.lexer.Current()
	p.lexer.Next()

	if !isUnary(t) {
		panic("invalid unary operator " + t.ToString())
	}

	node := &ast.UnaryOperation{
		Token: t,
	}

	node.Right = p.parseExpression(order.Unary)

	return node
}

func (p *Parser) parseGrouping() ast.Node {
	p.lexer.Next()
	node := p.parseExpression(order.Lowest)
	p.expect(token.Rparen)
	p.lexer.Next()
	return node
}

func (p *Parser) parseExpressionList() []ast.Node {
	t := p.lexer.Current()
	args := make([]ast.Node, 0)

	for {
		for t.Is(token.Newline) {
			t = p.lexer.Next()
		}

		arg := p.parseExpression(order.Lowest)
		if arg == nil {
			break
		}

		args = append(args, arg)

		t = p.lexer.Current()
		for t.Is(token.Comma) {
			t = p.lexer.Next()
		}
	}

	return args
}

func (p *Parser) parseAssignment(left ast.Node) ast.Node {
	if !isValidLeftAssignment(left) {
		panic("invalid left assignment " + left.String())
	}

	id := left

	assign := p.lexer.Current()

	p.lexer.Next()
	exp := p.parseExpression(order.Lowest)

	if exp == nil {
		panic("invalid expression " + p.lexer.Current().ToString())
	}

	return &ast.Assignment{
		Identifier: id,
		Token:      assign,
		Expression: exp,
	}
}

func (p *Parser) parseBinaryOperator(left ast.Node) ast.Node {
	t := p.lexer.Current()
	pr := priorityOf(t)

	node := &ast.BinaryOperation{
		Token: t,
	}

	p.lexer.Next()

	node.Left = left
	node.Right = p.parseExpression(pr)

	return node
}

func (p *Parser) parseFunctionCall(left ast.Node) ast.Node {
	p.lexer.Next()

	node := &ast.FunctionCall{
		Function:  left,
		Arguments: p.parseExpressionList(),
	}

	p.expect(token.Rparen)
	p.lexer.Next()

	return node
}

func (p *Parser) parseIndexing(left ast.Node) ast.Node {
	p.lexer.Next()

	a := p.parseExpression(order.Lowest)

	t := p.lexer.Current()
	if !t.Is(token.Colon) {
		p.expect(token.Rbracket)
		p.lexer.Next()
		return &ast.Index{
			Target: left,
			Value:  a,
		}
	}

	p.lexer.Next()
	b := p.parseExpression(order.Lowest)

	if a == nil {
		a = ast.Zero()
	}
	if b == nil {
		b = ast.Zero()
	}

	p.expect(token.Rbracket)
	p.lexer.Next()
	return &ast.Slice{
		Target: left,
		From:   a,
		To:     b,
	}
}

func (p *Parser) parseKeyword() ast.Node {
	t := p.lexer.Current()

	switch t.Literal {
	case "where":
		return p.parseWhereFunction()
	case "fn":
		return p.parseFnFunction()
	default:
		panic("invalid keyword " + t.ToString())
	}
}

func (p *Parser) parseList() ast.Node {
	p.expect(token.Lbracket)
	p.lexer.Next()

	v := p.parseExpressionList()

	p.expect(token.Rbracket)
	p.lexer.Next()
	return &ast.List{
		Values: v,
	}
}

func (p *Parser) parseWhereFunction() ast.Node {
	p.lexer.Next()

	block := p.parseExpression(order.Lowest)
	if block == nil {
		panic("invalid expression " + p.lexer.Current().ToString())
	}

	return &ast.FunctionDef{
		Params: []ast.Node{
			ast.Id("x"),
			ast.Id("y"),
			ast.Id("z"),
		},
		Block: block,
	}
}

func (p *Parser) parseFnFunction() ast.Node {
	p.lexer.Next()

	params := p.parseParameters()

	p.expect(token.Lbrace)
	block := p.parseBlock()

	return &ast.FunctionDef{
		Params: params,
		Block:  block,
	}
}

func (p *Parser) parseParameters() []ast.Node {
	if p.lexer.Current().Is(token.Lparen) {
		p.lexer.Next()
	}

	params := make([]ast.Node, 0)

	t := p.lexer.Current()
	for {
		for t.Is(token.Newline) {
			t = p.lexer.Next()
		}

		var arg ast.Node
		if t.Is(token.Spread) {
			p.lexer.Next()
			p.expect(token.Identifier)
			arg = &ast.SpreadArg{
				Identifier: p.parseIdentifier(),
			}

		} else if t.Is(token.Identifier) {
			arg = p.parseIdentifier()

		} else {
			break
		}

		if arg == nil {
			break
		}

		t = p.lexer.Current()
		if t.Is(token.Assignment) && t.Literal == "=" {
			arg = p.parseDefaultArg(arg)
		}

		params = append(params, arg)

		t = p.lexer.Current()
		for t.Is(token.Comma) {
			t = p.lexer.Next()
		}
	}

	if p.lexer.Current().Is(token.Rparen) {
		p.lexer.Next()
	}

	return params
}

func (p *Parser) parseDefaultArg(id ast.Node) ast.Node {
	p.lexer.Next()
	val := p.parseLiteral()
	return &ast.DefaultArg{
		Identifier: id,
		Value:      val,
	}
}

func (p *Parser) parseLiteral() ast.Node {
	t := p.lexer.Current()

	switch t.Type {
	case token.Number:
		return p.parseNumber()
	case token.String:
		return p.parseString()
	default:
		panic("invalid literal " + t.ToString())
	}
}

func (p *Parser) parseChain(left ast.Node) ast.Node {
	p.lexer.Next()

	p.expect(token.Identifier)
	right := p.parseExpression(order.Chain)
	if right == nil {
		panic("invalid expression " + p.lexer.Current().ToString())
	}

	return &ast.Chain{
		Left:  left,
		Right: right,
	}
}

// ----------------------------------------------------------------------------
// HELPERS
// ----------------------------------------------------------------------------
func isEndOfExpression(t *token.Token) bool {
	return t.Is(token.Semicolon) // t.Is(token.Newline) ||
}

func isEndOfBlock(t *token.Token) bool {
	return t.Is(token.Rbrace) || t.Is(token.Eof)
}

func isEndOfStatement(t *token.Token) bool {
	return t.Is(token.Semicolon) || t.Is(token.Eof) || t.Is(token.Newline) || t.Is(token.Rbrace)
}

func isUnary(t *token.Token) bool {
	return t.Is(token.Operator) && (t.Literal == "+" || t.Literal == "-" || t.Literal == "!")
}

func isValidLeftAssignment(n ast.Node) bool {
	switch n.(type) {
	case *ast.Identifier, *ast.Index, *ast.Slice:
		return true
	}
	return false
}

func validateLeftChain(n ast.Node) bool {
	switch n.(type) {
	case *ast.FunctionCall, *ast.Index, *ast.Slice, *ast.Identifier, *ast.Number, *ast.String:
		return true
	}
	return false
}

func validateRightChain(n ast.Node) bool {
	switch n.(type) {
	case *ast.FunctionCall:
		return true
	}
	return false
}
