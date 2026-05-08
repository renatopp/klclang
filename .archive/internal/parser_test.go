package internal_test

import (
	"testing"

	"github.com/renatopp/klclang/internal"
)

func TestParser(t *testing.T) {
	input := []byte(`a + 2`)

	lexer := internal.NewLexer(input)
	parser := internal.NewParser(lexer)

	node := parser.Parse()
	if node == nil {
		t.Fatal("node is nil")
	}

	// node.Traverse(0, func(i int, n asts.Node) {
	// 	fmt.Printf("%s%s\n", strings.Repeat("  ", i), n.String())
	// })
}
