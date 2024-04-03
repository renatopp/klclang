package internal

import (
	"github.com/renatopp/klclang/internal/runtime"
)

func NewEvaluator() *runtime.Runtime {
	return runtime.NewRuntime()
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

	runtime := NewEvaluator()
	obj := runtime.Eval(node)

	if runtime.HasErrors() {
		return nil, ConvertRuntimeErrors(code, runtime.Errors())
	}
	return obj, nil
}
