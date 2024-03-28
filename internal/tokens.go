package internal

import "github.com/renatopp/langtools"

const (
	TComment    langtools.TokenType = "comment"
	TIdentifier langtools.TokenType = "identifier"
	TNumber     langtools.TokenType = "number"
	TAssign     langtools.TokenType = "assign"
	TOperand    langtools.TokenType = "operand"
	TKeyword    langtools.TokenType = "keyword"
	TLParen     langtools.TokenType = "lparen"
	TRParen     langtools.TokenType = "rparen"
	TComma      langtools.TokenType = "comma"
	TEol        langtools.TokenType = "eol"
)
