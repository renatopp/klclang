package internal

import (
	"github.com/renatopp/langtools/tokens"
)

const (
	TInvalid    tokens.TokenType = tokens.INVALID
	TUnknown    tokens.TokenType = tokens.UNKNOWN
	TEof        tokens.TokenType = tokens.EOF
	TComment    tokens.TokenType = "comment"
	TIdentifier tokens.TokenType = "identifier"
	TNumber     tokens.TokenType = "number"
	TAssign     tokens.TokenType = "assign"
	TOperator   tokens.TokenType = "operator"
	TKeyword    tokens.TokenType = "keyword"
	TLParen     tokens.TokenType = "lparen"
	TRParen     tokens.TokenType = "rparen"
	TComma      tokens.TokenType = "comma"
	TEoe        tokens.TokenType = "eoe"
)
