package obj

type String struct {
	BaseObject
	Value string
}

func (n *String) Type() Type {
	return TString
}

func (n *String) AsString() string {
	return n.Value
}

func (n *String) AsBool() bool {
	return n.Value != ""
}

func (n *String) AsNumber() float64 {
	if n.AsBool() {
		return 1
	}

	return 0
}
