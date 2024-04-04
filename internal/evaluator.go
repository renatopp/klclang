package internal

import (
	"github.com/renatopp/klclang/internal/runtime"
)

func NewEvaluator() *runtime.Runtime {
	return runtime.NewRuntime()
}

func Run(code []byte) (runtime.Object, error) {
	runtime := NewEvaluator()
	return RunInRuntime(code, runtime)
}

func RunInRuntime(code []byte, runtime *runtime.Runtime) (runtime.Object, error) {
	lexer := NewLexer(code)
	if lexer.HasErrors() {
		return nil, ConvertLexerErrors(code, lexer.Errors())
	}

	parser := NewParser(lexer)
	node := parser.Parse()
	if parser.HasErrors() {
		return nil, ConvertParserErrors(code, parser.Errors())
	}

	obj := runtime.Eval(node)
	if runtime.HasErrors() {
		errors := runtime.Errors()
		runtime.ClearErrors()
		return nil, ConvertRuntimeErrors(code, errors)
	}
	return obj, nil
}
