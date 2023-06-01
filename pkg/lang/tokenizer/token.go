package tokenizer

import (
	"encoding/json"
	"fmt"
	"klc/pkg/lang/tokens"
)

type Token struct {
	Name    tokens.Type
	Literal string
}

func (t Token) ToString() string {
	json, _ := json.Marshal(t.Literal)
	return fmt.Sprintf("<%s:%s>", t.Name, json)
}
