package obj

type Function struct {
	BaseObject
}

func (n *Function) Type() Type {
	return TFunction
}

func (n *Function) AsString() string {
	return "NOT IMPLEMENTED"
}

func (n *Function) AsBool() bool {
	// evaluate?
	// return n.Value
	return false
}

func (n *Function) AsNumber() float64 {
	if n.AsBool() {
		return 1
	}

	return 0
}
