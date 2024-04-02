package internal

import (
	"github.com/renatopp/klclang/internal/runtime"
	"github.com/renatopp/langtools/asts"
)

func NewEvaluator() *runtime.Runtime {
	return runtime.NewRuntime()
}

func Eval(node asts.Node) runtime.Object {
	return NewEvaluator().Eval(node)
}

func Run(code []byte) runtime.Object {
	lexer := NewLexer(code)
	parser := NewParser(lexer)
	node := parser.Parse()
	return Eval(node)
}
