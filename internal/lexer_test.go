package internal_test

import (
	"testing"

	"github.com/renatopp/klclang/internal"
	"github.com/renatopp/langtools/tokens"
	"github.com/stretchr/testify/assert"
)

func TestLexer(t *testing.T) {

	var input = `
x = 2 or 3
y = -3.3e-2
fn(x, y) = x + y;
fn(2, 3)
2m to s
help(m)
example == 3 -- this is a comment
`

	var expected = []tokens.Token{
		{Type: internal.TEoe, Literal: "\n"},
		{Type: internal.TIdentifier, Literal: "x"},
		{Type: internal.TAssign, Literal: "="},
		{Type: internal.TNumber, Literal: "2"},
		{Type: internal.TOperator, Literal: "or"},
		{Type: internal.TNumber, Literal: "3"},
		{Type: internal.TEoe, Literal: "\n"},
		{Type: internal.TIdentifier, Literal: "y"},
		{Type: internal.TAssign, Literal: "="},
		{Type: internal.TOperator, Literal: "-"},
		{Type: internal.TNumber, Literal: "3.3e-2"},
		{Type: internal.TEoe, Literal: "\n"},
		{Type: internal.TIdentifier, Literal: "fn"},
		{Type: internal.TLParen, Literal: "("},
		{Type: internal.TIdentifier, Literal: "x"},
		{Type: internal.TComma, Literal: ","},
		{Type: internal.TIdentifier, Literal: "y"},
		{Type: internal.TRParen, Literal: ")"},
		{Type: internal.TAssign, Literal: "="},
		{Type: internal.TIdentifier, Literal: "x"},
		{Type: internal.TOperator, Literal: "+"},
		{Type: internal.TIdentifier, Literal: "y"},
		{Type: internal.TEoe, Literal: ";"},
		{Type: internal.TEoe, Literal: "\n"},
		{Type: internal.TIdentifier, Literal: "fn"},
		{Type: internal.TLParen, Literal: "("},
		{Type: internal.TNumber, Literal: "2"},
		{Type: internal.TComma, Literal: ","},
		{Type: internal.TNumber, Literal: "3"},
		{Type: internal.TRParen, Literal: ")"},
		{Type: internal.TEoe, Literal: "\n"},
		{Type: internal.TNumber, Literal: "2"},
		{Type: internal.TOperator, Literal: "*"},
		{Type: internal.TIdentifier, Literal: "m"},
		{Type: internal.TKeyword, Literal: "to"},
		{Type: internal.TIdentifier, Literal: "s"},
		{Type: internal.TEoe, Literal: "\n"},
		{Type: internal.TIdentifier, Literal: "help"},
		{Type: internal.TLParen, Literal: "("},
		{Type: internal.TIdentifier, Literal: "m"},
		{Type: internal.TRParen, Literal: ")"},
		{Type: internal.TEoe, Literal: "\n"},
		{Type: internal.TIdentifier, Literal: "example"},
		{Type: internal.TOperator, Literal: "=="},
		{Type: internal.TNumber, Literal: "3"},
		{Type: internal.TComment, Literal: " this is a comment"},
	}

	lexer := internal.NewLexer([]byte(input))
	results := lexer.All()

	for i := range expected {
		assert.Equal(t, expected[i].Type, results[i].Type)
		assert.Equal(t, expected[i].Literal, results[i].Literal)
	}
}

func TestMultShortcut(t *testing.T) {
	var input = `2a`

	var expected = []tokens.Token{
		{Type: internal.TNumber, Literal: "2"},
		{Type: internal.TOperator, Literal: "*"},
		{Type: internal.TIdentifier, Literal: "a"},
	}

	lexer := internal.NewLexer([]byte(input))
	results := lexer.All()

	for i := range expected {
		assert.Equal(t, expected[i].Type, results[i].Type)
		assert.Equal(t, expected[i].Literal, results[i].Literal)
	}
}

func TestMultShortcut2(t *testing.T) {
	var input = `2(a)`

	var expected = []tokens.Token{
		{Type: internal.TNumber, Literal: "2"},
		{Type: internal.TOperator, Literal: "*"},
		{Type: internal.TLParen, Literal: "("},
		{Type: internal.TIdentifier, Literal: "a"},
		{Type: internal.TRParen, Literal: ")"},
	}

	lexer := internal.NewLexer([]byte(input))
	results := lexer.All()

	for i := range expected {
		assert.Equal(t, expected[i].Type, results[i].Type)
		assert.Equal(t, expected[i].Literal, results[i].Literal)
	}
}
