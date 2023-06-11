package eval

import (
	"klc/lang/builtins/constants"
	"klc/lang/builtins/functions"
	"klc/lang/env"
)

func RegisterBuiltins(s *env.Stack) {
	s.Set("echo", functions.Echo)
	s.Set("doc", functions.Doc)
	s.Set("type", functions.Type)
	s.Set("assert", functions.Assert)
	s.Set("exit", functions.Exit)

	s.Set("range", functions.Range)
	s.Set("filter", functions.Filter)
	s.Set("select", functions.Filter)

	s.Set("boolean", functions.Boolean)
	s.Set("even", functions.Even)
	s.Set("odd", functions.Odd)

	s.Set("km", constants.KM)
}
