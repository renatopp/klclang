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

func Run(code []byte) (runtime.Object, error) {
	lexer := NewLexer(code)
	if lexer.HasErrors() {
		return nil, ConvertLexerErrors(code, lexer.Errors())
	}

	parser := NewParser(lexer)
	node := parser.Parse()
	if parser.HasErrors() {
		return nil, ConvertParserErrors(code, parser.Errors())
	}

	return Eval(node), nil
}
