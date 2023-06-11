package obj

import (
	"klc/lang/ast"
	"strings"
)

type ScopedStore interface{}

type FunctionParam struct {
	Name    string
	Default Object
	Spread  bool
}

func (n *FunctionParam) String(builder *strings.Builder) {
	if n.Spread {
		builder.WriteString("...")
	}
	builder.WriteString(n.Name)
	if n.Default != nil {
		builder.WriteString(" = ")
		builder.WriteString(n.Default.AsString())
	}
}

type Function struct {
	BaseObject

	Scope  ScopedStore
	Params []*FunctionParam
	Body   ast.Node
}

func (n *Function) Type() Type {
	return TFunction
}

func (n *Function) AsString() string {
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

func (n *Function) AsBool() bool {
	return true
}

func (n *Function) AsNumber() float64 {
	if n.AsBool() {
		return 1
	}

	return 0
}

func (n *Function) GetParams() []*FunctionParam {
	return n.Params
}

func (n *Function) GetScope() ScopedStore {
	return n.Scope
}
