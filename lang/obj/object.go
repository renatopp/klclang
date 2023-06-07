package obj

import "fmt"

type Type string

const (
	TNumber Type = "number"
	TString      = "string"
)

type Object interface {
	Type() Type
	Inspect() string
	Bool() bool
}

type Number struct {
	Value float64
}

func (n *Number) Type() Type {
	return TNumber
}

func (n *Number) Inspect() string {
	return fmt.Sprintf("%f", n.Value)
}

func (n *Number) Bool() bool {
	return n.Value != 0
}

type String struct {
	Value string
}

func (n *String) Type() Type {
	return TString
}

func (n *String) Inspect() string {
	return n.Value
}

func (n *String) Bool() bool {
	return n.Value != ""
}
