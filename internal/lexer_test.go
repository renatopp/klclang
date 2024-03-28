package internal_test

import (
	"testing"

	"github.com/renatopp/klclang/internal"
	"github.com/renatopp/langtools"
	"github.com/stretchr/testify/assert"
)

var input = `
x = 2
y = -3.3e-2
fn(x, y) = x + y;
fn(2, 3)
2m to s
help(m)
example == 3 -- this is a comment
`

var expected = []langtools.Token{
	{Type: internal.TEol, Literal: "\n"},
	{Type: internal.TIdentifier, Literal: "x"},
	{Type: internal.TAssign, Literal: "="},
	{Type: internal.TNumber, Literal: "2"},
	{Type: internal.TEol, Literal: "\n"},
	{Type: internal.TIdentifier, Literal: "y"},
	{Type: internal.TAssign, Literal: "="},
	{Type: internal.TOperand, Literal: "-"},
	{Type: internal.TNumber, Literal: "3.3e-2"},
	{Type: internal.TEol, Literal: "\n"},
	{Type: internal.TIdentifier, Literal: "fn"},
	{Type: internal.TLParen, Literal: "("},
	{Type: internal.TIdentifier, Literal: "x"},
	{Type: internal.TComma, Literal: ","},
	{Type: internal.TIdentifier, Literal: "y"},
	{Type: internal.TRParen, Literal: ")"},
	{Type: internal.TAssign, Literal: "="},
	{Type: internal.TIdentifier, Literal: "x"},
	{Type: internal.TOperand, Literal: "+"},
	{Type: internal.TIdentifier, Literal: "y"},
	{Type: internal.TEol, Literal: ";"},
	{Type: internal.TEol, Literal: "\n"},
	{Type: internal.TIdentifier, Literal: "fn"},
	{Type: internal.TLParen, Literal: "("},
	{Type: internal.TNumber, Literal: "2"},
	{Type: internal.TComma, Literal: ","},
	{Type: internal.TNumber, Literal: "3"},
	{Type: internal.TRParen, Literal: ")"},
	{Type: internal.TEol, Literal: "\n"},
	{Type: internal.TNumber, Literal: "2"},
	{Type: internal.TIdentifier, Literal: "m"},
	{Type: internal.TKeyword, Literal: "to"},
	{Type: internal.TIdentifier, Literal: "s"},
	{Type: internal.TEol, Literal: "\n"},
	{Type: internal.TIdentifier, Literal: "help"},
	{Type: internal.TLParen, Literal: "("},
	{Type: internal.TIdentifier, Literal: "m"},
	{Type: internal.TRParen, Literal: ")"},
	{Type: internal.TEol, Literal: "\n"},
	{Type: internal.TIdentifier, Literal: "example"},
	{Type: internal.TOperand, Literal: "=="},
	{Type: internal.TNumber, Literal: "3"},
	{Type: internal.TComment, Literal: "this is a comment"},
}

func TestLexer(t *testing.T) {
	lexer := internal.NewLexer([]byte(input))
	results := []langtools.Token{}

	for {
		t, eof := lexer.Next()
		if eof {
			break
		}

		results = append(results, t)
	}

	for i, r := range results {
		assert.Equal(t, expected[i].Type, r.Type)
		assert.Equal(t, expected[i].Literal, r.Literal)
	}
}
