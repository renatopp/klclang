package tokenizer

import (
	"klc/pkg/lang/tokens"
	"regexp"
)

type rule struct {
	Name    tokens.Type
	Regex   regexp.Regexp
	Ignore  bool
	Matcher func(*Token) *Token
}
