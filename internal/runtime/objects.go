package runtime

import (
	"fmt"
	"math"

	"github.com/renatopp/langtools/asts"
)

type ObjectType string

type Object interface {
	Type() ObjectType
	SetDocs(string)
	Docs() string

	Bool() bool
	Number() float64
	String() string
}

func withDocs(obj Object, docs string) Object {
	obj.SetDocs(docs)
	return obj
}

// ----------------------------------------------------------------------------
// Number
// ----------------------------------------------------------------------------
var NUMBER ObjectType = "Number"

type Number struct {
	docs  string
	Value float64
}

func NewNumber(value float64) *Number {
	return &Number{Value: value}
}

func (n *Number) SetDocs(d string) { n.docs = d }
func (n *Number) Docs() string     { return "" }
func (n *Number) Type() ObjectType { return NUMBER }
func (n *Number) Bool() bool       { return n.Value != 0 }
func (n *Number) Number() float64  { return n.Value }
func (n *Number) String() string {
	v := n.Value
	if math.IsInf(v, 1) {
		return "inf"
	} else if math.IsInf(v, -1) {
		return "-inf"
	}

	if math.Mod(v, 1.0) == 0 {
		return fmt.Sprintf("%.0f", v)
	}

	return fmt.Sprintf("%f", v)
}

// ----------------------------------------------------------------------------
// Function
// ----------------------------------------------------------------------------
var FUNCTION ObjectType = "Function"

type Function struct {
	docs  string
	Args  []asts.Node
	Body  asts.Node
	Scope *Scope
}

// func NewFunction(value float64) *Number {
// 	return &Number{Value: value}
// }

func (f *Function) SetDocs(d string) { f.docs = d }
func (f *Function) Docs() string     { return "" }
func (f *Function) Type() ObjectType { return FUNCTION }
func (f *Function) Bool() bool       { return true }
func (f *Function) Number() float64  { return 1 }
func (f *Function) String() string   { return "<function>" }
