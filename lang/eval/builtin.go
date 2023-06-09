package eval

import (
	"klc/lang/builtins/constants"
	"klc/lang/builtins/functions"
)

func RegisterBuiltins(s *EnvironmentStack) {
	s.Set("echo", functions.Echo)
	s.Set("doc", functions.Doc)
	s.Set("km", constants.KM)
}
