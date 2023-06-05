package token

import (
	"encoding/json"
	"fmt"
)

type Token struct {
	Type    Type
	Literal string
	At      At
}

type At struct {
	Cursor int
	Line   int
	Column int
}

func (t *Token) Is(ty Type) bool {
	return t.Type == ty
}

func (t *Token) ToString() string {
	json, _ := json.Marshal(t.Literal)
	return fmt.Sprintf("<%s@%d,%d:%s>", t.Type, t.At.Line, t.At.Column, json)
}

func Create(t Type, l string, at At) *Token {
	return &Token{
		Type:    t,
		Literal: l,
		At:      at,
	}
}
func CreateEmpty(t Type, at At) *Token {
	return &Token{
		Type:    t,
		Literal: "",
		At:      at,
	}
}
