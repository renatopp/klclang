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

type FunctionMatch struct {
	Args []asts.Node
	Body asts.Node
}

type Function struct {
	docs    string
	Matches []FunctionMatch
	Scope   *Scope
}

// TODO: Set scope here
func NewFunction() *Function {
	return &Function{}
}

func (f *Function) AddMatch(args []asts.Node, body asts.Node) error {
	// TODO: validate matches here
	f.Matches = append(f.Matches, FunctionMatch{Args: args, Body: body})
	return nil
}

func (f *Function) SetDocs(d string) { f.docs = d }
func (f *Function) Docs() string     { return "" }
func (f *Function) Type() ObjectType { return FUNCTION }
func (f *Function) Bool() bool       { return true }
func (f *Function) Number() float64  { return 1 }
func (f *Function) String() string   { return "<function>" }

// ----------------------------------------------------------------------------
// Builtin Function
// ----------------------------------------------------------------------------
var BUILTIN_FUNCTION ObjectType = "BuiltinFunction"

type BuiltinFunction struct {
	docs string
	Fn   func(env *Scope, args ...Object) Object
}

func NewBuiltinFunction(fn func(env *Scope, args ...Object) Object) *BuiltinFunction {
	return &BuiltinFunction{
		Fn: fn,
	}
}

func (f *BuiltinFunction) SetDocs(d string) { f.docs = d }
func (f *BuiltinFunction) Docs() string     { return "" }
func (f *BuiltinFunction) Type() ObjectType { return BUILTIN_FUNCTION }
func (f *BuiltinFunction) Bool() bool       { return true }
func (f *BuiltinFunction) Number() float64  { return 1 }
func (f *BuiltinFunction) String() string   { return "<function>" }
