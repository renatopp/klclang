package runtime

import "github.com/renatopp/langtools/asts"

type RuntimeError struct {
	node asts.Node
	msg  string
}

func (e RuntimeError) Error() string {
	return e.msg
}

func (e RuntimeError) At() (int, int) {
	token := e.node.GetToken()
	return token.Line, token.Column
}
