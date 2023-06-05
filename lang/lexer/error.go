package lexer

import "fmt"

type LexerError struct {
	Line   int
	Column int
	Reason string
}

func (e LexerError) Error() string {
	return fmt.Sprintf("Lexer error at %d:%d: %s", e.Line, e.Column, e.Reason)
}
