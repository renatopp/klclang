package core

import "slices"

type TokenKind string
type ErrorKind string

// Represents a klc script file.
type Script struct {
	CanonicalName string // "Canonical" id.
	Path          string // Absolute path to the script file.
	Content       []byte // Source code of the script file
}

// Span represents an interval of text in the source code.
type Span struct {
	Script     *Script // Script the span belongs to
	From       int     // Starting rune offset (0-based)
	To         int     // Ending rune offset (0-based, exclusive)
	FromLine   int     // Starting line number (1-based)
	FromColumn int     // Starting column number (1-based)
	ToLine     int     // Ending line number (1-based)
	ToColumn   int     // Ending column number (1-based)
}

// Token represents a lexical token in the source code.
type Token struct {
	Span    *Span     // Span of the token in the source code
	Kind    TokenKind // Kind of the token
	Literal string    // Literal value of the token
}

func NewToken(span *Span, kind TokenKind, literal string) *Token {
	return &Token{
		Span:    span,
		Kind:    kind,
		Literal: literal,
	}
}
func (t *Token) IsKind(k ...TokenKind) bool {
	if t == nil {
		return false
	}
	return slices.Contains(k, t.Kind)
}
func (t *Token) IsLiteral(l ...string) bool {
	if t == nil {
		return false
	}
	return slices.Contains(l, t.Literal)
}
func (t *Token) Absorb(other *Token) {
	if other == nil {
		return
	}
	t.Literal += other.Literal
	t.Span.To = other.Span.To
	t.Span.ToLine = other.Span.ToLine
	t.Span.ToColumn = other.Span.ToColumn
}

// Node represents a node in the abstract syntax tree (AST) of the source code.
type Node interface {
	GetToken() *Token // Token associated with the node
	GetSpan() *Span   // Span of the token in the source code
	GetDocs() *Token  // Documentation comment token associated with the node
}
