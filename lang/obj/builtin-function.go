package obj

import (
	"klc/lang/ast"
	"strings"
)

type Evaluator interface {
	Eval(n ast.Node) Object
	Call(fn Callable, args []Object) Object
}

type BuiltinFunctionFn func(ev Evaluator, args ...Object) Object

type BuiltinFunction struct {
	BaseObject

	Params []*FunctionParam
	Fn     BuiltinFunctionFn
}

func (n *BuiltinFunction) Type() Type {
	return TFunction
}

func (n *BuiltinFunction) AsString() string {
	builder := strings.Builder{}
	builder.WriteString("fn (")
	for i, p := range n.Params {
		p.String(&builder)
		if i < len(n.Params)-1 {
			builder.WriteString(", ")
		}
	}
	builder.WriteString(") { ... }")

	return builder.String()
}

func (n *BuiltinFunction) AsBool() bool {
	return true
}

func (n *BuiltinFunction) AsNumber() float64 {
	if n.AsBool() {
		return 1
	}

	return 0
}

func (n *BuiltinFunction) GetParams() []*FunctionParam {
	return n.Params
}
