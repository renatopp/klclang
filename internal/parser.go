package internal

import (
	"github.com/renatopp/klclang/internal/ast"
	"github.com/renatopp/langtools"
)

// type prefixFn func

type Parser struct {
	lexer      *langtools.Lexer
	prefixFns  map[langtools.TokenType]func() ast.INode
	infixFns   map[langtools.TokenType]func(ast.INode) ast.INode
	postfixFns map[langtools.TokenType]func(ast.INode) ast.INode
	errors     []ParserError
}
