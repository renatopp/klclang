package tokenizer

import (
	"klc/pkg/lang/tokens"
	"regexp"
)

type Tokenizer struct {
	rules  []rule
	cursor int
	line   int
	column int
	token  *Token
}

func (t *Tokenizer) Rule(name tokens.Type, regex string, matcher func(*Token) *Token) {
	t.rules = append(t.rules, rule{
		Name:    name,
		Regex:   *regexp.MustCompile(regex),
		Ignore:  false,
		Matcher: matcher,
	})
}

func (t *Tokenizer) Ignore(name tokens.Type, regex string) {
	t.rules = append(t.rules, rule{
		Name:    name,
		Regex:   *regexp.MustCompile(regex),
		Ignore:  true,
		Matcher: nil,
	})
}

func (t *Tokenizer) Parse(input string) []*Token {
	return nil
}

func (t *Tokenizer) Next() *Token {
	return nil
}

func (t *Tokenizer) Peek() *Token {
	return nil
}
