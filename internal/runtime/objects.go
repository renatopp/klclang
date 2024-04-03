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
func (n *Number) Docs() string     { return n.docs }
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
	Params []asts.Node
	Body   asts.Node
}

func (f FunctionMatch) DebugString() string {
	params := ""
	for _, p := range f.Params {
		params += p.String() + ", "
	}

	return fmt.Sprintf("%s => %s", params, f.Body.String())
}

type Function struct {
	docs    string
	Matches []FunctionMatch
	Scope   *Scope
}

func NewFunction(parent *Scope) *Function {
	return &Function{
		Matches: make([]FunctionMatch, 0),
		Scope:   parent.New(),
	}
}

func (f *Function) AddMatch(params []asts.Node, body asts.Node) error {
	f.Matches = append(f.Matches, FunctionMatch{Params: params, Body: body})
	return nil
}

func (f *Function) SetDocs(d string) { f.docs = d }
func (f *Function) Docs() string     { return f.docs }
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
func (f *BuiltinFunction) Docs() string     { return f.docs }
func (f *BuiltinFunction) Type() ObjectType { return BUILTIN_FUNCTION }
func (f *BuiltinFunction) Bool() bool       { return true }
func (f *BuiltinFunction) Number() float64  { return 1 }
func (f *BuiltinFunction) String() string   { return "<function>" }

// ----------------------------------------------------------------------------
// Error
// ----------------------------------------------------------------------------
var ERROR ObjectType = "Error"

type Error struct {
	docs string
	err  string
}

func NewError(err string, v ...any) *Error {
	return &Error{
		err: fmt.Sprintf(err, v...),
	}
}

func (f *Error) SetDocs(d string) { f.docs = d }
func (f *Error) Docs() string     { return f.docs }
func (f *Error) Type() ObjectType { return ERROR }
func (f *Error) Bool() bool       { return false }
func (f *Error) Number() float64  { return 0 }
func (f *Error) String() string   { return f.err }
