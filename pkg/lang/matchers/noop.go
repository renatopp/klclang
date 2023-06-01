package tokenizer

import (
	"klc/pkg/lang/tokenizer"
)

func Noop(t *tokenizer.Token) *tokenizer.Token {
	return t
}
