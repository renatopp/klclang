package obj

import (
	"strings"
)

type List struct {
	BaseObject

	Values []Object
}

func (n *List) Type() Type {
	return TList
}

func (n *List) AsString() string {
	b := strings.Builder{}
	s := len(n.Values)

	b.WriteString("[")
	for i, v := range n.Values {
		b.WriteString(v.AsString())
		if i < s-1 {
			b.WriteString(", ")
		}
	}
	b.WriteString("]")

	return b.String()
}

func (n *List) AsBool() bool {
	return len(n.Values) > 0
}

func (n *List) AsNumber() float64 {
	if n.AsBool() {
		return 1
	}

	return 0
}
