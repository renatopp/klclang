package obj

import (
	"fmt"
	"math"
)

type Number struct {
	BaseObject
	Value float64
}

func (n *Number) Type() Type {
	return TNumber
}

func (n *Number) AsString() string {
	if n.IsInteger() {
		return fmt.Sprintf("%.0f", n.Value)
	} else {
		return fmt.Sprintf("%f", n.Value)
	}
}

func (n *Number) AsBool() bool {
	return n.Value != 0
}

func (n *Number) AsNumber() float64 {
	return n.Value
}

func (n *Number) IsInteger() bool {
	return math.Mod(n.Value, 1.0) == 0
}
